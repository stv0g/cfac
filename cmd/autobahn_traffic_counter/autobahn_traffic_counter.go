package main

import (
	"os"

	"github.com/stv0g/cfac/pkg/mobility/traffic_counter"
)

func main() {
	var cars chan traffic_counter.Car

	if err := traffic_counter.TrafficCounter(os.Args[1], cars); err != nil {
		panic(err)
	}

}
