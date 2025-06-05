// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stv0g/cfac/pkg/oparl"
)

var REGEX_MEMBER_PICTURE = regexp.MustCompile(`(?m)<img src="([^"]+)".*class="sdnetrim-member-picture" \/>`)

func longestCommonPrefix(strs []string) string {
	longestPrefix := ""
	endPrefix := false

	if len(strs) > 0 {
		sort.Strings(strs)
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}

func ptrToBool(ptr *bool) bool {
	if ptr == nil {
		return false
	}

	return bool(*ptr)
}

type Error struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type Collector struct {
	*colly.Collector

	dir           string
	modifiedSince time.Time
	oparlEndpoint string

	prefix string
}

func NewCollector(dir string, oparlEndpoint string, modifiedSince time.Time) *Collector {
	c := &Collector{
		Collector: colly.NewCollector(
			colly.Async(true),
		),

		dir:           dir,
		modifiedSince: modifiedSince,
		oparlEndpoint: oparlEndpoint,
	}

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       200 * time.Millisecond,
	}); err != nil {
		panic(err)
	}

	c.IgnoreRobotsTxt = true

	c.OnRequest(c.onRequest)
	c.OnError(c.onError)

	c.OnResponse(func(r *colly.Response) {
		if err := c.onResponse(r); err != nil {
			logrus.WithError(err).Error("Failed to handle response")
		}
	})

	return c
}

func (c *Collector) Collect() error {
	if err := c.Visit(c.oparlEndpoint); err != nil {
		return err
	}

	c.Wait()

	return nil
}

func (c *Collector) Clone() *colly.Collector {
	c2 := c.Collector.Clone()

	c2.OnError(c.onError)
	c2.OnRequest(c.onRequest)

	return c2
}

func (c *Collector) fileName(id oparl.URL, ext string) string {
	fn := strings.TrimPrefix(string(id), string(c.prefix)) + ext
	fn = strings.TrimLeft(fn, "/")
	fn = strings.ReplaceAll(fn, "/", string(os.PathSeparator))
	fn = filepath.Join(c.dir, fn)

	return fn
}

func (c *Collector) onRequest(r *colly.Request) {
	logrus.WithField("url", r.URL).Info("Visiting")
}

func (c *Collector) onError(r *colly.Response, err error) {
	log := logrus.WithFields(logrus.Fields{
		"url": r.Request.URL,
	})

	var jErr Error
	if err := json.Unmarshal(r.Body, &jErr); err != nil {
		log.WithError(err).Error("Failed to decode JSON")
		log.WithError(err).Error("Failed to scrape")
	} else {
		if (jErr.Code == 0 && jErr.Error == "Keine Adressen vorhanden") ||
			(jErr.Code == 802 && jErr.Error == "Die angeforderte Ressource wurde nicht gefunden.") {
			log.Warn("Got empty response")
			return
		}

		log.WithError(err).WithFields(logrus.Fields{
			"code": jErr.Code,
			"msg":  jErr.Error,
			"type": jErr.Type,
		}).Error("Failed to scrape")
	}
}

func (c *Collector) onResponse(r *colly.Response) error {
	log := logrus.WithField("url", r.Request.URL)

	log.WithField("body", string(r.Body)).Trace("Body")

	// Check if response is empty
	if strings.TrimSpace(string(r.Body)) == "[]" {
		log.Warn("Got empty response")
		return nil
	}

	// Check if response is paginated
	var objs oparl.Paginated
	if err := json.Unmarshal(r.Body, &objs); err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal array into Go value") {
			return nil
		}

		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	if objs.Data == nil { // Non-paginated
		var obj oparl.Common
		if err := json.Unmarshal(r.Body, &obj); err != nil {
			return fmt.Errorf("failed to parse JSON: %w", err)
		}

		if err := c.onResource(r.Body, obj.ID, obj.Type, ptrToBool(obj.Deleted)); err != nil {
			return err
		}
	} else { // Paginated
		for _, raw := range objs.Data {
			var obj oparl.Common
			if err := json.Unmarshal(raw, &obj); err != nil {
				if strings.Contains(err.Error(), "failed to unmarshal array into Go value") {
					return nil
				}

				return fmt.Errorf("failed to parse JSON: %w", err)
			}

			if err := c.onResource(raw, obj.ID, obj.Type, ptrToBool(obj.Deleted)); err != nil {
				return err
			}
		}

		if objs.Links.Next != "" {
			if err := c.Visit(objs.Links.Next); err != nil && !errors.Is(err, &colly.AlreadyVisitedError{}) {
				return fmt.Errorf("failed to visit: %w", err)
			}
		}
	}

	return nil
}

