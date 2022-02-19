package geojson

// A list of the geojson types that are currently supported.
const (
	TypePoint           = "Point"
	TypeMultiPoint      = "MultiPoint"
	TypeLineString      = "LineString"
	TypeMultiLineString = "MultiLineString"
	TypePolygon         = "Polygon"
	TypeMultiPolygon    = "MultiPolygon"
)

type BBox []float32

type Coordinate [2]float32

type Object struct {
	Type string `json:"type"`
}

// A Geometry matches the structure of a GeoJSON Geometry.
type Geometry struct {
	Object

	Coordinates []Coordinate `json:"coordinates,omitempty"`
	Geometries  []*Geometry  `json:"geometries,omitempty"`
}

// A Feature corresponds to GeoJSON feature object
type Feature struct {
	ID         interface{}            `json:"id,omitempty"`
	Type       string                 `json:"type"`
	BBox       BBox                   `json:"bbox,omitempty"`
	Geometry   Geometry               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type FeatureCollection struct {
	Object

	Features []Feature
}

type CRS struct {
	Type       string `json:"type"`
	Properties struct {
		Name string `json:"name"`
	} `json:"properties"`
}
