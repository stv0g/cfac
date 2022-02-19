package bsis

import (
	"time"

	"github.com/stv0g/cfac/pkg/std/ogc/wfs"
)

type ConstructionSiteList wfs.FeatureCollection

type ConstructionSiteProperties struct {
	GID            int         `json:"gid"`
	SbodGID1       int         `json:"sbod_gid1"`
	SdatenID       int         `json:"sdatenid"`
	Name           string      `json:"name"`
	Bereich        string      `json:"bereich"`
	Strassen       string      `json:"strassen"`
	Zeitraum       string      `json:"zeitraum"`
	Art            string      `json:"art"`
	Symbol         string      `json:"symbol"`
	Beschreibung   string      `json:"beschreibu"`
	Firma          string      `json:"firma,omitempty"`
	Bauherr        string      `json:"bauherr"`
	Kontakt        string      `json:"kontakt,omitempty"`
	Einschraenkung string      `json:"einschraen"`
	AachenDe       interface{} `json:"aachen_de"`
	ExternDe       interface{} `json:"extern_de"`
	Date           time.Time   `json:"date"`
	Hotlink        string      `json:"hotlink"`
	Von            time.Time   `json:"von"`
	Bis            time.Time   `json:"bis"`
	SortColumn     int         `json:"sort_column"`
}
