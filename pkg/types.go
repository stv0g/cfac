package cfac

type Coordinate struct {
	Latitude  float32
	Longitude float32
}

type BoundingBox struct {
	NorthWest Coordinate
	SouthEast Coordinate
}

type ErrorCallback func(error)

type Percent int
