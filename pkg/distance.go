// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package cfac

import "math"

const (
	EarthRadius = 6378100 // in meters
)

// Distance function returns the distance (in meters) between two points of
//
//	a given longitude and latitude relatively accurately (using a spherical
//	approximation of the Earth) through the Haversin Distance Formula for
//	great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func (c Coordinate) DistanceTo(to Coordinate) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2 float64
	la1 = c.Latitude * math.Pi / 180
	lo1 = c.Longitude * math.Pi / 180
	la2 = to.Latitude * math.Pi / 180
	lo2 = to.Longitude * math.Pi / 180

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * EarthRadius * math.Asin(math.Sqrt(h))
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
