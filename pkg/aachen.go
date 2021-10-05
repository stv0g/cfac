package cfac

// TODO: Provide facts about Aachen from Wikidata

func GetBoundingBox() BoundingBox {
	return BoundingBox{
		NorthWest: Coordinate{50.642058, 5.798035},
		SouthEast: Coordinate{50.912991, 6.443481},
	}
}
