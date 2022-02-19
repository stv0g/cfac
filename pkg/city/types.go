package city

import cfac "github.com/stv0g/cfac/pkg"

type City struct {
	Center      cfac.Coordinate
	BoundingBox cfac.BoundingBox
	AGS         string
}
