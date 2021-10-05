package abfallapp

const (
	Url = "https://aachen-abfallapp.regioit.de/abfall-app-aachen/rest"

	UrlStandorte     = Url + "/standorte?ort={ort_name}&ortsteil={ortsteil}&standortart={standortart_name}"
	UrlStandortArten = Url + "/standorte/standortarten?ort={ort_name}"
	UrlOrte          = Url + "/orte"
	UrlFraktionen    = Url + "/fraktionen"
	UrlStrassen      = Url + "/orte/{ort_id}/strassen"
	UrlOrtsteile     = Url + "/standorte/standortarten/{standortart_name}/ortsteile?ort={ort_name}" // ?kat=7581006 von stoffe
	UrlStoffe        = Url + "/stoffe?ort={ort_name}"
)

type ResponseFraktionen []Fraktion
type ResponseStandortArten []StandortArt
type ResponseOrte []Ort
type ResponseStrassen []Strasse
type ResponseTermine []Termin
type ResponseOrtsteile []struct {
	Ort      string `json:"ort"`
	Ortsteil string `json:"ortsteil"`
}
type ResponseStandorte []Standort
type ResponseStoffe []Stoff

type Stoff struct {
	ID        int       `json:"id"`
	Stoff     string    `json:"stoff"`
	Tipps     string    `json:"tipps"`
	Kategorie Kategorie `json:"kategorieText"`
}

type Ort struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Strasse struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	HausNrList   []HausNummer `json:"hausNrList"`
	Plz          int          `json:"plz"`
	OrtsteilName string       `json:"ortsteilName"`
	Ort          Ort          `json:"ort"`
}

type HausNummer struct {
	ID  int    `json:"id"`
	Nr  string `json:"nr"`
	Plz string `json:"plz"`
}

type StandortArt struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	IconNr   int         `json:"iconNr"`
	PlusTage interface{} `json:"plustage"`
}

type Termin struct {
	ID     int    `json:"id"`
	Bezirk Bezirk `json:"bezirk"`
	Datum  string `json:"datum"`
}

type Bezirk struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	GueltigAb  interface{} `json:"gueltigAb"`
	FraktionID int         `json:"fraktionId"`
}

type Kategorie struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Beschreibung string      `json:"beschreibung"`
	IconNr       int         `json:"iconNr"`
	Ort          string      `json:"ort"`
	Optionen     interface{} `json:"optionen"`
}

type Position struct {
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
	Elevation interface{} `json:"elevation"`
}

type TerminSchadstoffmobil struct {
	ID         int    `json:"id"`
	Datum      string `json:"datum"`
	Uhrzeit    string `json:"uhrzeit"`
	Alarmzeit  string `json:"alarmzeit"`
	Orte       string `json:"orte"`
	StandortID int    `json:"standortId"`
}

type Fraktion struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IconNr   int    `json:"iconNr"`
	FarbeRgb string `json:"farbeRgb"`
}

type Standort struct {
	ID                     int                     `json:"id"`
	Standortart            string                  `json:"standortart"`
	Anlagenbezeichnung     string                  `json:"anlagenbezeichnung"`
	Plz                    string                  `json:"plz"`
	Ort                    string                  `json:"ort"`
	Ortsteil               string                  `json:"ortsteil"`
	Strasse                string                  `json:"strasse"`
	Hausnummer             string                  `json:"hausnummer"`
	Zusatz                 string                  `json:"zusatz"`
	Oeffnungszeiten        string                  `json:"oeffnungszeiten"`
	IconNr                 int                     `json:"iconNr"`
	Position               Position                `json:"position"`
	Fuellstand             interface{}             `json:"fuellstand"`
	KategorieList          []Kategorie             `json:"kategorieList"`
	TermineSchadstoffmobil []TerminSchadstoffmobil `json:"termineSchadstoffmobil"`
}
