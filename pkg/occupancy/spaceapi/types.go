package spaceapi

type Directory map[string]string

type Space struct {
	API string `json:"api"`

	// The versions your SpaceAPI endpoint supports
	APICompatibility []string `json:"api_compatibility"`

	// The name of your space
	Space string `json:"space"`

	// URL to your space logo
	Logo string `json:"logo"`

	// URL to your space website
	URL string `json:"url"`

	// Position data such as a postal address or geographic coordinates
	Location Location `json:"location"`

	// URL(s) of webcams in your space
	Cam []string `json:"cam"`

	// Contact information about your space.
	Contact Contact `json:"contact"`

	IssueReportChannels []string `json:"issue_report_channels"`

	// A collection of status-related data: actual open/closed status, icons, last change timestamp etc.
	State State `json:"state"`

	// Your project sites (links to GitHub, wikis or wherever your projects are hosted)
	Projects []string `json:"projects"`

	// A flag indicating if the hackerspace uses SpaceFED, a federated login scheme so that visiting hackers can use the space WiFi with their home space credentials.
	SpaceFed struct {
		SpaceNet  bool `json:"spacenet"`
		SpaceSAML bool `json:"spacesaml"`
	} `json:"spacefed"`

	// Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'
	Events []Event `json:"events"`

	// Arbitrary links that you'd like to share
	Links []Link `json:"link"`

	// A list of the different membership plans your hackerspace might have.
	MembershipPlans []MembershipPlan `json:"membership_plans"`
}

// A list of the different membership plans your hackerspace might have.
// Set the value according to your billing process.
// For example, if your membership fee is 10€ per month, but you bill it yearly (aka. the member pays the fee once per year), set the amount to 120 an the billing_interval to yearly.
type MembershipPlan struct {
	Name            string  `json:"name"`
	Value           float64 `json:"value"`
	Currency        string  `json:"currency"`
	BillingInterval string  `json:"billing_interval"`
	Description     string  `json:"description"`
}

type Link struct {
	// The link name.
	Name string `json:"name"`

	// An extra field for a more detailed description of the link.
	Description string `json:"description"`

	// The URL.
	URL string `json:"url"`
}

// Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'
type Event struct {
	// Name or other identity of the subject (e.g. J. Random Hacker, fridge, 3D printer, …)
	Name string `json:"name"`

	// Action (e.g. check-in, check-out, finish-print, …). Define your own actions and use them consistently, canonical actions are not (yet) specified
	Type string `json:"type"`

	// Unix timestamp when the event occurred
	Timestamp uint64 `json:"timestamp"`

	// A custom text field to give more information about the event
	Extra string `json:"extra"`
}

// Contact information about your space.
type Contact struct {
	// E-mail address for contacting your space. If this is a mailing list consider to use the contact/ml field.
	Email string `json:"email"`

	// A separate email address for issue reports. This value can be Base64-encoded.
	IssueEmail string `json:"issue_email"`

	// URL of the IRC channel, in the form irc://example.org/#channelname
	IRC string `json:"irc"`

	// The e-mail address of your mailing list. If you use Google Groups then the e-mail looks like your-group@googlegroups.com.
	ML string `json:"ml"`

	// A public Jabber/XMPP multi-user chatroom in the form chatroom@conference.example.net
	XMPP string `json:"xmpp"`

	// Twitter handle, with leading @
	Twitter string `json:"twitter"`

	// Mastodon username:
	// Example: @ordnung@chaos.social.
	Mastodon string `json:"mastodon"`

	// Facebook account URL.
	Facebook string `json:"facebook"`

	// Identi.ca or StatusNet account, in the form yourspace@example.org
	Identica string `json:"identica"`

	// Foursquare ID, in the form 4d8a9114d85f3704eab301dc.
	Foursquare string `json:"foursquare"`

	// A URL to find information about the Space in the Gopherspace. Example: gopher://gopher.binary-kitchen.de
	Gopher string `json:"gopher"`

	// Matrix channel/community for the Hackerspace. Example: #spaceroom:example.org or +spacecommunity:example.org
	Matrix string `json:"matrix"`

	// URL to a Mumble server/channel, as specified in https://wiki.mumble.info/wiki/Mumble_URL.
	// Example: mumble://mumble.example.org/spaceroom?version=1.2.0
	Mumble string `json:"mumble"`

	// URI for Voice-over-IP via SIP.
	// Example: sip:yourspace@sip.example.org
	SIP string `json:"sip"`

	// Phone number, including country code with a leading plus sign.
	// Example: +1 800 555 4567
	Phone string `json:"phone"`
}

