package cfac

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func LastUpdated(r *colly.Response) (time.Time, error) {
	lastUpdatedStr := r.Headers.Get("Date")
	return time.Parse(time.RFC1123, lastUpdatedStr)
}

type UrlArgs map[string]interface{}

func PrepareUrl(tpl string, args UrlArgs) string {
	oldnew := []string{}
	for k, v := range args {
		oldnew = append(oldnew, fmt.Sprintf("{%s}", k), fmt.Sprintf("%v", v))
	}

	r := strings.NewReplacer(oldnew...)

	return r.Replace(tpl)
}

func DumpResponse(r *colly.Response) {
	var i interface{}

	if err := json.Unmarshal(r.Body, &i); err == nil {
		if o, err := json.MarshalIndent(i, "", "  "); err == nil {
			os.Stdout.Write(o)
		} else {
			os.Stdout.Write(r.Body)
		}
	} else {
		os.Stdout.Write(r.Body)
	}
	os.Stdout.Write([]byte{'\n'})
}
