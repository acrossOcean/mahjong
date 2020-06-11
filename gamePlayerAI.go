package main

import (
	"fmt"
	"sort"
)

type GamePlayerAI struct {
	CommonGamePlayer
}

func NewAI(helper GamePlayerHelper) *GamePlayerAI {
	result := new(GamePlayerAI)
	result.Helper = helper
	result.NeededTiles = make(map[Tile]PlayerOperation)
	result.FinishedTiles = make([]TileList, 0)

	return result
}

func (receiver GamePlayerAI) String() string {
	return fmt.Sprint("AI(", receiver.Location, "): ")
}

func (receiver *GamePlayerAI) Peng(tile Tile) Tile {
	// 寻找相同的牌, 然后移出手牌, 组成list放入固定牌部分
	var list TileList
	var after TileList
	for i, handTile := range receiver.HandTiles {
		if len(list) == 2 {
			break
		}
		if tile.tileType == handTile.tileType && tile.tileNum == handTile.tileNum {
			var temp Tile
			temp, after = receiver.HandTiles.Remove(i)
			list = append(list, temp)
		}
	}

	// 如果找到了碰的牌
	if len(list) == 2 {
		list = append(list, tile)
		receiver.FinishedTiles = append(receiver.FinishedTiles, TileList(list))
		receiver.HandTiles = after
		return receiver.SendTile()
	}

	// 没找到, 返回去
	receiver.HandTiles = append(receiver.HandTiles, list...)
	return tile
}

func (receiver *GamePlayerAI) Gang(tile Tile, newTile Tile) Tile {
	// 寻找相同的牌, 然后移出手牌, 组成list放入固定牌部分
	var list TileList
	for i, handTile := range receiver.HandTiles {
		if len(list) == 3 {
			break
		}
		if tile.tileType == handTile.tileType && tile.tileNum == handTile.tileNum {
			temp, l := receiver.HandTiles.Remove(i)
			list = append(list, temp)
			receiver.HandTiles = l
		}
	}

	// 如果找到了杠的牌
	if len(list) == 3 {
		list = append(list, tile)
		receiver.FinishedTiles = append(receiver.FinishedTiles, TileList(list))

		receiver.AcceptTiles(TileList{newTile})
		return receiver.SendTile()
	}

	// 没找到, 返回去
	receiver.HandTiles = append(receiver.HandTiles, list...)
	return tile
}

func (receiver GamePlayerAI) Win(Tile) {

}

func (receiver *GamePlayerAI) SendTile() Tile {
	maxNeedCount := -1
	var uselessTile Tile
	//old := receiver.HandTiles

	// 轮询手中的牌, 每张牌扔一次, 然后进行 加一张牌胡 的判断, 看扔哪张牌胡的几率大, 如果相同, 那么先扔离的远的, 比如 不成对的风/单张的等
	for i := range receiver.HandTiles {
		//fmt.Println("手牌:", receiver.HandTiles)
		tile, list := receiver.HandTiles.Remove(i)
		//fmt.Println("移出:", tile, "剩余:", list)
		//fmt.Println("old:", old)
		//fmt.Println("-----------")

		neededList := receiver.getWinNeededTiles(list)
		if len(neededList) > maxNeedCount {
			maxNeedCount = len(neededList)
			uselessTile = tile
		}

		//receiver.HandTiles = old
		receiver.HandTiles = append(list, tile)
		sort.Sort(receiver.HandTiles)
	}

	if maxNeedCount < 1 {
		// 如果找不到一个合适的可以扔的牌, 那么扔一张影响最小的
		uselessTile, receiver.HandTiles = receiver.getMostUselessTile(receiver.HandTiles)
	} else {
		receiver.HandTiles, _ = receiver.removeSpecSameOnce(receiver.HandTiles, uselessTile, 1)
	}

	// 更新当前需要的牌
	receiver.NeededTiles = map[Tile]PlayerOperation{}
	list := receiver.getWinNeededTiles(receiver.HandTiles)
	for _, tile := range list {
		receiver.NeededTiles[tile] = OperationWin
	}

	list = receiver.getGangNeededTiles(receiver.HandTiles)
	for _, tile := range list {
		if _, ok := receiver.NeededTiles[tile]; !ok {
			receiver.NeededTiles[tile] = OperationGang
		}
	}

	list = receiver.getPengNeededTiles(receiver.HandTiles)
	for _, tile := range list {
		if op, ok := receiver.NeededTiles[tile]; !ok {
			receiver.NeededTiles[tile] = OperationPeng
		} else if op == OperationGang {
			// 判断碰合适还是杠合适
			b2 := Tile{tileType: tile.tileType, tileNum: tile.tileNum - 2}
			b1 := Tile{tileType: tile.tileType, tileNum: tile.tileNum - 1}
			a1 := Tile{tileType: tile.tileType, tileNum: tile.tileNum + 1}
			a2 := Tile{tileType: tile.tileType, tileNum: tile.tileNum + 2}

			newList := list
			newList, _ = receiver.removeSpecSameOnce(newList, b2, 3)
			newList, _ = receiver.removeSpecSameOnce(newList, b1, 3)
			newList, _ = receiver.removeSpecSameOnce(newList, a1, 3)
			newList, _ = receiver.removeSpecSameOnce(newList, a2, 3)

			if receiver.checkHasLineInList(tile, newList) {
				receiver.NeededTiles[tile] = OperationPeng
			}
		}
	}

	return uselessTile
}

func (receiver GamePlayerAI) IsAIPlayer() bool {
	return true
}
