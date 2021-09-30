package fitx

// https://www.fitx.de/fitnessstudio/38/workload
// https://services.fitx.de/survey/nps_courseplan_display
// https://services.fitx.de/survey/nps_studiodetail_v2

type StudioLocation struct {
	Street     string  `json:"street"`
	Zip        string  `json:"zip"`
	Place      string  `json:"place"`
	Region     string  `json:"region"`
	Phone      string  `json:"phone"`
	Directions string  `json:"directions"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	PlaceID    string  `json:"placeId"`
}

type StudioTour struct {
	Enabled  bool   `json:"enabled"`
	URL      string `json:"url"`
	EmbedURL string `json:"embedUrl"`
}

type StudioDates struct {
	Opening                 int         `json:"opening"`
	OpeningRegular          int         `json:"openingRegular"`
	OpeningCourses          int         `json:"openingCourses"`
	TemporarilyClosedReason interface{} `json:"temporarilyClosedReason"`
	TemporarilyClosedFrom   interface{} `json:"temporarilyClosedFrom"`
	TemporarilyClosedTo     interface{} `json:"temporarilyClosedTo"`
	Presale                 interface{} `json:"presale"`
}

type StudioWorkload struct {
	Term       string `json:"term"`
	Percentage int    `json:"percentage"`
}

type Studio struct {
	Distance                               interface{}    `json:"distance"`
	Name                                   string         `json:"name"`
	Alias                                  string         `json:"alias"`
	Identifier                             string         `json:"identifier"`
	BranchID                               int            `json:"branchId"`
	Location                               StudioLocation `json:"location"`
	VirtualTour                            StudioTour     `json:"virtualTour"`
	Notice                                 interface{}    `json:"notice"`
	Dates                                  StudioDates    `json:"dates"`
	Badge                                  interface{}    `json:"badge"`
	Promotions                             []string       `json:"promotions"`
	Status                                 int            `json:"status"`
	DisablePreregistration                 interface{}    `json:"disablePreregistration"`
	VisitorData                            [][]int        `json:"visitorData"`
	TrialworkoutEnabled                    bool           `json:"trialworkoutEnabled"`
	TrialworkoutRestricted                 bool           `json:"trialworkoutRestricted"`
	PreregistrationAgeRestricted           bool           `json:"preregistrationAgeRestricted"`
	PreregistrationAccountHolderRestricted bool           `json:"preregistrationAccountHolderRestricted"`
	Images                                 []string       `json:"images"`
	ListText                               string         `json:"listText"`
	ListTextColor                          interface{}    `json:"listTextColor"`
	GoogleRating                           float64        `json:"googleRating"`
	GoogleRatingCount                      int            `json:"googleRatingCount"`
	Workload                               StudioWorkload `json:"workload"`
}
