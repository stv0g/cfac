package city

import (
	cfac "github.com/stv0g/cfac/pkg"
)

// TODO: Provide facts about Aachen from Wikidata
// TODO: use ontologies

var Aachen = City{
	Center: cfac.Coordinate{
		Latitude:  50.776351,
		Longitude: 6.083862,
	},
	BoundingBox: cfac.BoundingBox{
		NorthWest: cfac.Coordinate{
			Latitude:  50.642058,
			Longitude: 5.798035,
		},
		SouthEast: cfac.Coordinate{
			Latitude:  50.912991,
			Longitude: 6.443481,
		},
	},
	AGS: "053340002002",
}
