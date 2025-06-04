// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package oparl

// https://dev.oparl.org/spezifikation

import (
	"encoding/json"
	"time"

	"github.com/stv0g/cfac/pkg/std/ietf/geojson"
)

const (
	schemaPrefix = "https://schema.oparl.org/1.1/"

	TypeSystem          = schemaPrefix + "System"
	TypeBody            = schemaPrefix + "Body"
	TypeOrganization    = schemaPrefix + "Organization"
	TypePerson          = schemaPrefix + "Person"
	TypeMeeting         = schemaPrefix + "Meeting"
	TypeMembership      = schemaPrefix + "Membership"
	TypePaper           = schemaPrefix + "Paper"
	TypeFile            = schemaPrefix + "File"
	TypeConsultation    = schemaPrefix + "Consultation"
	TypeAgendaItem      = schemaPrefix + "AgendaItem"
	TypeLegislativeTerm = schemaPrefix + "LegislativeTerm"
	TypeLocation        = schemaPrefix + "Location"
)

type Resource interface{}

func NewResource(typ string) Resource {
	switch typ {
	case TypeSystem:
		return &System{}
	case TypeBody:
		return &Body{}
	case TypeOrganization:
		return &Organization{}
	case TypePerson:
		return &Person{}
	case TypeMeeting:
		return &Meeting{}
	case TypeMembership:
		return &Membership{}
	case TypePaper:
		return &Paper{}
	case TypeFile:
		return &File{}
	case TypeConsultation:
		return &Consultation{}
	case TypeAgendaItem:
		return &AgendaItem{}
	case TypeLegislativeTerm:
		return &LegislativeTerm{}
	case TypeLocation:
		return &Location{}
	default:
		return nil
	}
}

type URL string

type Paginated struct {
	Data       []json.RawMessage `json:"data"`
	Pagination struct {
		TotalElements   int `json:"totalElements"`
		ElementsPerPage int `json:"elementsPerPage"`
		CurrentPage     int `json:"currentPage"`
		TotalPages      int `json:"totalPages"`
	} `json:"pagination"`
	Links struct {
		First string `json:"first"`
		Self  string `json:"self"`
		Last  string `json:"last"`
		Next  string `json:"next"`
	} `json:"links"`
}

