package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newAIPlayer() *GamePlayerAI {
	gm := NewGameManager()
	ai := NewAI(gm)

	return ai
}

// ======================公共方法======================

func TestCommonGamePlayer_SetLocation(t *testing.T) {
	ai := newAIPlayer()

	tile := NewTile(TileTypeWind, 1)
	ai.SetLocation(tile)
	assert.Equal(t, ai.GetLocation(), LocationEast)

	tile = NewTile(TileTypeWind, 2)
	ai.SetLocation(tile)
	assert.Equal(t, ai.GetLocation(), LocationSouth)

	tile = NewTile(TileTypeWind, 3)
	ai.SetLocation(tile)
	assert.Equal(t, ai.GetLocation(), LocationWest)

	tile = NewTile(TileTypeWind, 4)
	ai.SetLocation(tile)
	assert.Equal(t, ai.GetLocation(), LocationNorth)
}

func TestCommonGamePlayer_AcceptTiles(t *testing.T) {
	ai := newAIPlayer()
	tile := NewTile(TileTypeDot, 3)
	list := TileList{tile}

	ai.AcceptTiles(list)

	assert.Equal(t, ai.HandTiles, list)
}

func TestCommonGamePlayer_IsNeed(t *testing.T) {
	ai := newAIPlayer()

	tile := NewTile(TileTypeWind, 1)

	assert.Equal(t, ai.IsNeed(tile), OperationNull)

	ai.NeededTiles = map[Tile]PlayerOperation{
		tile: OperationPeng,
	}

	assert.Equal(t, ai.IsNeed(tile), OperationPeng)
}

func TestCommonGamePlayer_GetHandTiles(t *testing.T) {
	ai := newAIPlayer()
	tile := NewTile(TileTypeDot, 3)
	ai.HandTiles = TileList{tile}
	list := TileList{tile}

	assert.Equal(t, ai.GetHandTiles(), list)
}

func TestCommonGamePlayer_GetAllTiles(t *testing.T) {
	ai := newAIPlayer()
	tile := NewTile(TileTypeDot, 3)
	list := TileList{tile}
	finishedList := TileList{tile, tile, tile}

	ai.HandTiles = TileList{tile}
	ai.FinishedTiles = []TileList{finishedList}

	all := ai.GetAllTiles()
	actual := append(finishedList, list...)

	sort.Sort(all)
	sort.Sort(actual)

	assert.Equal(t, all, actual)
}

// ======================私有方法======================
func TestCommonGamePlayer_getWinNeededTiles(t *testing.T) {
	var gm *GameManager
	var ai *GamePlayerAI
	// 需要对
	list1 := TileList{
		NewTile(TileTypeWind, 1),
	}

	list1Need := TileList{
		NewTile(TileTypeWind, 1),
	}

	gm = NewGameManagerWithStack(*NewGameTileStack(list1Need))
	ai = NewAI(gm)

	afterMap1 := ai.getWinNeededTiles(list1)
	afterList1 := ai.tileMapToList(afterMap1)
	sort.Sort(afterList1)
	sort.Sort(list1Need)
	assert.Equal(t, afterList1, list1Need)

	// 需要一连
	list2 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 3),
	}

	list2Need := TileList{
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeCharacter, 4),
	}

	gm = NewGameManagerWithStack(*NewGameTileStack(list2Need))
	ai = NewAI(gm)

	afterMap2 := ai.getWinNeededTiles(list2)
	afterList2 := ai.tileMapToList(afterMap2)
	sort.Sort(afterList2)
	sort.Sort(list2Need)
	assert.Equal(t, afterList2, list2Need)

	// 需要三个的
	list3 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 1),
		NewTile(TileTypeBamboo, 1),
	}

	list3Need := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 1),
	}

	gm = NewGameManagerWithStack(*NewGameTileStack(list3Need))
	ai = NewAI(gm)

	afterMap3 := ai.getWinNeededTiles(list3)
	afterList3 := ai.tileMapToList(afterMap3)
	sort.Sort(afterList3)
	sort.Sort(list3Need)
	assert.Equal(t, afterList3, list3Need)

	// 复杂型胡牌
	list4 := TileList{
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 4),
		NewTile(TileTypeBamboo, 5),
	}

	list4Need := TileList{
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 5),
	}

	gm = NewGameManagerWithStack(*NewGameTileStack(list4Need))
	ai = NewAI(gm)

	afterMap4 := ai.getWinNeededTiles(list4)
	afterList4 := ai.tileMapToList(afterMap4)
	sort.Sort(afterList4)
	sort.Sort(list4Need)
	assert.Equal(t, afterList4, list4Need)
}

