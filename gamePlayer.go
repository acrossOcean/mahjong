package main

type PlayerOperation int

const (
	OperationNull PlayerOperation = 0
	OperationWin  PlayerOperation = 1
	OperationPeng PlayerOperation = 2
	OperationGang PlayerOperation = 3
)

// 通用游戏玩家接口(包含了AI/人类) 玩家均可用的动作定义
type IGamePlayer interface {
	// 安排位置, 给出东南西北四个牌的一个, 坐在这个位置
	SetLocation(tile Tile)
	// 接收多张牌
	AcceptTiles(list TileList)
	// 查询是否需要给定的牌
	IsNeed(tile Tile) PlayerOperation
	// 判断是否胡牌
	IsWin() bool

	// 返回玩家当前所有牌 (包括已经杠了和碰了的牌)
	GetAllTiles() TileList

	// 返回玩家手牌
	GetHandTiles() TileList

	// 判断是否是AI玩家
	IsAIPlayer() bool

	// 打一张牌
	SendTile() Tile

	// 操作
	Peng(Tile) Tile
	Gang(Tile, Tile) Tile
	Win(Tile)
}

type CommonGamePlayer struct {
	// 当前位置
	Location Location
	// 手牌
	HandTiles TileList
	// 已完成牌
	FinishedTiles []TileList
	// 需要的排
	NeededTiles map[Tile]struct{}
}

func (receiver CommonGamePlayer) SetLocation(tile Tile) {
	// 按风字牌定位
	if tile.tileType == TileTypeWind {
		switch tile.tileNum {
		case 1:
			receiver.Location = LocationEast
		case 2:
			receiver.Location = LocationSouth
		case 3:
			receiver.Location = LocationWest
		case 4:
			receiver.Location = LocationNorth
		}
	}
}

func (receiver CommonGamePlayer) AcceptTiles(list TileList) {
	receiver.HandTiles = append(receiver.HandTiles, list...)
}

func (receiver CommonGamePlayer) IsNeed(tile Tile) PlayerOperation {
	return OperationNull
}

func (receiver CommonGamePlayer) IsWin() bool {
	return false
}

func (receiver CommonGamePlayer) GetAllTiles() TileList {
	list := receiver.HandTiles

	for _, tiles := range receiver.FinishedTiles {
		list = append(list, tiles...)
	}

	return NewEmptyGameTileStack().GetTileList()
}

func (receiver CommonGamePlayer) GetHandTiles() TileList {
	return receiver.HandTiles
}