// A collection of status-related data: actual open/closed status, icons, last change timestamp etc.
type State struct {
	// Icons that show the status graphically
	Icon struct {
		// The URL to your customized space logo showing an open space
		Open string `json:"open"`

		// The URL to your customized space logo showing a closed space
		Closed string `json:"closed"`
	} `json:"icon"`

	// A flag which indicates whether the space is currently open or closed. The state 'undefined' can be achieved by omitting this field. A missing 'open' property carries the semantics of a temporary unavailability of the state, whereas the absence of the 'state' property itself means the state is generally not implemented for this space. This field is also allowed to explicitly have the value `null` for backwards compatibility with older schema versions, but this is deprecated and will be removed in a future version.
	Open bool `json:"open"`

	// The Unix timestamp when the space status changed most recently
	Lastchange int64 `json:"lastchange"`

	// The person who lastly changed the state e.g. opened or closed the space.
	TriggerPerson string `json:"trigger_person"`

	// An additional free-form string, could be something like 'open for public', 'members only' or whatever you want it to be
	Message string `json:"message"`
}

// Position data such as a postal address or geographic coordinates
type Location struct {
	// The postal address of your space (street, block, housenumber, zip code, city, whatever you usually need in your country, and the country itself).
	// Example: Netzladen e.V., Breite Straße 74, 53111 Bonn, Germany
	Address string `json:"address"`

	// Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.
	Lat float64 `json:"lat"`

	// Longitude of your space location, in degree with decimal places. Use positive values for locations east of Greenwich, and negative values for locations west of Greenwich.
	Lon float64 `json:"lon"`

	// The timezone the space is located in. It should be formatted according to the TZ database location names.
	// See: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	Timezone string `json:"timezone"`
}

// Data of various sensors in your space (e.g. temperature, humidity, amount of Club-Mate left, …).
// The only canonical property is the temp property, additional sensor types may be defined by you.
// In this case, you are requested to share your definition for inclusion in this specification.
type Sensors struct {
	// Temperature sensor. To convert from one unit of temperature to another consider Wikipedia.
	Temperature UnitSensorValue `json:"temperature"`

	// Sensor type to indicate if a certain door is locked.
	DoorLocked BooleanSensorValue `json:"door_locked"`

	// Barometer sensor
	Barometer UnitSensorValue `json:"barometer"`

	// Compound radiation sensor.
	// See: https://sites.google.com/site/diygeigercounter/gm-tubes-supported
	Radiation interface{} `json:"radiation"`

	// Humidity sensor
	Humidity UnitSensorValue `json:"humidity"`

	// How much Mate and beer is in your fridge?
	BeverageSupply UnitSensorValue `json:"beverage_supply"`

	// The power consumption of a specific device or of your whole space.
	PowerConsumption UnitSensorValue `json:"power_consumption"`

	// Your wind sensor.
	Wind UnitSensorValue `json:"wind"`

	// This sensor type is to specify the currently active ethernet or wireless network devices.
	// You can create different instances for each network type.
	NetworkConnections UnitSensorValue `json:"network_connections"`

	// How rich is your hackerspace?
	AccountBalance UnitSensorValue `json:"account_balance"`

	// Specify the number of space members.
	TotalMemberCount UnitSensorValue `json:"total_member_count"`

	// Specify the number of people that are currently in your space. Optionally you can define a list of names.
	PeopleNowPresent UnitSensorValue `json:"people_now_present"`

	// The current network traffic, in bits/second or packets/second (or both)
	NetworkTraffic interface{} `json:"network_traffic"`
}

type SensorValue struct {
	// The sensor value
	Value interface{} `json:"value"`

	// The location of your sensor such as front door, chill room or lab.
	Location string `json:"location"`

	// This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.
	Name string `json:"name"`

	// An extra field that you can use to attach some additional information to this sensor instance.
	Description string `json:"description"`
}

type UnitSensorValue struct {
	SensorValue

	// The unit of the sensor value.
	// Valid values: °C | °F | K | °De | °N | °R | °Ré | °Rø
	Unit string `json:"unit"`
}

type BooleanSensorValue struct {
	// The sensor value
	Value bool `json:"value"`

	// The location of your sensor such as front door, chill room or lab.
	Location string `json:"location"`

	// This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.
	Name string `json:"name"`

	// An extra field that you can use to attach some additional information to this sensor instance.
	Description string `json:"description"`
}
