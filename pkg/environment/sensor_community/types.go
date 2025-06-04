// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package sensor_community

type ResponseLabs []Lab

type SensorDataValues struct {
	Value     string `json:"value"`
	ValueType string `json:"value_type"`
	ID        int    `json:"id"`
}

type SensorType struct {
	ID           int    `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}

type SensorLocation struct {
	ID            int     `json:"id"`
	Indoor        int     `json:"indoor"`
	ExactLocation int     `json:"exact_location"`
	Latitude      float64 `json:"latitude,string"`
	Longitude     float64 `json:"longitude,string"`
	Altitude      float64 `json:"altitude,string"`
	Country       string  `json:"country"`
}

type SensorInfo struct {
	Pin        string     `json:"pin"`
	ID         int        `json:"id"`
	SensorType SensorType `json:"sensor_type"`
}

type Sensor struct {
	Sensor       SensorInfo         `json:"sensor"`
	Location     SensorLocation     `json:"location"`
	ID           int                `json:"id"`
	DataValues   []SensorDataValues `json:"sensordatavalues"`
	Timestamp    string             `json:"timestamp"`
	SamplingRate interface{}        `json:"sampling_rate"`
}

type Lab struct {
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Title         string  `json:"title"`
	Text          string  `json:"text"`
	MeetingsTitle string  `json:"meetings_title"`
	Meetings      []struct {
		Text string `json:"text"`
	} `json:"meetings"`
	ContactsTitle string `json:"contacts_title"`
	Contacts      []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"contacts"`
	Website       string `json:"website,omitempty"`
	Twitter       string `json:"twitter,omitempty"`
	Github        string `json:"github,omitempty"`
	Meetup        string `json:"meetup,omitempty"`
	FacebookPage  string `json:"facebook_page,omitempty"`
	Instagram     string `json:"instagram,omitempty"`
	Telegram      string `json:"telegram,omitempty"`
	Facebook      string `json:"facebook,omitempty"`
	FacebookGroup string `json:"facebook_group,omitempty"`
	TelegramGroup string `json:"telegram_group,omitempty"`
	Mastodon      string `json:"mastodon,omitempty"`
}
