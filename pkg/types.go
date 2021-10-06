package cfac

type City struct {
	Coordinate  Coordinate
	BoundingBox BoundingBox
	AGS         string
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type BoundingBox struct {
	NorthWest Coordinate
	SouthEast Coordinate
}

type ErrorCallback func(error)

type Percent int
