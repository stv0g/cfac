package autobahn_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stv0g/cfac/pkg/mobility/autobahn"
)

func TestGetRoads(t *testing.T) {
	roads, err := autobahn.GetRoads()
	if err != nil {
		t.Errorf("Failed to fetch roads: %s\n", err)
		t.FailNow()
	}

	fmt.Printf("Roads: %+v", roads)
}

func TestGetWebcams(t *testing.T) {
	roads, err := autobahn.GetRoads()
	if err != nil {
		t.Errorf("Failed to fetch roads: %s\n", err)
		t.FailNow()
	}

	all_webcams := []autobahn.Webcam{}

	for _, road := range roads {
		if strings.Contains(string(road), "/") {
			continue
		}

		webcams, err := autobahn.GetWebcams(road)
		if err != nil {
			t.Errorf("Failed to fetch webcams: %s\n", err)
			t.FailNow()
		}

		all_webcams = append(all_webcams, webcams...)
	}

	b, err := json.MarshalIndent(&all_webcams, "", " ")
	if err != nil {
		t.Fail()
	}

	fmt.Printf("Webcams: %s", string(b))
}
