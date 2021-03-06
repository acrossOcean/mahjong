package main

//TileType 牌类型, 包括(饼/条/万/风/龙/花/四季)
type TileType int

const (
	_ TileType = iota
	//TileTypeDot 饼 1-9
	TileTypeDot

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

func TileTypeString(n TileType) string {
	return TileTypeMap()[n]
}

func SpecialTileTypeString(m int, n TileType) string {
	specialMap := map[TileType]map[int]string{
		TileTypeWind: map[int]string{
			1: "东",
			2: "南",
			3: "西",
			4: "北",
		},
		TileTypeDragon: map[int]string{
			1: "红中",
			2: "发财",
			3: "白板",
		},
		TileTypeFlower: map[int]string{
			1: "梅",
			2: "兰",
			3: "菊",
			4: "竹",
		},
		TileTypeSeason: map[int]string{
			1: "春",
			2: "夏",
			3: "秋",
			4: "冬",
		},
	}

	return specialMap[n][m]
}

func TileTypeMap() map[TileType]string {
	var result = map[TileType]string{
		TileTypeDot:       "饼",
		TileTypeBamboo:    "条",
		TileTypeCharacter: "万",
		TileTypeWind:      "风",
		TileTypeDragon:    "",
		TileTypeFlower:    "花",
		TileTypeSeason:    "季",
	}

	return result
}
