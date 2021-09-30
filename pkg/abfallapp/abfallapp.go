package abfallapp

const (
	UrlAPI          = "https://aachen-abfallapp.regioit.de/abfall-app-aachen/rest"
	UrlAPIStandorte = UrlAPI + "/standorte"
)

type Kategorie struct {
	Text         string `xml:",chardata"`
	Beschreibung string `xml:"beschreibung"`
	IconNr       string `xml:"iconNr"`
	ID           string `xml:"id"`
	Name         string `xml:"name"`
}

type Position struct {
	Text      string `xml:",chardata"`
	Latitude  string `xml:"latitude"`
	Longitude string `xml:"longitude"`
}

type TerminSchadstoffmobil struct {
	Text       string `xml:",chardata"`
	Alarmzeit  string `xml:"alarmzeit"`
	Datum      string `xml:"datum"`
	ID         string `xml:"id"`
	Orte       string `xml:"orte"`
	StandortId string `xml:"standortId"`
	Uhrzeit    string `xml:"uhrzeit"`
}

type Standort struct {
	Text                   string                  `xml:",chardata"`
	AnlagenBezeichnung     string                  `xml:"anlagenbezeichnung"`
	Hausnummer             string                  `xml:"hausnummer"`
	IconNr                 string                  `xml:"iconNr"`
	ID                     string                  `xml:"id"`
	Kategorien             []Kategorie             `xml:"kategorieList"`
	Oeffnungszeiten        string                  `xml:"oeffnungszeiten"`
	Ort                    string                  `xml:"ort"`
	Ortsteil               string                  `xml:"ortsteil"`
	Plz                    string                  `xml:"plz"`
	Position               Position                `xml:"position"`
	Standortart            string                  `xml:"standortart"`
	Strasse                string                  `xml:"strasse"`
	Zusatz                 string                  `xml:"zusatz"`
	TermineSchadstoffmobil []TerminSchadstoffmobil `xml:"termineSchadstoffmobil"`
}