func (c *Collector) onResource(buf []byte, id oparl.URL, typ string, deleted bool) error {
	if c.prefix == "" && typ == oparl.TypeSystem {
		obj := &oparl.System{}
		if err := json.Unmarshal(buf, obj); err != nil {
			return fmt.Errorf("failed to parse JSON: %w", err)
		}

		c.prefix = longestCommonPrefix([]string{
			string(obj.Body),
			string(obj.ID),
		})

		logrus.WithField("prefix", c.prefix).Info("Found prefix")
	}

	fn := c.fileName(id, ".json")

	log := logrus.WithFields(logrus.Fields{
		"id":   id,
		"file": fn,
	})

	if deleted {
		if _, err := os.Stat(fn); err == nil {
			log.Info("Deleting")
			if err := os.Remove(fn); err != nil {
				return fmt.Errorf("failed to delete file '%s': %w", fn, err)
			}
		}

		return nil
	}

	obj := oparl.NewResource(typ)
	if obj == nil {
		return fmt.Errorf("unsupported type: " + typ)
	}

	if err := json.Unmarshal(buf, obj); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	log.WithFields(logrus.Fields{
		"file": fn,
		"type": typ,
	}).Info("Got resource")

	dn := filepath.Dir(fn)
	if err := os.MkdirAll(dn, 0o755); err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", dn, err)
	}

	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open file '%s': %w", fn, err)
	}
	defer f.Close() //nolint:errcheck

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(obj); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	switch obj := obj.(type) {
	case *oparl.System:
		if err := c.Visit(string(obj.Body)); err != nil && !errors.Is(err, &colly.AlreadyVisitedError{}) {
			return fmt.Errorf("failed to visit: %w", err)
		}

	case *oparl.Body:
		for _, id := range []oparl.URL{
			obj.AgendaItem,
			obj.Consultation,
			obj.File,
			obj.LegislativeTermList,
			obj.LocationList,
			obj.Meeting,
			obj.Membership,
			obj.Organization,
			obj.Paper,
			obj.Person,
		} {
			u := string(id)

			if !c.modifiedSince.IsZero() {
				v := url.Values{}
				v.Add("modified_since", c.modifiedSince.Format(time.RFC3339))
				u += "?" + v.Encode()
			}

			if err := c.Visit(u); err != nil && !errors.Is(err, &colly.AlreadyVisitedError{}) {
				return fmt.Errorf("failed to visit: %w", err)
			}
		}

	case *oparl.Person:
		if obj.Web != nil {
			c2 := c.Clone()
			c2.OnResponse(func(r *colly.Response) {
				if err := c.onPersonWebProfile(r, obj); err != nil {
					log.WithError(err).Error("Failed to handle person")
				}
			})

			if err := c2.Visit(string(*obj.Web)); err != nil && !errors.Is(err, &colly.AlreadyVisitedError{}) {
				return fmt.Errorf("failed to visit person website: %w", err)
			}
		}

	case *oparl.File:
		c2 := c.Clone()
		c2.OnResponse(func(r *colly.Response) {
			if err := c.onFile(r, obj); err != nil {
				log.WithError(err).Error("Failed to handle file")
			}
		})

		if err := c2.Visit(string(obj.AccessURL)); err != nil && !errors.Is(err, &colly.AlreadyVisitedError{}) {
			return fmt.Errorf("failed to visit file: %w", err)
		}
	}

	return nil
}

