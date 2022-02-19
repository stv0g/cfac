package oparl

// https://dev.oparl.org/spezifikation

import (
	"time"

	"github.com/stv0g/cfac/pkg/std/ietf/geojson"
)

type URL string

type Common struct {
	ID       URL       `json:"id"`
	Type     string    `json:"type"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Keyword  []string  `json:"keyword"`
	Deleted  bool      `json:"deleted"`

	// Lizenz, unter der durch diese API abruf­ba­ren Daten stehen, sofern nicht am einzel­nen Objekt anders ange­ge­ben.
	License URL `json:"license"`
	Web     URL `json:"web"`
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
	Name string `json:"name"`

	// E-Mail-Adresse für Anfra­gen zur OParl-API.
	// Die Angabe einer E-Mail-Adresse dient sowohl NutzerIn­nen wie auch Entwick­le­rin­nen von Clients zur Kontakt­auf­nahme mit dem Betrei­ber.
	ContactEmail string `json:"contactEmail"`

	// Name der Ansprech­part­ne­rin bzw. des Ansprech­part­ners oder der Abtei­lung, die über die in contactEmail ange­ge­bene Adresse erreicht werden kann.
	ContactName string `json:"contactName"`

	// URL der Website des parla­men­ta­ri­schen Infor­ma­ti­ons­sys­tems.
	Website URL `json:"website"`

	// URL der Website des Soft­wa­rean­bie­ters, von dem die OParl-Server-Soft­ware stammt.
	Vendor URL `json:"vendor"`

	// URL zu Infor­ma­tio­nen über die auf dem System genutzte OParl-Server-Soft­ware.
	Product URL `json:"product"`

	// Dient der Angabe von System-Objek­ten mit ande­ren OParl-Versio­nen.
	OtherOparlVersions []string `json:"otherOparlVersions"`
}

// Der Objekt­typ oparl:Body dient dazu, eine Körper­schaft zu reprä­sen­tie­ren.
// Eine Körper­schaft ist in den meis­ten Fällen eine Gemeinde, eine Stadt oder ein Land­kreis.
// In der Regel sind auf einem OParl-Server Daten von genau einer Körper­schaft gespei­chert und es wird daher auch nur ein Body-Objekt ausge­ge­ben.
// Sind auf dem Server jedoch Daten von mehre­ren Körper­schaf­ten gespei­chert, muss für jede Körper­schaft ein eige­nes Body-Objekt ausge­ge­ben werden.
type Body struct {
	Common

	// System, zu dem dieses Objekt gehört.
	System URL `json:"system"`

	// Dient der Angabe einer Kontakt-E-Mail-Adresse.
	// Die Adresse soll die Kontakt­auf­nahme zu einer für die Körper­schaft und idea­ler­weise das parla­men­ta­ri­sche Infor­ma­ti­ons­sys­tem zustän­di­gen Stelle ermög­li­chen.
	ContactEmail string `json:"contactEmail"`

	// Name oder Bezeich­nung der mit contactEmail erreich­ba­ren Stelle.
	ContactName string `json:"contactName"`

	// Der acht­stel­lige Amtli­che Gemein­de­schlüs­sel.
	AGS string `json:"ags"`

	// Der zwölf­stel­lige Regio­nal­schlüs­sel.
	RGS string `json:"rgs"`

	// Dient der Angabe zusätz­li­cher URLs, die dieselbe Körper­schaft reprä­sen­tie­ren
	// Hier können beispiels­weise der entspre­chende Eintrag der gemein­sa­men Norm­da­tei der Deut­schen Natio­nal­bi­blio­thek, der DBPe­dia oder der Wiki­pe­dia ange­ge­ben werden.
	// Body- oder System-Objekte mit ande­ren OParl-Versio­nen dürfen nicht Teil der Liste sein.
	Equivalent []URL `json:"equivalent"`

	// Kurzer Name der Körper­schaft.
	ShortName string `json:"shortName"`

	// Der offi­zi­elle lange Name der Körper­schaft.
	Name string `json:"name"`

	// Allge­meine Website der Körper­schaft.
	Website URL `json:"website"`

	// Zeit­punkt, seit dem die unter license ange­ge­bene Lizenz gilt.
	// Vorsicht bei Ände­run­gen der Lizenz die zu restrik­ti­veren Bedin­gun­gen führen!
	LicenseValidSince time.Time `json:"licenseValidSince"`

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
	Location *Location `json:"location"`

	// Art der Körper­schaft.
	Classification string `json:"classification"`

	// Zeit­punkt, ab dem OParl für dieses Body bereit­ge­stellt wurde.
	// Dies hilft, um die Daten­qua­li­tät einzu­schät­zen, denn erst ab der Einrich­tung für OParl kann sicher­ge­stellt werden, dass sämt­li­che Werte korrekt in der Origi­nal-Quelle vorlie­gen.
	OParlSince time.Time `json:"oparlSince`
}

