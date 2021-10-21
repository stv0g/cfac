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
