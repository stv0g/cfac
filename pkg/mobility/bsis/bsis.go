package bsis

import "time"

const (
	UrlApi = "https://bsis.aachen.de/geoserver/ows"
)

// curl 'https://bsis.aachen.de/geoserver/ows?service=WFS&typeName=BSIS:PUNKTE_ALLE&version=1.1.0&request=GetFeature&outputFormat=application%2Fjson&srsName=EPSG%3A3857' \
//   -H 'Connection: keep-alive' \
//   -H 'sec-ch-ua: "Chromium";v="93", " Not;A Brand";v="99"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'User-Agent: Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36' \
//   -H 'sec-ch-ua-platform: "Linux"' \
//   -H 'Accept: */*' \
//   -H 'Sec-Fetch-Site: same-origin' \
//   -H 'Sec-Fetch-Mode: cors' \
//   -H 'Sec-Fetch-Dest: empty' \
//   -H 'Referer: https://bsis.aachen.de/' \
//   -H 'Accept-Language: en-US,en;q=0.9' \
//   --compressed

// {
// 	"type":"FeatureCollection",
// 	"features":[
// 	   {
// 		  "type":"Feature",
// 		  "id":"PUNKTE_ALLE.35222",
// 		  "geometry":{
// 			 "type":"MultiPoint",
// 			 "coordinates":[
// 				[
// 				   677734.36648271,
// 				   6581715.56276454
// 				]
// 			 ]
// 		  },
// 		  "geometry_name":"the_geom",
// 		  "properties":{
// 			 "gid":35222,
// 			 "sbod_gid1":35222,
// 			 "sdatenid":68640,
// 			 "name":"Abrissarbeiten Parkhaus Büchel",
// 			 "bereich":"Mefferdatisstraße - Peterstraße",
// 			 "strassen":"Büchel, Peterstraße",
// 			 "zeitraum":"29.07.2021 - 31.12.2021",
// 			 "art":"Strassenbau",
// 			 "symbol":"Verkehrsbehinderung - aktuell",
// 			 "beschreibu":"Abrissarbeiten Parkhaus Büchel",
// 			 "firma":"BL Schröer GmbH  Baustellenabsicherung-Straßenausstattung",
// 			 "bauherr":"Stadt Aachen Gebäudemanagament",
// 			 "kontakt":null,
// 			 "einschraen":"Einrichtung einer Überfahrung.\r\nEinrichtung einer Baustellensignalisierung.",
// 			 "aachen_de":null,
// 			 "extern_de":null,
// 			 "date":"2021-09-21",
// 			 "hotlink":"https://bsis.aachen.de/html/Dokumente/00010C00/0x010C20/index.htm",
// 			 "von":"2021-07-29",
// 			 "bis":"2021-12-31",
// 			 "sort_column":1
// 		  }
// 	   },
// 	   {
// 		  "type":"Feature",
// 		  "id":"PUNKTE_ALLE.35689",
// 		  "geometry":{
// 			 "type":"MultiPoint",
// 			 "coordinates":[
// 				[
// 				   677602.31487247,
// 				   6580247.91105099
// 				]
// 			 ]
// 		  },
// 		  "geometry_name":"the_geom",
// 		  "properties":{
// 			 "gid":35689,
// 			 "sbod_gid1":35689,
// 			 "sdatenid":69441,
// 			 "name":"Verlegungsarbeiten Fernmeldeleitungen",
// 			 "bereich":"Hausnummer 2, bis Burtscheider Straße 1, inkl. SQ Friedlandstraße und SQ Burtscheider Straße",
// 			 "strassen":"Burtscheider Straße, Friedlandstraße, Reumontstraße",
// 			 "zeitraum":"03.09.2021 - 24.09.2021",
// 			 "art":"Strassenbau",
// 			 "symbol":"Verkehrsbehinderung - aktuell",
// 			 "beschreibu":"Glasfaserverlegung",
// 			 "firma":"Özkan Baugesellschaft mbH",
// 			 "bauherr":"RelAix Networks GmbH",
// 			 "kontakt":"keine Angabe",
// 			 "einschraen":"Arbeiten in den Nebenanlagen und am Fahrbahnrand. Es wird immer eine sichere Fußgängerführung hergestellt. Alle Fahrbeziehungen bleiben erhalten.",
// 			 "aachen_de":null,
// 			 "extern_de":null,
// 			 "date":"2021-09-21",
// 			 "hotlink":"https://bsis.aachen.de/html/Dokumente/00010C00/0x010F41/index.htm",
// 			 "von":"2021-09-03",
// 			 "bis":"2021-09-24",
// 			 "sort_column":1
// 		  }
// 	   },

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