type LegislativeTerm struct {
	Common

	// Rück­re­fe­renz auf die Körper­schaft, welche nur dann ausge­ge­ben werden muss, wenn das Legis­la­ti­veTerm-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Body URL `json:"body"`

	// Nutzer­freund­li­che Bezeich­nung der Wahl­pe­ri­ode.
	Name string `json:"name"`

	// Der erste Tag der Wahl­pe­ri­ode.
	StartDate string `json:"startDate"`

	// Der letzte Tag der Wahl­pe­ri­ode.
	EndDate string `json:"endDate"`
}

// TODO
type Location struct {
	Common

	Description string          `json:"description"`
	GeoJSON     geojson.Feature `json:"geojson"`
}

type Organization struct {
	Common

	// Körper­schaft, zu der diese Grup­pie­rung gehört.
	Body URL `json:"body"`

	// Exter­ner OParl Body, der dieser Orga­ni­sa­tion entspricht.
	// Diese Eigen­schaft ist dafür gedacht auf even­tu­elle konkre­tere OParl-Schnitt­stel­len zu verwei­sen.
	// Ein Beispiel hier­für wäre eine Stadt, die sowohl ein über­grei­fen­des parla­men­ta­ri­sches Infor­ma­ti­ons­sys­tem, als auch bezirkss­pe­zi­fi­sche Systeme hat.
	ExternalBody URL `json:"externalBody"`

	// Offi­zi­elle (lange) Form des Namens der Grup­pie­rung.
	Name string `json:"name"`

	// Der Name der Grup­pie­rung als Kurz­form.
	ShortName string `json:"shortName"`

	// Grün­dungs­da­tum der Grup­pie­rung.
	// Kann z. B. das Datum der konsti­tu­ie­ren­den Sitzung sein.
	StartDate string `json:"startDate"`

	// Datum des letz­ten Tages der Exis­tenz der Grup­pie­rung.
	EndDate string `json:"endDate"`

	// Grobe Kate­go­ri­sie­rung der Grup­pie­rung.
	// Mögli­che Werte sind “Gremium”, “Partei”, “Frak­tion”, “Verwal­tungs­be­reich”, “exter­nes Gremium”, “Insti­tu­tion” und “Sons­ti­ges”.
	OrganizationType string `json:"organizationType"`

	// Ort, an dem die Orga­ni­sa­tion behei­ma­tet ist
	Location *Location `json:"location"`

	// Posi­tio­nen, die für diese Grup­pie­rung vorge­se­hen sind.
	Post []string `json:"post"`

	// URL auf eine externe Objekt­liste mit den Bera­tun­gen dieser Grup­pie­rung.
	// Invers zur Eigen­schaft organization der Klasse oparl:Consultation
	Consultation URL `json:"consultation"`

	// URL auf eine externe Objekt­liste mit den Sitzun­gen dieser Grup­pie­rung.
	// Invers zur Eigen­schaft organization der Klasse oparl:Meeting
	Meeting URL `json:"meeting"`

	// URL einer even­tu­el­len über­ge­ord­ne­ten Grup­pie­rung.
	SubOrganizationOf URL `json:"subOrganizationOf"`

	// Mitglied­schaf­ten dieser Grup­pie­rung.
	Membership []URL `json:"membership"`

	// Die Art der Grup­pie­rung.
	// In Frage kommen z.B. “Parla­ment”, “Ausschuss”, “Beirat”, “Projekt­bei­rat”, “Kommis­sion”, “AG”, “Verwal­tungs­rat”, “Frak­tion” oder “Partei”.
	// Die Angabe sollte möglichst präzise erfol­gen.
	// Außer­dem soll­ten Abkür­zun­gen vermie­den werden. Für die höchste demo­kra­ti­sche Instanz in der Kommune sollte immer der Begriff “Parla­ment” verwen­det werden, nicht “Rat” oder “Haupt­aus­schuss”.
	Classification string `json:"classification"`

	// Allge­meine Website der Grup­pie­rung.
	Website URL `json:"website"`
}

