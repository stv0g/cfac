package cccac

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
	Time   uint   `json:"time"`
	Type   string `json:"type"`
}

const (
	UrlApi       = "https://status.aachen.ccc.de/api/v0"
	UrlApiStatus = UrlApi + "/status/current?public"
)

func FetchStatus() (Status, error) {
	resp, err := http.Get(UrlApi)
	if err != nil {
		return Status{}, err
	}

	dec := json.NewDecoder(resp.Body)

	var sts struct {
		Changed Status `json:"changed"`
	}

	err = dec.Decode(&sts)
	if err != nil {
		return Status{}, err
	}

	return sts.Changed, nil
}