func TestCommonGamePlayer_getPengNeededTiles(t *testing.T) {
	ai := newAIPlayer()

	ai.HandTiles = TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 4),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 4),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 2),
	}

	ai.HandTiles = ai.HandTiles.Shuffle()

	pengList := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 4),
		NewTile(TileTypeCharacter, 2),
	}

	list := ai.getPengNeededTiles(ai.GetHandTiles())

	sort.Sort(pengList)
	sort.Sort(list)

	assert.Equal(t, pengList, list)
}

func TestCommonGamePlayer_getGangNeededTiles(t *testing.T) {
	ai := newAIPlayer()

	ai.HandTiles = TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 4),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 4),
		NewTile(TileTypeBamboo, 4),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 2),
	}

	ai.HandTiles = ai.HandTiles.Shuffle()

	gangList := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 4),
	}

	list := ai.getGangNeededTiles(ai.GetHandTiles())

	sort.Sort(gangList)
	sort.Sort(list)

	assert.Equal(t, gangList, list)
}

func TestCommonGamePlayer_findTileInList(t *testing.T) {
	ai := newAIPlayer()
	tile := NewTile(TileTypeWind, 1)
	list := TileList{
		tile,
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeBamboo, 1),
		tile,
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 9),
	}

	ok, index := ai.findTileInList(tile, list, -1)
	assert.Equal(t, ok, true)
	assert.Equal(t, index, 0)

	ok, index = ai.findTileInList(tile, list, 1)
	assert.Equal(t, ok, true)
	assert.Equal(t, index, 3)

	ok, index = ai.findTileInList(tile, list, 4)
	assert.Equal(t, ok, false)
}

func TestCommonGamePlayer_checkHasLineInList(t *testing.T) {
	ai := newAIPlayer()

	list := TileList{
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 5),
		NewTile(TileTypeBamboo, 4),
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 2),
		NewTile(TileTypeDot, 4),
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeCharacter, 9),
		NewTile(TileTypeCharacter, 1),
	}

	list = list.Shuffle()

	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeBamboo, 1), list), true)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeBamboo, 2), list), true)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeBamboo, 3), list), true)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeBamboo, 4), list), true)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeBamboo, 5), list), true)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeBamboo, 6), list), true)

	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeDot, 3), list), true)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeDot, 1), list), false)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeDot, 5), list), false)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeDot, 9), list), false)

	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeWind, 1), list), false)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeWind, 2), list), false)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeWind, 4), list), false)

	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeCharacter, 2), list), false)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeCharacter, 8), list), false)
	assert.Equal(t, ai.checkHasLineInList(NewTile(TileTypeCharacter, 1), list), false)
}

func TestCommonGamePlayer_splitTileList(t *testing.T) {
	ai := newAIPlayer()

	list1 := TileList{
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 4),
	}

	list2 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 4),
	}

	list3 := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 2),
		NewTile(TileTypeDot, 4),
	}

	list := append(list1, list2...)
	list = append(list, list3...)
	list = list.Shuffle()

	after := ai.splitTileList(list)
	for _, l := range after {
		if assert.NotZero(t, l.Len()) {
			sort.Sort(l)
			switch l[0].tileType {
			case TileTypeCharacter:
				sort.Sort(list1)
				assert.Equal(t, l, list1)
			case TileTypeWind:
				sort.Sort(list2)
				assert.Equal(t, l, list2)
			case TileTypeDot:
				sort.Sort(list3)
				assert.Equal(t, l, list3)
			}
		}
	}
}