type Person struct {
	Common

	// Körper­schaft, zu der die Person gehört.
	Body URL `json:"body"`

	// Der voll­stän­dige Name der Person mit akade­mi­schem Grad und dem gebräuch­li­chen Vorna­men, wie er zur Anzeige durch den Client genutzt werden kann.
	Name string `json:"name"`

	// Fami­li­enname bzw. Nach­name.
	FamilyName string `json:"familyName"`

	// Vorname bzw. Tauf­name.
	GivenName string `json:"givenName"`

	// Akade­mi­sche Titel
	Title []string `json:"title"`

	// Namens­zu­satz (z.B. jun. oder MdL.)
	Affix string `json:"affix`

	// Anrede.
	FormOfAddress string `json:"formOfAddress"`

	// Geschlecht. Vorge­ge­bene Werte sind female und male, weitere werden durch die durch­ge­hend klein geschrie­bene engli­sche Bezeich­nung ange­ge­ben.
	// Für den Fall, dass das Geschlecht der Person unbe­kannt ist, sollte die Eigen­schaft nicht ausge­ge­ben werden.
	Gender string `json:"gender"`

	// E-Mail-Adres­sen der Person.
	Email []string `json:"email"`

	// Tele­fon­num­mern der Person.
	Phone []string `json:"phone"`

	//
	Status []string `json:"status"`

	// Mitglied­schaf­ten der Person in Grup­pie­run­gen, z. B. Gremien und Frak­tio­nen.
	// Es sollen sowohl aktu­elle als auch vergan­gene Mitglied­schaf­ten ange­ge­ben werden.
	Membership []Membership `json:"membership"`

	// Refe­renz der Kontakt-Anschrift der Person.
	Location URL `json:"location"`

	// Kontakt-Anschrift der Person.
	// Wenn diese Eigen­schaft ausge­ge­ben wird, dann muss auch die Eigen­schaft location ausge­ge­ben werden und auf das glei­che Loca­tion-Objekt verwei­sen.
	// Dieses Feld sollte die eigent­li­che Ausga­be­form von location in OParl 1.0 werden. vgl. https://github.com/OParl/spec/issues/373.
	// Neu in OParl 1.1
	LocationObject *Location `json:"locationObject"`

	// Kurzer Infor­ma­ti­ons­text zur Person. Eine Länge von weni­ger als 300 Zeichen ist empfoh­len
	Life string `json:"life"`

	// Angabe der Quelle, aus der die Infor­ma­tio­nen für life stam­men. Bei Angabe von life ist diese Eigen­schaft empfoh­len
	LifeSource string `json:"lifeSource"`
}

type Membership struct {
	Common

	// Rück­re­fe­renz auf Person, welches nur dann ausge­ge­ben werden muss, wenn das Member­ship-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Person URL `json:"person"`

	// Die Grup­pie­rung, in der die Person Mitglied ist oder war.
	Organization URL `json:"organization"`

	// Rolle der Person für die Grup­pie­rung.
	// Kann genutzt werden, um verschie­dene Arten von Mitglied­schaf­ten zum Beispiel in Gremien zu unter­schei­den.
	Role string `json:"role"`

	// Gibt an, ob die Person in der Grup­pie­rung stimm­be­rech­tig­tes Mitglied ist.
	VotingRight bool `json:"votingRight"`

	// Datum, an dem die Mitglied­schaft beginnt.
	StartDate string `json:"startDate"`

	// Datum, an dem die Mitglied­schaft endet.
	EndDate string `json:"endDate,omitempty"`

	// Die Grup­pie­rung, für die die Person in der unter organization ange­ge­be­nen Orga­ni­sa­tion sitzt.
	// Beispiel: Mitglied­schaft als Vertre­ter einer Rats­frak­tion, einer Grup­pie­rung oder einer exter­nen Orga­ni­sa­tion.
	OnBehalfOf URL `json:"onBehalfOf"`
}