type Common struct {
	ID       URL        `json:"id"`
	Type     string     `json:"type"`
	Created  *time.Time `json:"created,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
	Keyword  []string   `json:"keyword,omitempty"`
	Deleted  *bool      `json:"deleted,omitempty"`

	// Lizenz, unter der durch diese API abruf­ba­ren Daten stehen, sofern nicht am einzell­nen Objekt anders ange­ge­ben.
	License *URL `json:"license,omitempty"`
	Web     *URL `json:"web,omitempty"`
}

// Ein oparl:System-Objekt reprä­sen­tiert eine OParl-Schnitt­stelle für eine bestimmte OParl-Version.
// Es ist außer­dem der Start­punkt für Clients beim Zugriff auf einen Server.
// Möchte ein Server mehrere zuein­an­der inkom­pa­ti­ble OParl-Versio­nen unter­stüt­zen, dann muss der Server für jede Version eine eige­nen OParl-Schnitt­stelle mit einem eige­nen System-Objekt ausge­ben.
type System struct {
	Common

	// Die URL der OParl-Spezi­fi­ka­tion, die von diesem Server unter­stützt wird.
	// Aktu­ell kommt hier nur ein Wert in Frage.
	// Mit zukünf­ti­gen OParl-Versio­nen kommen weitere mögli­che URLs hinzu.
	// Wert: https://schema.oparl.org/1.1/
	OparlVersion string `json:"oparlVersion"`

	// Link zur Objekt­liste mit allen Körper­schaf­ten, die auf dem System exis­tie­ren.
	Body URL `json:"body"`

	// Nutzer­freund­li­cher Name für das System, mit dessen Hilfe Nutze­rin­nen und Nutzer das System erken­nen und von ande­ren unter­schei­den können.
	Name *string `json:"name,omitempty"`

	// E-Mail-Adresse für Anfra­gen zur OParl-API.
	// Die Angabe einer E-Mail-Adresse dient sowohl NutzerIn­nen wie auch Entwick­le­rin­nen von Clients zur Kontakt­auf­nahme mit dem Betrei­ber.
	ContactEmail *string `json:"contactEmail,omitempty"`

	// Name der Ansprech­part­ne­rin bzw. des Ansprech­part­ners oder der Abtei­lung, die über die in contactEmail ange­ge­bene Adresse erreicht werden kann.
	ContactName *string `json:"contactName,omitempty"`

	// URL der Website des parla­men­ta­ri­schen Infor­ma­ti­ons­sys­tems.
	Website *URL `json:"website,omitempty"`

	// URL der Website des Soft­wa­rean­bie­ters, von dem die OParl-Server-Soft­ware stammt.
	Vendor *URL `json:"vendor,omitempty"`

	// URL zu Infor­ma­tio­nen über die auf dem System genutzte OParl-Server-Soft­ware.
	Product *URL `json:"product,omitempty"`

	// Dient der Angabe von System-Objek­ten mit ande­ren OParl-Versio­nen.
	OtherOparlVersions *[]string `json:"otherOparlVersions,omitempty"`
}

// Der Objekt­typ oparl:Body dient dazu, eine Körper­schaft zu reprä­sen­tie­ren.
// Eine Körper­schaft ist in den meis­ten Fällen eine Gemeinde, eine Stadt oder ein Land­kreis.
// In der Regel sind auf einem OParl-Server Daten von genau einer Körper­schaft gespei­chert und es wird daher auch nur ein Body-Objekt ausge­ge­ben.
// Sind auf dem Server jedoch Daten von mehre­ren Körper­schaf­ten gespei­chert, muss für jede Körper­schaft ein eige­nes Body-Objekt ausge­ge­ben werden.
type Body struct {
	Common

	// System, zu dem dieses Objekt gehört.
	System *URL `json:"system,omitempty"`

	// Dient der Angabe einer Kontakt-E-Mail-Adresse.
	// Die Adresse soll die Kontakt­auf­nahme zu einer für die Körper­schaft und idea­ler­weise das parla­men­ta­ri­sche Infor­ma­ti­ons­sys­tem zustän­di­gen Stelle ermög­li­chen.
	ContactEmail *string `json:"contactEmail,omitempty"`

	// Name oder Bezeich­nung der mit contactEmail erreich­ba­ren Stelle.
	ContactName *string `json:"contactName,omitempty"`

	// Der acht­stel­lige Amtli­che Gemein­de­schlüs­sel.
	AGS *string `json:"ags,omitempty"`

	// Der zwölf­stel­lige Regio­nal­schlüs­sel.
	RGS *string `json:"rgs,omitempty"`

	// Dient der Angabe zusätz­li­cher URLs, die dieselbe Körper­schaft reprä­sen­tie­ren
	// Hier können beispiels­weise der entspre­chende Eintrag der gemein­sa­men Norm­da­tei der Deut­schen Natio­nal­bi­blio­thek, der DBPe­dia oder der Wiki­pe­dia ange­ge­ben werden.
	// Body- oder System-Objekte mit ande­ren OParl-Versio­nen dürfen nicht Teil der Liste sein.
	Equivalent []URL `json:"equivalent,omitempty"`

	// Kurzer Name der Körper­schaft.
	ShortName *string `json:"shortName,omitempty"`

	// Der offi­zi­elle lange Name der Körper­schaft.
	Name string `json:"name"`

	// Allge­meine Website der Körper­schaft.
	Website *URL `json:"website,omitempty"`

	// Zeit­punkt, seit dem die unter license ange­ge­bene Lizenz gilt.
	// Vorsicht bei Ände­run­gen der Lizenz die zu restrik­ti­veren Bedin­gun­gen führen!
	LicenseValidSince *time.Time `json:"licenseValidSince,omitempty"`

	// Link zur Objekt­liste mit allen Grup­pie­run­gen der Körper­schaft.
	Organization URL `json:"organization"`

	// Link zur Objekt­liste mit allen Perso­nen der Körper­schaft.
	Person URL `json:"person"`

	// Link zur Objekt­liste mit allen Sitzun­gen der Körper­schaft.
	Meeting URL `json:"meeting"`

	// Link zur Objekt­liste mit allen Druck­sa­chen der Körper­schaft.
	Paper URL `json:"paper"`

	// Link zur Objekt­liste mit allen Tages­ord­nungs­punk­ten der Körper­schaft.
	// Neu in OParl 1.1.
	AgendaItem URL `json:"agendaItem"`

	// Link zur Objekt­liste mit allen Bera­tun­gen der Körper­schaft.
	// Neu in OParl 1.1.
	Consultation URL `json:"consultation"`

	// Link zur Objekt­liste mit allen Dateien der Körper­schaft.
	// Neu in OParl 1.1.
	File URL `json:"file"`

	// Link zur Objekt­liste mit allen Orts­an­ga­ben der Körper­schaft.
	// Neu in OParl 1.1.
	LocationList URL `json:"locationList"`

	// Link zur Objekt­liste mit allen Mitglied­schaf­ten der Körper­schaft.
	// Neu in OParl 1.1.
	Membership URL `json:"membership"`

	// Link zur Objekt­liste mit allen Legis­la­tur­pe­ri­oden der Körper­schaft.
	// Die externe Objekt­liste enthält die glei­chen Objekte wie legislativeTerm
	// Neu in OParl 1.1.
	LegislativeTermList URL `json:"legislativeTermList"`

	// Objekt­liste mit den Wahl­pe­ri­oden der Körper­schaft.
	LegislativeTerm []LegislativeTerm `json:"legislativeTerm"`

	// Ort, an dem die Körper­schaft behei­ma­tet ist.
	Location *Location `json:"location,omitempty"`

	// Art der Körper­schaft.
	Classification *string `json:"classification,omitempty"`

	// Zeit­punkt, ab dem OParl für dieses Body bereit­ge­stellt wurde.
	// Dies hilft, um die Daten­qua­li­tät einzu­schät­zen, denn erst ab der Einrich­tung für OParl kann sicher­ge­stellt werden, dass sämt­li­che Werte korrekt in der Origi­nal-Quelle vorlie­gen.
	OParlSince *time.Time `json:"oparlSince,omitempty"`
}

// LegislativeTerm dient der Beschrei­bung einer Wahl­pe­ri­ode.
type LegislativeTerm struct {
	Common

	// Rück­re­fe­renz auf die Körper­schaft, welche nur dann ausge­ge­ben werden muss, wenn das Legis­la­ti­veTerm-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Body *URL `json:"body,omitempty"`

	// Nutzer­freund­li­che Bezeich­nung der Wahl­pe­ri­ode.
	Name *string `json:"name,omitempty"`

	// Der erste Tag der Wahl­pe­ri­ode.
	StartDate *string `json:"startDate,omitempty"`

	// Der letzte Tag der Wahl­pe­ri­ode.
	EndDate *string `json:"endDate,omitempty"`
}

// Location dient dazu, einen Orts­be­zug formal abzu­bil­den.
// Orts­an­ga­ben können sowohl aus Text­in­for­ma­tio­nen beste­hen (beispiels­weise dem Namen einer Straße/eines Plat­zes oder eine genaue Adresse) als auch aus Geoda­ten.
// Orts­an­ga­ben sind auch nicht auf einzelne Posi­tio­nen beschränkt, sondern können eine Viel­zahl von Posi­tio­nen, Flächen, Stre­cken etc. abde­cken.
type Location struct {
	Common

	Description *string          `json:"description,omitempty"`
	GeoJSON     *geojson.Feature `json:"geojson,omitempty"`

	// Straße und Haus­num­mer der Anschrift.
	StreetAddress *string `json:"streetAddress,omitempty"`

	// Raum­an­gabe der Anschrift.
	Room *string `json:"room,omitempty"`

	// Post­leit­zahl der Anschrift.
	PostalCode *string `json:"postalCode,omitempty"`

	// Unter­ge­ord­nete Orts­an­gabe der Anschrift, z.B. Stadt­be­zirk, Orts­teil oder Dorf.
	SubLocality *string `json:"subLocality,omitempty"`

	// Orts­an­gabe der Anschrift.
	Locality *string `json:"locality,omitempty"`

	// Rück­re­fe­ren­zen auf Body-Objekte. Wird nur dann ausge­ge­ben, wenn das Loca­tion-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Bodies []URL `json:"bodies,omitempty"`

	// Rück­re­fe­ren­zen auf Orga­ni­za­tion-Objekte. Wird nur dann ausge­ge­ben, wenn das Loca­tion-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Organizations []URL `json:"organizations,omitempty"`

	// Rück­re­fe­ren­zen auf Person-Objekte. Wird nur dann ausge­ge­ben, wenn das Loca­tion-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Persons []URL `json:"persons,omitempty"`

	// Rück­re­fe­ren­zen auf Meeting-Objekte. Wird nur dann ausge­ge­ben, wenn das Loca­tion-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Meetings []URL `json:"meetings,omitempty"`

	// Rück­re­fe­ren­zen auf Paper-Objekte. Wird nur dann ausge­ge­ben, wenn das Loca­tion-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Papers []URL `json:"papers,omitempty"`
}

// Organization dient dazu, Grup­pie­run­gen von Perso­nen abzu­bil­den, die in der parla­men­ta­ri­schen Arbeit eine Rolle spie­len.
// Dazu zählen in der Praxis insbe­son­dere Frak­tio­nen und Gremien.
type Organization struct {
	Common

	// Körper­schaft, zu der diese Grup­pie­rung gehört.
	Body *URL `json:"body,omitempty"`

	// Exter­ner OParl Body, der dieser Orga­ni­sa­tion entspricht.
	// Diese Eigen­schaft ist dafür gedacht auf even­tu­elle konkre­tere OParl-Schnitt­stel­len zu verwei­sen.
	// Ein Beispiel hier­für wäre eine Stadt, die sowohl ein über­grei­fen­des parla­men­ta­ri­sches Infor­ma­ti­ons­sys­tem, als auch bezirkss­pe­zi­fi­sche Systeme hat.
	ExternalBody *URL `json:"externalBody,omitempty"`

	// Offi­zi­elle (lange) Form des Namens der Grup­pie­rung.
	Name *string `json:"name,omitempty"`

	// Der Name der Grup­pie­rung als Kurz­form.
	ShortName *string `json:"shortName,omitempty"`

	// Grün­dungs­da­tum der Grup­pie­rung.
	// Kann z. B. das Datum der konsti­tu­ie­ren­den Sitzung sein.
	StartDate *string `json:"startDate,omitempty"`

	// Datum des letz­ten Tages der Exis­tenz der Grup­pie­rung.
	EndDate *string `json:"endDate,omitempty"`

	// Grobe Kate­go­ri­sie­rung der Grup­pie­rung.
	// Mögli­che Werte sind “Gremium”, “Partei”, “Frak­tion”, “Verwal­tungs­be­reich”, “exter­nes Gremium”, “Insti­tu­tion” und “Sons­ti­ges”.
	OrganizationType *string `json:"organizationType,omitempty"`

	// Ort, an dem die Orga­ni­sa­tion behei­ma­tet ist
	Location *Location `json:"location,omitempty"`

	// Posi­tio­nen, die für diese Grup­pie­rung vorge­se­hen sind.
	Post []string `json:"post,omitempty"`

	// URL auf eine externe Objekt­liste mit den Bera­tun­gen dieser Grup­pie­rung.
	// Invers zur Eigen­schaft organization der Klasse oparl:Consultation
	Consultation *URL `json:"consultation,omitempty"`

	// URL auf eine externe Objekt­liste mit den Sitzun­gen dieser Grup­pie­rung.
	// Invers zur Eigen­schaft organization der Klasse oparl:Meeting
	Meeting *URL `json:"meeting,omitempty"`

	// URL einer even­tu­el­len über­ge­ord­ne­ten Grup­pie­rung.
	SubOrganizationOf *URL `json:"subOrganizationOf,omitempty"`

	// Mitglied­schaf­ten dieser Grup­pie­rung.
	Membership []URL `json:"membership,omitempty"`

	// Die Art der Grup­pie­rung.
	// In Frage kommen z.B. “Parla­ment”, “Ausschuss”, “Beirat”, “Projekt­bei­rat”, “Kommis­sion”, “AG”, “Verwal­tungs­rat”, “Frak­tion” oder “Partei”.
	// Die Angabe sollte möglichst präzise erfol­gen.
	// Außer­dem soll­ten Abkür­zun­gen vermie­den werden. Für die höchste demo­kra­ti­sche Instanz in der Kommune sollte immer der Begriff “Parla­ment” verwen­det werden, nicht “Rat” oder “Haupt­aus­schuss”.
	Classification *string `json:"classification,omitempty"`

	// Allge­meine Website der Grup­pie­rung.
	Website *URL `json:"website,omitempty"`
}

// Person is eine natür­li­che Person, die in der parla­men­ta­ri­schen Arbeit tätig und insbe­son­dere Mitglied in einer Grup­pie­rung (oparl:Orga­ni­za­tion) ist, wird mit einem Objekt vom Typ oparl:Person abge­bil­det.
type Person struct {
	Common

	// Körper­schaft, zu der die Person gehört.
	Body *URL `json:"body,omitempty"`

	// Der voll­stän­dige Name der Person mit akade­mi­schem Grad und dem gebräuch­li­chen Vorna­men, wie er zur Anzeige durch den Client genutzt werden kann.
	Name *string `json:"name,omitempty"`

	// Fami­li­enname bzw. Nach­name.
	FamilyName *string `json:"familyName,omitempty"`

	// Vorname bzw. Tauf­name.
	GivenName *string `json:"givenName,omitempty"`

	// Akade­mi­sche Titel
	Title []string `json:"title,omitempty"`

	// Namens­zu­satz (z.B. jun. oder MdL.)
	Affix *string `json:"affix,omitempty"`

	// Anrede.
	FormOfAddress *string `json:"formOfAddress,omitempty"`

	// Geschlecht. Vorge­ge­bene Werte sind female und male, weitere werden durch die durch­ge­hend klein geschrie­bene engli­sche Bezeich­nung ange­ge­ben.
	// Für den Fall, dass das Geschlecht der Person unbe­kannt ist, sollte die Eigen­schaft nicht ausge­ge­ben werden.
	Gender *string `json:"gender,omitempty"`

	// E-Mail-Adres­sen der Person.
	Email []string `json:"email,omitempty"`

	// Tele­fon­num­mern der Person.
	Phone []string `json:"phone,omitempty"`

	//
	Status []string `json:"status,omitempty"`

	// Mitglied­schaf­ten der Person in Grup­pie­run­gen, z. B. Gremien und Frak­tio­nen.
	// Es sollen sowohl aktu­elle als auch vergan­gene Mitglied­schaf­ten ange­ge­ben werden.
	Membership []Membership `json:"membership,omitempty"`

	// Refe­renz der Kontakt-Anschrift der Person.
	Location *URL `json:"location,omitempty"`

	// Kontakt-Anschrift der Person.
	// Wenn diese Eigen­schaft ausge­ge­ben wird, dann muss auch die Eigen­schaft location ausge­ge­ben werden und auf das glei­che Loca­tion-Objekt verwei­sen.
	// Dieses Feld sollte die eigent­li­che Ausga­be­form von location in OParl 1.0 werden. vgl. https://github.com/OParl/spec/issues/373.
	// Neu in OParl 1.1
	LocationObject *Location `json:"locationObject,omitempty"`

	// Kurzer Infor­ma­ti­ons­text zur Person. Eine Länge von weni­ger als 300 Zeichen ist empfoh­len
	Life *string `json:"life,omitempty"`

	// Angabe der Quelle, aus der die Infor­ma­tio­nen für life stam­men. Bei Angabe von life ist diese Eigen­schaft empfoh­len
	LifeSource *string `json:"lifeSource,omitempty"`
}

// Über Objekte dieses Typs wird die Mitglied­schaft von Perso­nen in Grup­pie­run­gen darge­stellt.
// Diese Mitglied­schaf­ten können zeit­lich begrenzt sein.
// Zudem kann abge­bil­det werden, dass eine Person eine bestimmte Rolle bzw. Posi­tion inner­halb der Grup­pie­rung inne­hat, beispiels­weise den Vorsitz einer Frak­tion.
type Membership struct {
	Common

	// Rück­re­fe­renz auf Person, welches nur dann ausge­ge­ben werden muss, wenn das Member­ship-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Person *URL `json:"person,omitempty"`

	// Die Grup­pie­rung, in der die Person Mitglied ist oder war.
	Organization *URL `json:"organization,omitempty"`

	// Rolle der Person für die Grup­pie­rung.
	// Kann genutzt werden, um verschie­dene Arten von Mitglied­schaf­ten zum Beispiel in Gremien zu unter­schei­den.
	Role *string `json:"role,omitempty"`

	// Gibt an, ob die Person in der Grup­pie­rung stimm­be­rech­tig­tes Mitglied ist.
	VotingRight *bool `json:"votingRight,omitempty"`

	// Datum, an dem die Mitglied­schaft beginnt.
	StartDate *string `json:"startDate,omitempty"`

	// Datum, an dem die Mitglied­schaft endet.
	EndDate *string `json:"endDate,omitempty,omitempty"`

	// Die Grup­pie­rung, für die die Person in der unter organization ange­ge­be­nen Orga­ni­sa­tion sitzt.
	// Beispiel: Mitglied­schaft als Vertre­ter einer Rats­frak­tion, einer Grup­pie­rung oder einer exter­nen Orga­ni­sa­tion.
	OnBehalfOf *URL `json:"onBehalfOf,omitempty"`
}

// Meeting ist die Versamm­lung einer oder mehre­rer Grup­pie­run­gen (oparl:Orga­ni­za­tion) zu einem bestimm­ten Zeit­punkt an einem bestimm­ten Ort.
//
// Die gela­de­nen Teil­neh­mer der Sitzung sind jeweils als Objekte vom Typ oparl:Person, die in entspre­chen­der Form refe­ren­ziert werden.
// Verschie­dene Dateien (Einla­dung, Ergeb­nis- und Wort­pro­to­koll, sons­tige Anla­gen) können refe­ren­ziert werden.
// Die Inhalte einer Sitzung werden durch Tages­ord­nungs­punkte (oparl:Agen­daItem) abge­bil­det.
type Meeting struct {
	Common

	// Name der Sitzung.
	Name *string `json:"name,omitempty"`

	// Datum und Uhrzeit des Anfangs­zeit­punkts der Sitzung.
	// Bei einer zukünf­ti­gen Sitzung ist dies der geplante Zeit­punkt, bei einer statt­ge­fun­de­nen kann es der tatsäch­li­che Start­zeit­punkt sein.
	Start *time.Time `json:"start,omitempty"`

	// Endzeit­punkt der Sitzung als Datum/Uhrzeit.
	// Bei einer zukünf­ti­gen Sitzung ist dies der geplante Zeit­punkt, bei einer statt­ge­fun­de­nen kann es der tatsäch­li­che Endzeit­punkt sein.
	End *time.Time `json:"end,omitempty"`

	// Wenn die Sitzung ausfällt, wird cancel­led auf true gesetzt.
	Cancelled *bool `json:"cancelled,omitempty"`

	// Sitzungs­ort.
	Location *Location `json:"location,omitempty"`

	// Grup­pie­run­gen, denen die Sitzung zuge­ord­net ist. Im Regel­fall wird hier eine Grup­pie­rung verknüpft sein, es kann jedoch auch gemein­same Sitzun­gen mehre­rer Grup­pie­run­gen geben. Das erste Element sollte dann das feder­füh­rende Gremium sein.
	Organization []URL `json:"organization,omitempty"`

	// Einla­dungs­do­ku­ment zur Sitzung.
	Invitation *File `json:"invitation,omitempty"`

	// Ergeb­nispro­to­koll zur Sitzung.
	// Diese Eigen­schaft kann selbst­ver­ständ­lich erst nach­dem Statt­fin­den der Sitzung vorkom­men.
	ResultsProtocol *File `json:"resultsProtocol,omitempty"`

	// Wort­pro­to­koll zur Sitzung.
	// Diese Eigen­schaft kann selbst­ver­ständ­lich erst nach dem Statt­fin­den der Sitzung vorkom­men.
	VerbatimProtocol *File `json:"verbatimProtocol,omitempty"`

	// Datei­an­hang zur Sitzung.
	// Hier­mit sind Dateien gemeint, die übli­cher­weise mit der Einla­dung zu einer Sitzung verteilt werden, und die nicht bereits über einzelne Tages­ord­nungs­punkte refe­ren­ziert sind.
	AuxiliaryFile []File `json:"auxiliaryFile,omitempty"`

	// Tages­ord­nungs­punkte der Sitzung. Die Reihen­folge ist rele­vant. Es kann Sitzun­gen ohne TOPs geben.
	AgendaItem []AgendaItem `json:"agendaItem,omitempty"`

	// Aktu­el­ler Status der Sitzung. Empfoh­len ist die Verwen­dung von terminiert (geplant), eingeladen (vor der Sitzung bis zur Frei­gabe des Proto­kolls) und durchgeführt (nach Frei­gabe des Proto­kolls).
	MeetingState *string `json:"meetingState,omitempty"`
}

// AgendaItem ist ein Bestand­teil von Sitzun­gen (oparl:Meeting).
// Jeder Tages­ord­nungs­punkt widmet sich inhalt­lich einem bestimm­ten Thema, wozu in der Regel auch die Bera­tung bestimm­ter Druck­sa­chen gehört.
//
// Die Bezie­hung zwischen einem Tages­ord­nungs­punkt und einer Druck­sa­che wird über ein Objekt vom Typ oparl:Consultation herge­stellt, das über die Eigen­schaft consultation refe­ren­ziert werden kann.
type AgendaItem struct {
	Common

	// Rück­re­fe­renz auf das Meeting, welches nur dann ausge­ge­ben werden muss, wenn das agen­daItem-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Meeting *URL `json:"meeting,omitempty"`

	// Glie­de­rungs-“Nummer” des Tages­ord­nungs­punk­tes.
	// Eine belie­bige Zeichen­kette, wie z. B. “10.”, “10.1”, “C”, “c)” o. ä. Die Reihen­folge wird nicht dadurch, sondern durch die Reihen­folge der TOPs im agendaItem-Attri­but von oparl:Meeting fest­ge­legt, sollte aller­dings zu dieser iden­tisch sein.
	Number *string `json:"number,omitempty"`

	// Die Posi­tion des Tages­ord­nungs­punkts in der Sitzung, wenn alle Tages­ord­nungs­punkte von 0 an durch­ge­hend nume­riert werden.
	// Diese Nummer entspricht der Posi­tion in Meeting:agendaItem
	// Neu in OParl 1.1
	Order int `json:"order"`

	// Das Thema des Tages­ord­nungs­punk­tes.
	Name *string `json:"name,omitempty"`

	// Kenn­zeich­net, ob der Tages­ord­nungs­punkt zur Behand­lung in öffent­li­cher Sitzung vorge­se­hen ist/war.
	// Es wird ein Wahr­heits­wert (true oder false) erwar­tet.
	Public *bool `json:"public,omitempty"`

	// Bera­tung, die diesem Tages­ord­nungs­punkt zuge­wie­sen ist.
	Consultation *URL `json:"consultation,omitempty"`

	// Kate­go­ri­sche Infor­ma­tion darüber, welches Ergeb­nis die Bera­tung des Tages­ord­nungs­punk­tes erbracht hat, in der Bedeu­tung etwa “Unver­än­dert beschlos­sen” oder “Geän­dert beschlos­sen”.
	Result *string `json:"result,omitempty"`

	// Falls in diesem Tages­ord­nungs­punkt ein Beschluss gefasst wurde, kann hier ein Text ange­ge­ben werden.
	// Das ist beson­ders dann in der Praxis rele­vant, wenn der gefasste Beschluss (z. B. durch Ände­rungs­an­trag) von der Beschluss­vor­lage abweicht.
	ResolutionText *string `json:"resolutionText,omitempty"`

	// Falls in diesem Tages­ord­nungs­punkt ein Beschluss gefasst wurde, kann hier eine Datei ange­ge­ben werden.
	// Das ist beson­ders dann in der Praxis rele­vant, wenn der gefasste Beschluss (z. B. durch Ände­rungs­an­trag) von der Beschluss­vor­lage abweicht.
	ResolutionFile *File `json:"resolutionFile,omitempty"`

	// Weitere Datei­an­hänge zum Tages­ord­nungs­punkt.
	AuxiliaryFile []File `json:"auxiliaryFile,omitempty"`

	// Datum und Uhrzeit des Anfangs­zeit­punkts des Tages­ord­nungs­punk­tes.
	// Bei zukünf­ti­gen Tages­ord­nungs­punk­ten ist dies der geplante Zeit­punkt, bei einem statt­ge­fun­de­nen kann es der tatsäch­li­che Start­zeit­punkt sein.
	Start *time.Time `json:"start,omitempty"`

	// Endzeit­punkt des Tages­ord­nungs­punk­tes als Datum/Uhrzeit.
	// Bei zukünf­ti­gen Tages­ord­nungs­punk­ten ist dies der geplante Zeit­punkt, bei einer statt­ge­fun­de­nen kann es der tatsäch­li­che Endzeit­punkt sein.
	End *time.Time `json:"end,omitempty"`
}

// File reprä­sen­tiert eine Datei, beispiels­weise eine PDF-Datei, ein RTF- oder ODF-Doku­ment, und hält Meta­da­ten zu der Datei sowie URLs zum Zugriff auf die Datei bereit.
//
// Objekte vom Typ oparl:File können unter ande­rem mit Druck­sa­chen (oparl:Paper) oder Sitzun­gen (oparl:Meeting) in Bezie­hung stehen.
// Dies wird durch die Eigen­schaft paper bzw. meeting ange­zeigt.
// Mehrere Objekte vom Typ oparl:File können mit einan­der in direk­ter Bezie­hung stehen, z.B. wenn sie den selben Inhalt in unter­schied­li­chen tech­ni­schen Forma­ten wieder­ge­ben. Hier­für werden die Eigen­schaf­ten masterFile bzw. derivativeFile einge­setzt. Das gezeigte Beispiel-Objekt reprä­sen­tiert eine PDF-Datei (zu erken­nen an der Eigen­schaft mimeType) und zeigt außer­dem über die Eigen­schaft masterFile an, von welcher ande­ren Datei es abge­lei­tet wurde. Umge­kehrt kann über die Eigen­schaft derivativeFile ange­zeigt werden, welche Ablei­tun­gen einer Datei exis­tie­ren.
type File struct {
	Common

	// Ein zur Anzeige für Endnut­zer bestimm­ter Name für dieses Objekt.
	// Leer­zei­chen dürfen enthal­ten sein, Datei-Endun­gen wie “.pdf” soll­ten nicht enthal­ten sein.
	Name *string `json:"name,omitempty"`

	// Dateiname, unter dem die Datei in einem Datei­sys­tem gespei­chert werden kann.
	// Beispiel: “eineda­tei.pdf”. Da der Name den komplet­ten Unicode-Zeichen­um­fang nutzen kann, soll­ten Clients ggfs. selbst dafür sorgen, diesen beim Spei­chern in ein Datei­sys­tem den loka­len Erfor­der­nis­sen anzu­pas­sen.
	FileName *string `json:"fileName,omitempty"`

	// MIME-Type der Datei.
	MimeType *string `json:"mimeType,omitempty"`

	// Datum, welches als Start­punkt für Fris­ten u.ä. verwen­det ist.
	Date *string `json:"date,omitempty"`

	// Größe der Datei in Bytes.
	Size *int `json:"size,omitempty"`

	// Reine Text-Wieder­gabe des Datei­in­halts, sofern dieser in Text­form wieder­ge­ge­ben werden kann.
	Text *string `json:"text,omitempty"`

	// SHA1-Prüf­summe des Datei­in­halts in Hexa­de­zi­mal-Schreib­weise.
	// Sollte nicht mehr verwen­det werden, da sha1 als unsi­cher gilt.
	// Statt­des­sen sollte sha512checksum verwen­det werden.
	SHA1Checksum *string `json:"sha1Checksum,omitempty"`

	// SHA512-Prüf­summe des Datei­in­halts in Hexa­de­zi­mal-Schreib­weise.
	SHA256Checksum *string `json:"sha256Checksum,omitempty"`

	// URL zum allge­mei­nen Zugriff auf die Datei.
	AccessURL URL `json:"accessUrl"`

	// Externe URL, welche eine zusätz­li­che Zugriffs­mög­lich­keit bietet.
	// Beispiel: YouTube-Video.
	ExternalServiceURL *URL `json:"externalServiceUrl,omitempty"`

	// URL zum Down­load der Datei.
	DownloadURL *URL `json:"downloadUrl,omitempty"`

	// Datei, von der das aktu­elle Objekt abge­lei­tet wurde.
	MasterFile *URL `json:"masterFile,omitempty"`

	// Dateien, die von dem aktu­el­len Objekt abge­lei­tet wurden.
	DerivativeFile []URL `json:"derivativeFile,omitempty"`

	// Lizenz, unter der die Datei ange­bo­ten wird.
	// Wenn diese Eigen­schaft nicht verwen­det wird, ist der Wert von license bezie­hungs­weise die Lizenz eines über­ge­ord­ne­ten Objek­tes maßgeb­lich.
	FileLicense *string `json:"fileLicense,omitempty"`

	// Rück­re­fe­ren­zen auf Meeting-Objekte.
	// Wird nur dann ausge­ge­ben, wenn das File-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Meeting []URL `json:"meeting,omitempty"`

	// Rück­re­fe­ren­zen auf Agen­daItem-Objekte.
	// Wird nur dann ausge­ge­ben, wenn das File-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	AgendaItem *URL `json:"agendaItem,omitempty"`

	// Rück­re­fe­ren­zen auf Paper-Objekte.
	// Wird nur dann ausge­ge­ben, wenn das File-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Paper *URL `json:"paper,omitempty"`
}

// Paper dient der Abbil­dung von Druck­sa­chen in der parla­men­ta­ri­schen Arbeit, wie zum Beispiel Anfra­gen, Anträ­gen und Beschluss­vor­la­gen.
// Druck­sa­chen werden in Form einer Bera­tung (oparl:Consul­ta­tion) im Rahmen eines Tages­ord­nungs­punkts (oparl:Agen­daItem) einer Sitzung (oparl:Meeting) behan­delt.
//
// Druck­sa­chen spie­len in der schrift­li­chen wie münd­li­chen Kommu­ni­ka­tion eine beson­dere Rolle, da in vielen Texten auf bestimmte Druck­sa­chen Bezug genom­men wird.
// Hier­bei kommen in parla­men­ta­ri­schen Infor­ma­ti­ons­sys­te­men in der Regel unver­än­der­li­che Kennun­gen der Druck­sa­chen zum Einsatz.
type Paper struct {
	Common

	// Körper­schaft, zu der die Druck­sa­che gehört.
	Body *URL `json:"body,omitempty"`

	// Körper­schaft, zu der die Druck­sa­che gehört.
	Name *string `json:"name,omitempty"`

	// Kennung bzw. Akten­zei­chen der Druck­sa­che, mit der sie in der parla­men­ta­ri­schen Arbeit eindeu­tig refe­ren­ziert werden kann.
	Reference *string `json:"reference,omitempty"`

	// Datum, welches als Start­punkt für Fris­ten u.ä. verwen­det ist.
	Date *string `json:"date,omitempty"`

	// Art der Druck­sa­che, z. B. Beant­wor­tung einer Anfrage.
	PaperType *string `json:"paperType,omitempty"`

	// Inhalt­lich verwandte Druck­sa­chen.
	RelatedPaper []URL `json:"relatedPaper,omitempty"`

	// Über­ge­ord­nete Druck­sa­chen.
	SuperordinatedPaper []URL `json:"superordinatedPaper,omitempty"`

	// Unter­ge­ord­nete Druck­sa­chen.
	SubordinatedPaper []URL `json:"subordinatedPaper,omitempty"`

	// Die Haupt­da­tei zu dieser Druck­sa­che.
	// Beispiel: Die Druck­sa­che reprä­sen­tiert eine Beschluss­vor­lage und die Haupt­da­tei enthält den Text der Beschluss­vor­lage.
	// Sollte keine eindeu­tige Haupt­da­tei vorhan­den sein, wird diese Eigen­schaft nicht ausge­ge­ben.
	MainFile *File `json:"mainFile,omitempty"`

	// Alle weite­ren Dateien zur Druck­sa­che ausge­nom­men der gege­be­nen­falls in mainFile ange­ge­be­nen.
	AuxiliaryFile []File `json:"auxiliaryFile,omitempty"`

	// Sofern die Druck­sa­che einen inhalt­li­chen Orts­be­zug hat, beschreibt diese Eigen­schaft den Ort in Text­form und/oder in Form von Geoda­ten.
	Location []Location `json:"location,omitempty"`

	// Urhe­ber der Druck­sa­che, falls der Urhe­ber eine Person ist. Es können auch mehrere Perso­nen ange­ge­ben werden.
	OriginatorPerson []URL `json:"originatorPerson,omitempty"`

	// Urhe­ber der Druck­sa­che, falls der Urhe­ber eine Grup­pie­rung ist. Es können auch mehrere Grup­pie­run­gen ange­ge­ben werden.
	OriginatorOrganization []URL `json:"originatorOrganization,omitempty"`

	// Bera­tun­gen der Druck­sa­che.
	Consultation []Consultation `json:"consultation,omitempty"`

	// Feder­füh­rung. Amt oder Abtei­lung, für die Inhalte oder Beant­wor­tung der Druck­sa­che verant­wort­lich.
	UnderDirectionOf []URL `json:"underDirectionOf,omitempty"`
}

// Consultation dient dazu, die Bera­tung einer Druck­sa­che (oparl:Paper) in einer Sitzung abzu­bil­den.
// Dabei ist es nicht entschei­dend, ob diese Bera­tung in der Vergan­gen­heit statt­ge­fun­den hat oder diese für die Zukunft geplant ist.
// Die Gesamt­heit aller Objekte des Typs oparl:Consultation zu einer bestimm­ten Druck­sa­che bildet das ab, was in der Praxis als “Bera­tungs­folge” der Druck­sa­che bezeich­net wird.
type Consultation struct {
	Common

	// Refe­renz auf den Tages­ord­nungs­punkt, unter dem die Druck­sa­che bera­ten wird, welcher nur dann ausge­ge­ben werden muss, wenn das Consul­ta­tion-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	AgendaItem *URL `json:"agendaItem,omitempty"`

	// Refe­renz auf die Sitzung, in der die Druck­sa­che bera­ten wird oder wurde, welche nur dann ausge­ge­ben werden muss, wenn das Consul­ta­tion-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Meeting *URL `json:"meeting,omitempty"`

	// Gremium, in dem die Druck­sa­che bera­ten wird. Hier kann auch eine mit Liste von Gremien ange­ge­ben werden (die verschie­de­nen oparl:Body und oparl:System ange­hö­ren können).
	// Die Liste ist dann geord­net.
	// Das erste Gremium der Liste ist feder­füh­rend.
	Organization []URL `json:"organization,omitempty"`

	// Drückt aus, ob bei dieser Bera­tung ein Beschluss zu der Druck­sa­che gefasst wird oder wurde (true) oder nicht (false).
	Authoritative *bool `json:"authoritative,omitempty"`

	// Rolle oder Funk­tion der Bera­tung. Zum Beispiel Anhö­rung, Entschei­dung, Kennt­nis­nahme, Vorbe­ra­tung usw.
	Role *string `json:"role,omitempty"`

	// Refe­renz auf das Paper, welche nur dann ausge­ge­ben werden muss, wenn das Consul­ta­tion-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Paper *URL `json:"paper,omitempty"`
}
