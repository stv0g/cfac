package apag_test

import (
	"fmt"
	"testing"

	"github.com/stv0g/cfac/pkg/mobility/apag"
)

func TestFetchAllHouses(t *testing.T) {
	houses, err := apag.FetchAllHouses()
	if err != nil {
		t.Errorf("Failed to fetch house list: %s", err)
	}

	if houses == nil || len(houses) < 10 {
		t.Fail()
	}

	fmt.Printf("Houses: %#v", houses)
	t.Fail()
}
