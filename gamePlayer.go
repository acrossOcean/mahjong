package main

import (
	"sort"
)

type PlayerOperation int

const (
	OperationNull PlayerOperation = 0
	OperationWin  PlayerOperation = 1
	OperationPeng PlayerOperation = 2
	OperationGang PlayerOperation = 3
)

func (receiver PlayerOperation) String() string {
	switch receiver {
	case OperationNull:
		return "无"
	case OperationWin:
		return "胡"
	case OperationPeng:
		return "碰"
	case OperationGang:
		return "杠"
	default:
		return "未知"
	}
}

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
	// 需要的牌
	NeededTiles map[Tile]PlayerOperation
	// 助手
	Helper GamePlayerHelper
}

func (receiver *CommonGamePlayer) SetLocation(tile Tile) {
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

func (receiver *CommonGamePlayer) AcceptTiles(list TileList) {
	receiver.HandTiles = append(receiver.HandTiles, list...)
	sort.Sort(receiver.HandTiles)
}

func (receiver CommonGamePlayer) IsNeed(tile Tile) PlayerOperation {
	if operation, ok := receiver.NeededTiles[tile]; ok {
		return operation
	}

	return OperationNull
}

func (receiver CommonGamePlayer) getWinNeededTiles(list TileList) TileList {
	sort.Sort(list)
	var neededTiles TileList

	// 判断需要的牌, 即 有这张牌之后可以赢的牌
	for _, tile := range list {
		if tile.tileType == TileTypeDot || tile.tileType == TileTypeBamboo || tile.tileType == TileTypeCharacter {
			// 饼条万 还可以 前补一张,后补一张
			if tile.tileNum > 1 {
				beforeTile := Tile{
					tileType: tile.tileType,
					tileNum:  tile.tileNum - 1,
				}
				if receiver.Helper.getRestNumWithoutList(beforeTile, receiver.GetAllTiles()) >= 1 {
					testList := append(list, beforeTile)
					if receiver.checkWin(testList) {
						neededTiles = append(neededTiles, beforeTile)
					}
				}
			}

			if tile.tileNum < 9 {
				afterTile := Tile{
					tileType: tile.tileType,
					tileNum:  tile.tileNum,
				}

				if receiver.Helper.getRestNumWithoutList(afterTile, receiver.GetAllTiles()) >= 1 {
					testList := append(list, afterTile)
					if receiver.checkWin(testList) {
						neededTiles = append(neededTiles, afterTile)
					}
				}
			}

			sameTile := tile
			if receiver.Helper.getRestNumWithoutList(sameTile, receiver.GetAllTiles()) >= 1 {
				testList := append(list, sameTile)
				if receiver.checkWin(testList) {
					neededTiles = append(neededTiles, sameTile)
				}
			}
		}
	}

	return neededTiles
}

func (receiver CommonGamePlayer) getPengNeededTiles(list TileList) TileList {
	sort.Sort(list)

	var result TileList
	m := make(map[Tile]struct{})

	if len(list) < 1 {
		return result
	}

	for i := 0; i < len(list)-1; i++ {
		tile := list[i]
		if _, ok := m[tile]; !ok {
			if list[i] == list[i+1] {
				m[tile] = struct{}{}
				result = append(result, tile)
			}
		}
	}

	return result
}

func (receiver CommonGamePlayer) getGangNeededTiles(list TileList) TileList {
	sort.Sort(list)

	var result TileList
	m := make(map[Tile]struct{})

	if len(list) < 2 {
		return result
	}

	for i := 0; i < len(list)-2; i++ {
		tile := list[i]
		if _, ok := m[tile]; !ok {
			if list[i] == list[i+1] && list[i] == list[i+2] {
				m[tile] = struct{}{}
				result = append(result, tile)
			}
		}
	}

	return result
}

func (receiver CommonGamePlayer) checkWin(list TileList) bool {
	sort.Sort(list)
	old := list

	// 要有一个对
	// 剩下的组成 三个三个的组合即可
	for i := 0; i < list.Len()-1; i++ {
		if list[i] == list[i+1] {
			// 发现一个对, 除去这个对, 并检查剩下的是否满足 3个一组
			_, newList := list.Remove(i)
			_, newList = list.Remove(i)

			if receiver.checkIsAllThree(newList) {
				return true
			}

			list = old
		}
	}

	return false
}

// 检查是否全是可以组成三个一组的列表
func (receiver CommonGamePlayer) checkIsAllThree(list TileList) bool {
	sort.Sort(list)

	typeSplitList := receiver.splitTileList(list)
	for _, splitList := range typeSplitList {
		if len(splitList)%3 != 0 {
			return false
		}

		for {
			newList, ok := receiver.removeSameOnce(splitList, 3)
			if !ok {
				splitList = newList
				break
			}
		}

		// 剩下的牌 应该都是可以 直接+1 +1 连上的
		for {
			newList, ok := receiver.removeLineOnce(splitList)
			if !ok {
				splitList = newList
			}
		}

		// 没有移除完, 说明有不符的, 所以失败
		if len(splitList) != 0 {
			return false
		}
	}

	return true
}

func (receiver CommonGamePlayer) removeSameOnce(list TileList, count int) (TileList, bool) {
	if count < 1 {
		return list, false
	}

	sort.Sort(list)
	isRemove := false
	for i := 0; i < len(list)-count; i++ {
		tempList := TileList{list[i]}
		isSame := true
		for j := 1; j < count; j++ {
			if list[i+j] != tempList[0] {
				isSame = false
				i += j - 1
				break
			}
		}

		if isSame {
			// 可以移除这些,然后返回了
			for i := 0; i < count; i++ {
				_, list = list.Remove(i)
			}

			isRemove = true
			break
		}
	}

	return list, isRemove
}

func (receiver CommonGamePlayer) removeSpecSameOnce(list TileList, tile Tile, count int) (TileList, bool) {
	if count < 1 {
		return list, false
	}

	sort.Sort(list)
	isRemove := false
	for i := 0; i < len(list)-count; i++ {
		if list[i] != tile {
			continue
		}

		tempList := TileList{list[i]}
		isSame := true
		for j := 1; j < count; j++ {
			if list[i+j] != tempList[0] {
				isSame = false
				i += j - 1
				break
			}
		}

		if isSame {
			// 可以移除这些,然后返回了
			for i := 0; i < count; i++ {
				_, list = list.Remove(i)
			}

			isRemove = true
			break
		}
	}

	return list, isRemove
}

// 移除 一连, 即 三个连在一起的牌
func (receiver CommonGamePlayer) removeLineOnce(list TileList) (TileList, bool) {
	sort.Sort(list)
	if len(list) < 3 {
		return list, false
	}

	var first, second, third Tile
	var i, j, k int
	isRemove := false
	for i = 0; i < len(list)-2; i++ {
		findAll := false
		first = list[i]
		for j = i + 1; j < len(list)-1; j++ {
			if list[j] != first {
				second = list[j]
			}
			for k = j + 1; k < len(list); k++ {
				if list[k] != first && list[k] != second {
					third = list[k]
					findAll = true
				}
			}
		}

		if findAll &&
			(first.tileType == TileTypeDot || first.tileType == TileTypeBamboo || first.tileType == TileTypeCharacter) &&
			first.tileType == second.tileType && first.tileType == third.tileType &&
			first.tileNum+1 == second.tileNum && second.tileNum+1 == third.tileNum {
			// 找到了
			isRemove = true
			_, list = list.Remove(k)
			_, list = list.Remove(j)
			_, list = list.Remove(i)
		}
	}

	return list, isRemove
}

func (receiver CommonGamePlayer) IsWin() bool {
	return false
}

func (receiver CommonGamePlayer) GetAllTiles() TileList {
	list := receiver.HandTiles

	for _, tiles := range receiver.FinishedTiles {
		list = append(list, tiles...)
	}

	return list
}

func (receiver CommonGamePlayer) GetHandTiles() TileList {
	return receiver.HandTiles
}

// 按照所给的list,把不同 tileType的 tile 放到不同的list中
func (receiver CommonGamePlayer) splitTileList(list TileList) []TileList {
	result := make([]TileList, 0)

	sort.Sort(list)

	var currentType TileType
	var currentTileList TileList

	for i, tile := range list {
		if i == 0 {
			currentType = tile.tileType
		}

		if currentType == tile.tileType {
			currentTileList = append(currentTileList, tile)
		} else {
			currentType = tile.tileType
			result = append(result, currentTileList)
			currentTileList = TileList{tile}
		}
	}

	if len(currentTileList) > 0 {
		result = append(result, currentTileList)
	}

	return result
}

// 获取最无用的牌
func (receiver CommonGamePlayer) getMostUselessTile(list TileList) (Tile, TileList) {
	sort.Sort(list)

	for i, tile := range list {
		if tile.tileType != TileTypeDot && tile.tileType != TileTypeBamboo || tile.tileType != TileTypeCharacter {
			// 不是饼条万, 如果没有对, 就是这个最垃圾, 返回就行
			if _, ok := receiver.removeSpecSameOnce(list, tile, 2); !ok {
				return list.Remove(i)
			}
		} else {
			// 有对, 跳过
			if _, ok := receiver.removeSpecSameOnce(list, tile, 2); !ok {
				continue
			}

			// 有没有三个连起来的, 有就跳过
			if receiver.checkHasLineInList(tile, list) {
				continue
			} else {
				return list.Remove(i)
			}
		}
	}

	return list.Remove(0)
}

func (receiver CommonGamePlayer) checkHasLineInList(tile Tile, list TileList) bool {
	// 有没有三个连起来的, 有就跳过
	var b2, b1, a1, a2 bool
	if tile.tileNum > 2 {
		b2, _ = receiver.findTileInList(Tile{tileType: tile.tileType, tileNum: tile.tileNum - 2}, list, 0)
	}
	if tile.tileNum > 1 {
		b1, _ = receiver.findTileInList(Tile{tileType: tile.tileType, tileNum: tile.tileNum - 1}, list, 0)
	}
	if tile.tileNum < 9 {
		a1, _ = receiver.findTileInList(Tile{tileType: tile.tileType, tileNum: tile.tileNum + 1}, list, 0)
	}
	if tile.tileNum < 8 {
		a2, _ = receiver.findTileInList(Tile{tileType: tile.tileType, tileNum: tile.tileNum + 2}, list, 0)
	}

	if (b1 && b2) || (a1 && a2) || (b1 && a1) {
		//有三个连一起, 可以跳过
		return true
	}
	return false
}

func (receiver CommonGamePlayer) findTileInList(tile Tile, list TileList, startAt int) (bool, int) {
	if startAt < 0 {
		startAt = 0
	}
	for i := startAt; i < len(list); i++ {
		if tile == list[i] {
			return true, i
		}
	}

	return false, -1
}