func TestCommonGamePlayer_removeLineOnce(t *testing.T) {
	ai := newAIPlayer()

	lineList := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 2),
		NewTile(TileTypeDot, 3),
	}

	notLine1 := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 8),
		NewTile(TileTypeDot, 9),
	}

	notLine2 := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 2),
		NewTile(TileTypeDot, 4),
	}

	notLine3 := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 2),
	}

	list := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeWind, 4),
	}

	all := append(list, lineList...)
	all = all.Shuffle()

	after, ok := ai.removeLineOnce(all)
	assert.Equal(t, ok, true)
	assert.Equal(t, after, list)

	all = append(list, lineList...)
	all = append(all, lineList...)
	all = all.Shuffle()
	after, ok = ai.removeLineOnce(all)
	after, ok = ai.removeLineOnce(after)
	assert.Equal(t, ok, true)
	assert.Equal(t, after, list)

	after, ok = ai.removeLineOnce(list)
	assert.Equal(t, ok, false)

	all2 := append(list, notLine1...)
	all2 = all2.Shuffle()
	after, ok = ai.removeLineOnce(all2)
	assert.Equal(t, ok, false)
	assert.Equal(t, after, all2)

	all3 := append(list, notLine2...)
	all3 = all3.Shuffle()
	after, ok = ai.removeLineOnce(all3)
	assert.Equal(t, ok, false)
	assert.Equal(t, after, all3)

	after, ok = ai.removeLineOnce(notLine3)
	assert.Equal(t, ok, false)
	assert.Equal(t, after, notLine3)
}

func TestCommonGamePlayer_removeSameOnce(t *testing.T) {
	ai := newAIPlayer()

	list := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeWind, 1),
	}

	same1 := TileList{
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 2),
	}

	same2 := TileList{
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 3),
	}

	all := append(list, same1...)
	all = all.Shuffle()
	after, ok := ai.removeSameOnce(all, 2)
	sort.Sort(after)
	sort.Sort(list)

	assert.Equal(t, ok, true)
	assert.Equal(t, after, list)

	all = append(list, same2...)
	all = all.Shuffle()
	after, ok = ai.removeSameOnce(all, 3)
	sort.Sort(after)
	sort.Sort(list)

	assert.Equal(t, ok, true)
	assert.Equal(t, after, list)

	all = append(list, same2...)
	all = all.Shuffle()
	after, ok = ai.removeSameOnce(all, 99)
	sort.Sort(after)
	sort.Sort(list)

	assert.Equal(t, ok, false)
	assert.Equal(t, after, all)

	after, ok = ai.removeSameOnce(list, 2)
	sort.Sort(after)
	sort.Sort(list)

	assert.Equal(t, ok, false)
	assert.Equal(t, after, list)
}

func TestCommonGamePlayer_removeSpecSameOnce(t *testing.T) {
	ai := newAIPlayer()

	list := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 1),
		NewTile(TileTypeBamboo, 1),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 4),
	}

	same1 := TileList{
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 2),
	}

	same2 := TileList{
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 3),
	}

	all := append(list, same1...)
	all = all.Shuffle()
	after, ok := ai.removeSpecSameOnce(all, NewTile(TileTypeBamboo, 2), 2)
	sort.Sort(after)
	sort.Sort(list)

	assert.Equal(t, ok, true)
	assert.Equal(t, after, list)

	all = append(list, same2...)
	all = all.Shuffle()
	after, ok = ai.removeSpecSameOnce(all, NewTile(TileTypeCharacter, 3), 3)
	sort.Sort(after)
	sort.Sort(list)

	assert.Equal(t, ok, true)
	assert.Equal(t, after, list)

	all = append(list, same2...)
	all = all.Shuffle()
	after, ok = ai.removeSpecSameOnce(all, NewTile(TileTypeCharacter, 3), 99)
	sort.Sort(after)
	sort.Sort(list)

	assert.Equal(t, ok, false)
	assert.Equal(t, after, all)

	after, ok = ai.removeSpecSameOnce(list, NewTile(TileTypeFlower, 2), 2)
	sort.Sort(after)
	sort.Sort(list)

	assert.Equal(t, ok, false)
	assert.Equal(t, after, list)
}

