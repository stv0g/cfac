package bsis

import "time"

// TODO: move to dedicated OWS package

type OWSFeatureGeometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type OWSFeatureProperties struct {
	Gid            int         `json:"gid"`
	SbodGid1       int         `json:"sbod_gid1"`
	SdatenID       int         `json:"sdatenid"`
	Name           string      `json:"name"`
	Bereich        interface{} `json:"bereich"`
	Strassen       string      `json:"strassen"`
	Zeitraum       string      `json:"zeitraum"`
	Art            string      `json:"art"`
	Symbol         string      `json:"symbol"`
	Beschreibung   string      `json:"beschreibu"`
	Firma          interface{} `json:"firma"`
	Bauherr        string      `json:"bauherr"`
	Kontakt        interface{} `json:"kontakt"`
	Einschraenkung string      `json:"einschraen"`
	AachenDe       interface{} `json:"aachen_de"`
	ExternDe       interface{} `json:"extern_de"`
	Date           string      `json:"date"`
	Hotlink        string      `json:"hotlink"`
	Von            string      `json:"von"`
	Bis            string      `json:"bis"`
	SortColumn     int         `json:"sort_column"`
}

type OWSFeature struct {
	Type         string               `json:"type"`
	ID           string               `json:"id"`
	Geometry     OWSFeatureGeometry   `json:"geometry"`
	GeometryName string               `json:"geometry_name"`
	Properties   OWSFeatureProperties `json:"properties"`
}

type OWSCRS struct {
	Type       string `json:"type"`
	Properties struct {
		Name string `json:"name"`
	} `json:"properties"`
}

type OWS struct {
	Type           string       `json:"type"`
	Features       []OWSFeature `json:"features"`
	TotalFeatures  int          `json:"totalFeatures"`
	NumberMatched  int          `json:"numberMatched"`
	NumberReturned int          `json:"numberReturned"`
	TimeStamp      time.Time    `json:"timeStamp"`
	CRS            OWSCRS       `json:"crs"`
}
