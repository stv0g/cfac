package sensorthings

type GeoJSONLocation struct {
	Type     string          `json:"type"`
	Geometry GeoJSONGeometry `json:"geometry"`
}

type GeoJSONPolygon struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
