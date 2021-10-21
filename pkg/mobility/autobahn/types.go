package autobahn

type ResponseWebcams struct {
	Webcams []Webcam `json:"webcams"`
}
type ResponseRoads struct {
	Roads []Road `json:"roads"`
}

type Coordinate struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type Road string

type Webcam struct {
	Extent                   string        `json:"extent"`
	Identifier               string        `json:"identifier"`
	RouteRecommendation      []interface{} `json:"routeRecommendation"`
	Coordinate               Coordinate    `json:"coordinate"`
	Footer                   []string      `json:"footer"`
	Icon                     string        `json:"icon"`
	IsBlocked                string        `json:"isBlocked"`
	Description              []interface{} `json:"description"`
	Title                    string        `json:"title"`
	Operator                 string        `json:"operator"`
	Point                    string        `json:"point"`
	DisplayType              string        `json:"display_type"`
	LorryParkingFeatureIcons []interface{} `json:"lorryParkingFeatureIcons"`
	Future                   bool          `json:"future"`
	ImageURL                 string        `json:"imageurl"`
	Subtitle                 string        `json:"subtitle"`
	LinkURL                  string        `json:"linkurl"`
}
