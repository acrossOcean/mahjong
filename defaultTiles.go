package main

func DefaultAllDot() TileList {
	var result = []Tile{
		Tile{
			tileType: TileTypeDot,
			tileNum:  1,
		},
		Tile{
			tileType: TileTypeDot,
			tileNum:  2,
		},
		Tile{
			tileType: TileTypeDot,
			tileNum:  3,
		},
		Tile{
			tileType: TileTypeDot,
			tileNum:  4,
		},
		Tile{
			tileType: TileTypeDot,
			tileNum:  5,
		},
		Tile{
			tileType: TileTypeDot,
			tileNum:  6,
		},
		Tile{
			tileType: TileTypeDot,
			tileNum:  7,
		},
		Tile{
			tileType: TileTypeDot,
			tileNum:  8,
		},
		Tile{
			tileType: TileTypeDot,
			tileNum:  9,
		},
	}

	return result
}

func DefaultAllBamboo() TileList {
	var result = []Tile{
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  1,
		},
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  2,
		},
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  3,
		},
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  4,
		},
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  5,
		},
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  6,
		},
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  7,
		},
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  8,
		},
		Tile{
			tileType: TileTypeBamboo,
			tileNum:  9,
		},
	}

	return result
}

func DefaultAllCharacter() TileList {
	var result = []Tile{
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  1,
		},
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  2,
		},
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  3,
		},
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  4,
		},
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  5,
		},
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  6,
		},
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  7,
		},
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  8,
		},
		Tile{
			tileType: TileTypeCharacter,
			tileNum:  9,
		},
	}

	return result
}

func DefaultAllDragon() TileList {
	var result = []Tile{
		Tile{
			tileType: TileTypeDragon,
			tileNum:  1,
		},
		Tile{
			tileType: TileTypeDragon,
			tileNum:  2,
		},
		Tile{
			tileType: TileTypeDragon,
			tileNum:  3,
		},
	}

	return result
}

func DefaultAllWinds() TileList {
	var result = []Tile{
		Tile{
			tileType: TileTypeWind,
			tileNum:  1,
		},
		Tile{
			tileType: TileTypeWind,
			tileNum:  2,
		},
		Tile{
			tileType: TileTypeWind,
			tileNum:  3,
		},
		Tile{
			tileType: TileTypeWind,
			tileNum:  4,
		},
	}

	return result
}

func DefaultAllFlower() TileList {
	var result = []Tile{
		Tile{
			tileType: TileTypeFlower,
			tileNum:  1,
		},
		Tile{
			tileType: TileTypeFlower,
			tileNum:  2,
		},
		Tile{
			tileType: TileTypeFlower,
			tileNum:  3,
		},
		Tile{
			tileType: TileTypeFlower,
			tileNum:  4,
		},
	}

	return result
}

func DefaultAllSeason() TileList {
	var result = []Tile{
		Tile{
			tileType: TileTypeSeason,
			tileNum:  1,
		},
		Tile{
			tileType: TileTypeSeason,
			tileNum:  2,
		},
		Tile{
			tileType: TileTypeSeason,
			tileNum:  3,
		},
		Tile{
			tileType: TileTypeSeason,
			tileNum:  4,
		},
	}

	return result
}

// DefaultAllTiles return a TileList contain all tiles
func DefaultAllTiles() TileList {
	result := append(DefaultAllDot(), DefaultAllBamboo()...)
	result = append(result, DefaultAllCharacter()...)
	result = append(result, DefaultAllDragon()...)
	result = append(result, DefaultAllWinds()...)

	//TODO: 之后再加上花牌和补花功能
	//result = append(result, DefaultAllFlower()...)
	//result = append(result, DefaultAllSeason()...)

	return result
}