type Meeting struct {
	Common

	// Name der Sitzung.
	Name string `json:"name"`

	// Datum und Uhrzeit des Anfangs­zeit­punkts der Sitzung.
	// Bei einer zukünf­ti­gen Sitzung ist dies der geplante Zeit­punkt, bei einer statt­ge­fun­de­nen kann es der tatsäch­li­che Start­zeit­punkt sein.
	Start time.Time `json:"start"`

	// Endzeit­punkt der Sitzung als Datum/Uhrzeit.
	// Bei einer zukünf­ti­gen Sitzung ist dies der geplante Zeit­punkt, bei einer statt­ge­fun­de­nen kann es der tatsäch­li­che Endzeit­punkt sein.
	End time.Time `json:"end"`

	// Wenn die Sitzung ausfällt, wird cancel­led auf true gesetzt.
	Cancelled bool `json:"cancelled"`

	// Sitzungs­ort.
	Location *Location `json:"location"`

	// Grup­pie­run­gen, denen die Sitzung zuge­ord­net ist. Im Regel­fall wird hier eine Grup­pie­rung verknüpft sein, es kann jedoch auch gemein­same Sitzun­gen mehre­rer Grup­pie­run­gen geben. Das erste Element sollte dann das feder­füh­rende Gremium sein.
	Organization []URL `json:"organization"`

	// Einla­dungs­do­ku­ment zur Sitzung.
	Invitation *File `json:"invitation"`

	// Ergeb­nispro­to­koll zur Sitzung.
	// Diese Eigen­schaft kann selbst­ver­ständ­lich erst nach­dem Statt­fin­den der Sitzung vorkom­men.
	ResultsProtocol *File `json:"resultsProtocol"`

	// Wort­pro­to­koll zur Sitzung.
	// Diese Eigen­schaft kann selbst­ver­ständ­lich erst nach dem Statt­fin­den der Sitzung vorkom­men.
	VerbatimProtocol *File `json:"verbatimProtocol"`

	// Datei­an­hang zur Sitzung.
	// Hier­mit sind Dateien gemeint, die übli­cher­weise mit der Einla­dung zu einer Sitzung verteilt werden, und die nicht bereits über einzelne Tages­ord­nungs­punkte refe­ren­ziert sind.
	AuxiliaryFile []File `json:"auxiliaryFile"`

	// Tages­ord­nungs­punkte der Sitzung. Die Reihen­folge ist rele­vant. Es kann Sitzun­gen ohne TOPs geben.
	AgendaItem []AgendaItem `json:"agendaItem"`

	// Aktu­el­ler Status der Sitzung. Empfoh­len ist die Verwen­dung von terminiert (geplant), eingeladen (vor der Sitzung bis zur Frei­gabe des Proto­kolls) und durchgeführt (nach Frei­gabe des Proto­kolls).
	MeetingState string `json:"meetingState"`
}

