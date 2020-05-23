package main

import "strings"

func ToUpperNum(n int) string {
	var m = map[int]string{1: "一", 2: "二", 3: "三", 4: "四", 5: "五", 6: "六", 7: "七", 8: "八", 9: "九"}
	return m[n]
}

//FormatTiles return formatted tiles string
func FormatTiles(list TileList) string {
	var result strings.Builder

	result.WriteString("{\n")
	if len(list) > 0 {
		lastTileType := list[0].tileType

		for i := 0; i < len(list); i++ {
			if lastTileType != list[i].tileType {
				lastTileType = list[i].tileType
				result.WriteString("\n")
			}
			result.WriteString(list[i].String())
		}
	}

	result.WriteString("\n}")
	return result.String()
}
