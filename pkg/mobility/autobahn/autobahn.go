package autobahn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	UrlApi = "https://verkehr.autobahn.de/o/autobahn/"
)

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

func GetRoads() ([]Road, error) {
	resp, err := http.Get(UrlApi)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respRoads ResponseRoads
	if err := json.Unmarshal(body, &respRoads); err != nil {
		return nil, err
	}

	return respRoads.Roads, nil
}

func GetWebcams(r Road) ([]Webcam, error) {
	resp, err := http.Get(UrlApi + "/" + path.Join(string(r), "services", "webcam"))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respWebcams ResponseWebcams
	if err := json.Unmarshal(body, &respWebcams); err != nil {
		return nil, err
	}

	return respWebcams.Webcams, nil
}