type AgendaItem struct {
	Common

	// Rück­re­fe­renz auf das Meeting, welches nur dann ausge­ge­ben werden muss, wenn das agen­daItem-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Meeting URL `json:"meeting"`

	// Glie­de­rungs-“Nummer” des Tages­ord­nungs­punk­tes.
	// Eine belie­bige Zeichen­kette, wie z. B. “10.”, “10.1”, “C”, “c)” o. ä. Die Reihen­folge wird nicht dadurch, sondern durch die Reihen­folge der TOPs im agendaItem-Attri­but von oparl:Meeting fest­ge­legt, sollte aller­dings zu dieser iden­tisch sein.
	Number string `json:"number"`

	// Die Posi­tion des Tages­ord­nungs­punkts in der Sitzung, wenn alle Tages­ord­nungs­punkte von 0 an durch­ge­hend nume­riert werden.
	// Diese Nummer entspricht der Posi­tion in Meeting:agendaItem
	// Neu in OParl 1.1
	Order int `json:"order"`

	// Das Thema des Tages­ord­nungs­punk­tes.
	Name string `json:"name"`

	// Kenn­zeich­net, ob der Tages­ord­nungs­punkt zur Behand­lung in öffent­li­cher Sitzung vorge­se­hen ist/war.
	// Es wird ein Wahr­heits­wert (true oder false) erwar­tet.
	Public bool `json:"public"`

	// Bera­tung, die diesem Tages­ord­nungs­punkt zuge­wie­sen ist.
	Consultation URL `json:"consultation"`

	// Kate­go­ri­sche Infor­ma­tion darüber, welches Ergeb­nis die Bera­tung des Tages­ord­nungs­punk­tes erbracht hat, in der Bedeu­tung etwa “Unver­än­dert beschlos­sen” oder “Geän­dert beschlos­sen”.
	Result string `json:"result"`

	// Falls in diesem Tages­ord­nungs­punkt ein Beschluss gefasst wurde, kann hier ein Text ange­ge­ben werden.
	// Das ist beson­ders dann in der Praxis rele­vant, wenn der gefasste Beschluss (z. B. durch Ände­rungs­an­trag) von der Beschluss­vor­lage abweicht.
	ResolutionText string `json:"resolutionText"`

	// Falls in diesem Tages­ord­nungs­punkt ein Beschluss gefasst wurde, kann hier eine Datei ange­ge­ben werden.
	// Das ist beson­ders dann in der Praxis rele­vant, wenn der gefasste Beschluss (z. B. durch Ände­rungs­an­trag) von der Beschluss­vor­lage abweicht.
	ResolutionFile *File `json:"resolutionFile"`

	// Weitere Datei­an­hänge zum Tages­ord­nungs­punkt.
	AuxiliaryFile []File `json:"auxiliaryFile"`

	// Datum und Uhrzeit des Anfangs­zeit­punkts des Tages­ord­nungs­punk­tes.
	// Bei zukünf­ti­gen Tages­ord­nungs­punk­ten ist dies der geplante Zeit­punkt, bei einem statt­ge­fun­de­nen kann es der tatsäch­li­che Start­zeit­punkt sein.
	Start time.Time `json:"start"`

	// Endzeit­punkt des Tages­ord­nungs­punk­tes als Datum/Uhrzeit.
	// Bei zukünf­ti­gen Tages­ord­nungs­punk­ten ist dies der geplante Zeit­punkt, bei einer statt­ge­fun­de­nen kann es der tatsäch­li­che Endzeit­punkt sein.
	End time.Time `json:"end"`
}

