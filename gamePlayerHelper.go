package main

// 给游戏玩家提供帮助的
type GamePlayerHelper interface {
	// 获取指定牌在可见范围下还有多少张
	getRestNum(Tile) int
	// 获取指定牌在可见范围下 出list外还有多少张
	getRestNumWithoutList(tile Tile, list TileList) int
}
