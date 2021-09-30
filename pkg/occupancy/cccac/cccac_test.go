package cccac_test

import (
	"log"
	"testing"

	"github.com/stv0g/cfac/pkg/cccac"
)

func TestFetchStatus(t *testing.T) {
	sts, err := cccac.FetchStatus()
	if err != nil {
		t.Fail()
	}

	log.Printf("Status: %v", sts)
}
