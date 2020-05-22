package main

//TileType 牌类型, 包括(饼/条/万/风/龙/花/四季)
type TileType int

//Tile 单个牌
type Tile struct {
	//tileType 牌类型
	tileType TileType
	//tileNum 数值, 代表含义根据类型不同决定
	tileNum int
}

const (
	//TileTypeDot 饼 1-9
	TileTypeDot = iota

	//TileTypeBamboo 条 1-9
	TileTypeBamboo

	//TileTypeCharacter 万 1-9
	TileTypeCharacter

	//TileTypeWind 风 (1:东  2:南  3:西  4:北)
	TileTypeWind

	//TileTypeDragon 龙 (1:中  2:发   3:白)
	TileTypeDragon

	//TileTypeFlower 花 (1:梅  2:兰   3:菊   4:竹)
	TileTypeFlower

	//TileTypeSeason 季节 (1:春  2:夏   3:秋   4:冬)
	TileTypeSeason
)
