package main

//Tile 单个牌
type Tile struct {
	//tileType 牌类型
	tileType TileType
	//tileNum 数值, 代表含义根据类型不同决定
	tileNum int
}

func NewTile(tileType TileType, num int) Tile {
	return Tile{tileType: tileType, tileNum: num}
}

func (receiver Tile) String() string {
	result := " "

	if receiver.tileType == TileTypeDot || receiver.tileType == TileTypeBamboo || receiver.tileType == TileTypeCharacter {
		result += ToUpperNum(receiver.tileNum)
	} else {
		result += SpecialTileTypeString(receiver.tileNum, receiver.tileType)
	}

	result += TileTypeString(receiver.tileType) + " "
	return result
}
