package tdrm

import cfac "github.com/stv0g/cfac/pkg"

type Station struct {
	ID         int
	Name       string
	Region     string
	Coordinate cfac.Coordinate
	AvgValue   float64
}