type File struct {
	Common

	// Ein zur Anzeige für Endnut­zer bestimm­ter Name für dieses Objekt.
	// Leer­zei­chen dürfen enthal­ten sein, Datei-Endun­gen wie “.pdf” soll­ten nicht enthal­ten sein.
	Name string `json:"name"`

	// Dateiname, unter dem die Datei in einem Datei­sys­tem gespei­chert werden kann.
	// Beispiel: “eineda­tei.pdf”. Da der Name den komplet­ten Unicode-Zeichen­um­fang nutzen kann, soll­ten Clients ggfs. selbst dafür sorgen, diesen beim Spei­chern in ein Datei­sys­tem den loka­len Erfor­der­nis­sen anzu­pas­sen.
	FileName string `json:"fileName"`

	// MIME-Type der Datei.
	MimeType string `json:"mimeType"`

	// Datum, welches als Start­punkt für Fris­ten u.ä. verwen­det ist.
	Date string `json:"date"`

	// Größe der Datei in Bytes.
	Size int `json:"size"`

	// Reine Text-Wieder­gabe des Datei­in­halts, sofern dieser in Text­form wieder­ge­ge­ben werden kann.
	Text string `json:"text"`

	// SHA1-Prüf­summe des Datei­in­halts in Hexa­de­zi­mal-Schreib­weise.
	// Sollte nicht mehr verwen­det werden, da sha1 als unsi­cher gilt.
	// Statt­des­sen sollte sha512checksum verwen­det werden.
	SHA1Checksum string `json:"sha1Checksum"`

	// SHA512-Prüf­summe des Datei­in­halts in Hexa­de­zi­mal-Schreib­weise.
	SHA256Checksum string `json:"sha256Checksum"`

	// URL zum allge­mei­nen Zugriff auf die Datei.
	AccessURL URL `json:"accessUrl"`

	// Externe URL, welche eine zusätz­li­che Zugriffs­mög­lich­keit bietet.
	// Beispiel: YouTube-Video.
	ExternalServiceURL URL `json:"externalServiceUrl"`

	// URL zum Down­load der Datei.
	DownloadURL URL `json:"downloadUrl"`

	// Datei, von der das aktu­elle Objekt abge­lei­tet wurde.
	MasterFile URL `json:"masterFile"`

	// Dateien, die von dem aktu­el­len Objekt abge­lei­tet wurden.
	DerivativeFile []URL `json:"derivativeFile"`

	// Lizenz, unter der die Datei ange­bo­ten wird.
	// Wenn diese Eigen­schaft nicht verwen­det wird, ist der Wert von license bezie­hungs­weise die Lizenz eines über­ge­ord­ne­ten Objek­tes maßgeb­lich.
	FileLicense string `json:"fileLicense"`

	// Rück­re­fe­ren­zen auf Meeting-Objekte.
	// Wird nur dann ausge­ge­ben, wenn das File-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Meeting []URL `json:"meeting"`

	// Rück­re­fe­ren­zen auf Agen­daItem-Objekte.
	// Wird nur dann ausge­ge­ben, wenn das File-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	AgendaItem URL `json:"agendaItem"`

	// Rück­re­fe­ren­zen auf Paper-Objekte.
	// Wird nur dann ausge­ge­ben, wenn das File-Objekt nicht als einge­bet­te­tes Objekt aufge­ru­fen wird.
	Paper URL `json:"paper"`
}

// TODO
type Paper struct {
	Common

	//
	Body URL `json:"body"`

	//
	Name string `json:"name"`

	//
	Reference string `json:"reference"`

	//
	Date string `json:"date"`

	//
	PaperType string `json:"paperType"`

	//
	RelatedPaper []URL `json:"relatedPaper"`

	//
	MainFile *File `json:"mainFile"`

	//
	AuxiliaryFile []File `json:"auxiliaryFile"`

	//
	Location []Location `json:"location"`

	//
	OriginatorPerson []URL `json:"originatorPerson"`

	//
	OriginatorOrganization []URL `json:"originatorOrganization"`

	//
	Consultation []Consultation `json:"consultation"`

	//
	UnderDirectionOf []URL `json:"underDirectionOf"`
}

type Consultation struct {
	Common

	// Refe­renz auf den Tages­ord­nungs­punkt, unter dem die Druck­sa­che bera­ten wird, welcher nur dann ausge­ge­ben werden muss, wenn das Consul­ta­tion-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	AgendaItem URL `json:"agendaItem"`

	// Refe­renz auf die Sitzung, in der die Druck­sa­che bera­ten wird oder wurde, welche nur dann ausge­ge­ben werden muss, wenn das Consul­ta­tion-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Meeting URL `json:"meeting"`

	// Gremium, in dem die Druck­sa­che bera­ten wird. Hier kann auch eine mit Liste von Gremien ange­ge­ben werden (die verschie­de­nen oparl:Body und oparl:System ange­hö­ren können).
	// Die Liste ist dann geord­net.
	// Das erste Gremium der Liste ist feder­füh­rend.
	Organization []URL `json:"organization"`

	// Drückt aus, ob bei dieser Bera­tung ein Beschluss zu der Druck­sa­che gefasst wird oder wurde (true) oder nicht (false).
	Authoritative bool `json:"authoritative"`

	// Rolle oder Funk­tion der Bera­tung. Zum Beispiel Anhö­rung, Entschei­dung, Kennt­nis­nahme, Vorbe­ra­tung usw.
	Role string `json:"role"`

	// Refe­renz auf das Paper, welche nur dann ausge­ge­ben werden muss, wenn das Consul­ta­tion-Objekt einzeln abge­ru­fen wird, d.h. nicht Teil einer inter­nen Ausgabe ist.
	Paper URL `json:"paper"`
}