func TestCommonGamePlayer_checkWin(t *testing.T) {
	ai := newAIPlayer()
	notWinList1 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeWind, 4),
	}

	notWinList2 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 4),
	}

	notWinList3 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
	}

	winList1 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
	}

	winList2 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 2),
	}

	winList3 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 2),
	}

	winList4 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 4),
	}

	assert.Equal(t, false, ai.checkWin(notWinList1))
	assert.Equal(t, false, ai.checkWin(notWinList2))
	assert.Equal(t, false, ai.checkWin(notWinList3))
	assert.Equal(t, true, ai.checkWin(winList1))
	assert.Equal(t, true, ai.checkWin(winList2))
	assert.Equal(t, true, ai.checkWin(winList3))
	assert.Equal(t, true, ai.checkWin(winList4))
}

func TestCommonGamePlayer_checkIsAllThree(t *testing.T) {
	ai := newAIPlayer()
	line := TileList{
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 3),
	}

	all := append(line, line...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), true)

	same := TileList{
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 3),
	}

	all = append(same, same...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), true)

	all = append(line, same...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), true)

	same2 := TileList{
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeCharacter, 1),
	}

	all = append(line, same...)
	all = append(all, same2...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), true)

	same3 := TileList{
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 2),
	}

	all = append(line, same...)
	all = append(all, same3...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), true)

	all = append(line, same...)
	all = append(all, same2...)
	all = append(all, same3...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), true)

	same4 := TileList{
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 4),
	}

	all = append(line, same...)
	all = append(all, same4...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), true)

	all = append(line, same...)
	all = append(all, same2...)
	all = append(all, same3...)
	all = append(all, same4...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), true)

	notLine := TileList{
		NewTile(TileTypeBamboo, 1),
		NewTile(TileTypeBamboo, 8),
		NewTile(TileTypeBamboo, 9),
	}

	all = notLine
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), false)

	notSame := TileList{
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeWind, 4),
	}

	all = notSame
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), false)

	all = notSame
	all = append(all, notLine...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), false)

	all = notSame
	all = append(all, notLine...)
	all = append(all, notSame...)
	all = append(all, notLine...)
	all = all.Shuffle()
	assert.Equal(t, ai.checkIsAllThree(all), false)
}

func TestCommonGamePlayer_getMostUselessTile(t *testing.T) {
	ai := newAIPlayer()

	list1 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 1),
	}

	list2 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 1),
	}

	list3 := TileList{
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 3),
	}

	list4 := TileList{
		NewTile(TileTypeBamboo, 1),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 5),
	}

	list5 := TileList{
		NewTile(TileTypeBamboo, 1),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeCharacter, 1),
	}

	tile1, _ := ai.getMostUselessTile(list1)
	assert.Equal(t, tile1, NewTile(TileTypeWind, 1))

	tile2, _ := ai.getMostUselessTile(list2)
	assert.Equal(t, tile2, NewTile(TileTypeBamboo, 1))

	tile3, _ := ai.getMostUselessTile(list3)
	assert.Equal(t, tile3, NewTile(TileTypeBamboo, 2))

	tile4, _ := ai.getMostUselessTile(list4)
	assert.Equal(t, tile4, NewTile(TileTypeBamboo, 5))

	tile5, _ := ai.getMostUselessTile(list5)
	assert.Equal(t, tile5, NewTile(TileTypeCharacter, 1))
}

func TestCommonGamePlayer_getMostUselessTile1(t *testing.T) {
	ai := newAIPlayer()

	list := TileList{
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 5),
		NewTile(TileTypeCharacter, 8),
		NewTile(TileTypeCharacter, 9),
		NewTile(TileTypeCharacter, 9),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 5),
		NewTile(TileTypeDot, 2),
		NewTile(TileTypeDot, 3),
	}

	should := NewTile(TileTypeBamboo, 5)

	tile, _ := ai.getMostUselessTile(list)
	assert.Equal(t, tile, should)
}
