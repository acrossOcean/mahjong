package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTileList_Sort(t *testing.T) {
	tiles := TileList{
		Tile{tileType: TileTypeCharacter, tileNum: 3},
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
	}
	result := TileList{
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
		Tile{tileType: TileTypeCharacter, tileNum: 3},
	}

	fmt.Println("before:", tiles)

	sort.Sort(tiles)

	fmt.Println("after:", tiles)
	fmt.Println("want:", result)

	assert.Equal(t, tiles, result)
}

func TestTileList_Remove(t *testing.T) {
	tiles := TileList{
		Tile{tileType: TileTypeCharacter, tileNum: 3},
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
	}

	a, b := tiles.Remove(1)
	assert.Equal(t, a, Tile{tileType: TileTypeBamboo, tileNum: 1})
	assert.Equal(t, b, TileList{{tileType: TileTypeCharacter, tileNum: 3}, {tileType: TileTypeCharacter, tileNum: 2}})

	tiles = TileList{
		Tile{tileType: TileTypeCharacter, tileNum: 3},
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
	}
	a, b = tiles.Remove(-1)
	assert.Equal(t, a, Tile{tileType: TileTypeCharacter, tileNum: 3})
	assert.Equal(t, b, TileList{{tileType: TileTypeBamboo, tileNum: 1}, {tileType: TileTypeCharacter, tileNum: 2}})

	tiles = TileList{
		Tile{tileType: TileTypeCharacter, tileNum: 3},
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
	}
	a, b = tiles.Remove(99)
	assert.Equal(t, a, Tile{tileType: TileTypeCharacter, tileNum: 2})
	assert.Equal(t, b, TileList{{tileType: TileTypeCharacter, tileNum: 3}, {tileType: TileTypeBamboo, tileNum: 1}})
}

func TestTileList_Shuffle(t *testing.T) {
	tiles := DefaultAllTiles()

	fmt.Println("before:", tiles)
	after := tiles.Shuffle()
	fmt.Println("after:", after)
	assert.Equal(t, tiles.Len(), after.Len())
}

func TestTileList_GetOneFromHead(t *testing.T) {
	first := Tile{tileType: TileTypeCharacter, tileNum: 3}
	list := TileList{
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
	}
	tiles := TileList(append([]Tile{first}, list...))

	tile, result := tiles.GetOneFromHead()

	assert.Equal(t, first, tile)
	assert.Equal(t, result, list)
}

func TestTileList_GetOneFromEnd(t *testing.T) {
	last := Tile{tileType: TileTypeCharacter, tileNum: 3}
	list := TileList{
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
	}
	tiles := append(list, last)

	tile, result := tiles.GetOneFromEnd()

	assert.Equal(t, last, tile)
	assert.Equal(t, result, list)
}

func TestTileList_GetFromHead(t *testing.T) {
	first := TileList{
		{tileType: TileTypeCharacter, tileNum: 3},
		{tileType: TileTypeCharacter, tileNum: 6},
	}
	list := TileList{
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeBamboo, tileNum: 4},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
	}
	tiles := append(first, list...)

	tile, result := tiles.GetFromHead(len(first))

	assert.Equal(t, first, tile)
	assert.Equal(t, result, list)
}

func TestTileList_GetFromEnd(t *testing.T) {
	last := TileList{
		{tileType: TileTypeCharacter, tileNum: 3},
		{tileType: TileTypeCharacter, tileNum: 6},
	}
	list := TileList{
		Tile{tileType: TileTypeBamboo, tileNum: 1},
		Tile{tileType: TileTypeBamboo, tileNum: 4},
		Tile{tileType: TileTypeCharacter, tileNum: 2},
	}
	tiles := append(list, last...)

	tile, result := tiles.GetFromEnd(len(last))

	assert.Equal(t, last, tile)
	assert.Equal(t, result, list)
}
