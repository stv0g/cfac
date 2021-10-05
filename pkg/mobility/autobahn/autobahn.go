package autobahn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Url = "https://verkehr.autobahn.de/o/autobahn/"
)

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
	Imageurl                 string        `json:"imageurl"`
	Subtitle                 string        `json:"subtitle"`
	Linkurl                  string        `json:"linkurl"`
}

func GetRoads() ([]Road, error) {
	resp, err := http.Get(Url)
	if err != nil {
		return nil, err
	}

	var bodyJson struct {
		Roads []Road `json:"roads"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &bodyJson); err != nil {
		return nil, err
	}

	return bodyJson.Roads, nil
}

func GetWebcams(r Road) ([]Webcam, error) {
	resp, err := http.Get(Url + "/" + url.PathEscape(string(r)) + "/services/webcam")
	if err != nil {
		return nil, err
	}

	var bodyJson struct {
		Webcams []Webcam `json:"webcam"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &bodyJson); err != nil {
		return nil, err
	}

	return bodyJson.Webcams, nil
}