func (c *Collector) onFile(r *colly.Response, f *oparl.File) error {
	fn := c.fileName(f.ID, extensionOfFile(f))

	logrus.WithField("file", fn).Info("Got file")

	dn := filepath.Dir(fn)
	if err := os.MkdirAll(dn, 0o755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return r.Save(fn)
}

func (c *Collector) onPersonWebProfile(r *colly.Response, p *oparl.Person) error {
	log := logrus.WithFields(logrus.Fields{
		"name": *p.Name,
		"id":   p.ID,
	})

	log.Info("Got person profile")

	m := REGEX_MEMBER_PICTURE.FindStringSubmatch(string(r.Body))
	if len(m) > 0 {
		url := m[1]
		log.WithField("url", url).Info("Found profile picture")

		c2 := c.Clone()
		c2.OnResponse(func(r *colly.Response) {
			if err := c.onPersonProfilePicture(r, p); err != nil {
				log.WithError(err).Error("handle profile picture")
			}
		})

		if err := c2.Visit(url); err != nil && !errors.Is(err, &colly.AlreadyVisitedError{}) {
			return fmt.Errorf("failed to visit: %w", err)
		}
	} else {
		log.Warning("Missing profile picture?")
	}

	return nil
}

func (c *Collector) onPersonProfilePicture(r *colly.Response, p *oparl.Person) error {
	log := logrus.WithFields(logrus.Fields{
		"name": *p.Name,
		"id":   p.ID,
	})

	exts, err := mime.ExtensionsByType(r.Headers.Get("Content-Type"))
	if err != nil {
		return fmt.Errorf("failed to get file extension for profile picture: %w", err)
	}

	fn := c.fileName(p.ID, exts[0])
	log.WithField("file", fn).Info("Got person picture")

	return r.Save(fn)
}

func extensionOfFile(f *oparl.File) string {
	if mt := f.MimeType; mt != nil {
		exts, err := mime.ExtensionsByType(*mt)
		if err != nil {
			panic(err)
		} else if len(exts) > 0 {
			return exts[0]
		}
	}

	return filepath.Ext(string(f.AccessURL))
}

func main() {
	var dir string
	var endpoint string
	var modifiedSince string

	// Root command
	rootCmd := &cobra.Command{
		Use:   "scraper_oparl",
		Short: "Scraper for OParl data",
		Run: func(cmd *cobra.Command, args []string) {
			// Parse modifiedSince into time.Time
			var modSince time.Time
			if modifiedSince != "" {
				var err error
				modSince, err = time.Parse(time.RFC3339, modifiedSince)
				if err != nil {
					logrus.WithError(err).Fatal("Invalid modifiedSince format")
				}
			}

			// Create collector and start scraping
			c := NewCollector(dir, endpoint, modSince)
			if err := c.Collect(); err != nil {
				logrus.WithError(err).Fatal("Failed to fetch data")
			}
		},
	}

	// Bind flags
	rootCmd.Flags().StringVarP(&dir, "dir", "d", "data/ratsinfo.simmerath.de", "Directory to store scraped data")
	rootCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "https://ratsinfo.simmerath.de/webservice/oparl/v1.1/system", "OParl endpoint URL")
	rootCmd.Flags().StringVarP(&modifiedSince, "modified-since", "m", "", "Fetch data modified since (RFC3339 format)")

	// Bind flags to Viper
	viper.BindPFlag("dir", rootCmd.Flags().Lookup("dir"))                       //nolint:errcheck
	viper.BindPFlag("endpoint", rootCmd.Flags().Lookup("endpoint"))             //nolint:errcheck
	viper.BindPFlag("modified-since", rootCmd.Flags().Lookup("modified-since")) //nolint:errcheck

	viper.AutomaticEnv()
	viper.SetEnvPrefix("cfac")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Failed to execute command")
	}
}
