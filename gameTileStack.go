package main

type GameTileStack struct {
	list TileList
}

func NewEmptyGameTileStack() *GameTileStack {
	result := new(GameTileStack)

	return result
}

func NewGameTileStack(list TileList) *GameTileStack {
	result := new(GameTileStack)
	result.list = list

	return result
}

func (receiver GameTileStack) GetTileList() TileList {
	return receiver.list
}

func (receiver *GameTileStack) Shuffle() {
	result := receiver.list.Shuffle()
	receiver.list = result
}

func (receiver *GameTileStack) GetFromHead(n int) TileList {
	result, list := receiver.list.GetFromHead(n)
	receiver.list = list

	return result
}

func (receiver *GameTileStack) GetOneFromHead() Tile {
	result := receiver.GetFromHead(1)

	return result[0]
}

func (receiver *GameTileStack) GetFromEnd(n int) TileList {
	result, list := receiver.list.GetFromEnd(n)
	receiver.list = list

	return result
}

func (receiver *GameTileStack) GetOneFromEnd() Tile {
	result := receiver.GetFromEnd(1)

	return result[0]
}
