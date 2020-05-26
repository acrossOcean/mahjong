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
