package main

type GamePlayerAI struct {
	CommonGamePlayer
}

func (receiver GamePlayerAI) Peng(tile Tile) Tile {
	// 寻找相同的牌, 然后移出手牌, 组成list放入固定牌部分
	var list TileList
	for i, handTile := range receiver.HandTiles {
		if len(list) == 2 {
			break
		}
		if tile.tileType == handTile.tileType && tile.tileNum == handTile.tileNum {
			temp, l := receiver.HandTiles.Remove(i)
			list = append(list, temp)
			receiver.HandTiles = l
		}
	}

	// 如果找到了碰的牌
	if len(list) == 2 {
		list = append(list, tile)
		receiver.FinishedTiles = append(receiver.FinishedTiles, TileList(list))
		return receiver.SendTile()
	}

	// 没找到, 返回去
	receiver.HandTiles = append(receiver.HandTiles, list...)
	return tile
}

func (receiver GamePlayerAI) Gang(tile Tile, newTile Tile) Tile {
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

func (receiver GamePlayerAI) SendTile() Tile {
	// 判断最大可能性胡牌的牌型, 现在是先判断哪些胡的张数多, 不考虑已经出现的牌的情况
	return Tile{}
}

func (receiver GamePlayerAI) IsAIPlayer() bool {
	return true
}
