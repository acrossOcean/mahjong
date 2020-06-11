package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameTileStack_Shuffle(t *testing.T) {
	stack := NewEmptyGameTileStack()
	stack.list = append(stack.list, NewTile(TileTypeBamboo, 1))
	stack.list = append(stack.list, NewTile(TileTypeBamboo, 3))
	stack.list = append(stack.list, NewTile(TileTypeBamboo, 2))
	stack.list = append(stack.list, NewTile(TileTypeBamboo, 5))

	stack.Shuffle()
	assert.Equal(t, len(stack.list), 4)
}

func TestGameTileStack_GetTileList(t *testing.T) {
	list := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 3),
		NewTile(TileTypeDot, 4),
		NewTile(TileTypeDot, 8),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 6),
		NewTile(TileTypeBamboo, 9),
	}

	stack := NewGameTileStack(list)

	assert.Equal(t, stack.GetTileList(), list)
}

func TestGameTileStack_GetOneFromHead(t *testing.T) {
	first := NewTile(TileTypeWind, 2)
	list := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 3),
		NewTile(TileTypeDot, 4),
		NewTile(TileTypeDot, 8),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 6),
		NewTile(TileTypeBamboo, 9),
	}

	stack := NewGameTileStack(append(TileList{first}, list...))

	assert.Equal(t, stack.GetOneFromHead(), first)
	assert.Equal(t, stack.GetTileList().Len(), len(list))
	assert.Equal(t, stack.GetTileList(), list)
}

func TestGameTileStack_GetOneFromEnd(t *testing.T) {
	end := NewTile(TileTypeFlower, 1)
	list := TileList{
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 3),
		NewTile(TileTypeDot, 4),
		NewTile(TileTypeDot, 8),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 6),
		NewTile(TileTypeBamboo, 9),
	}

	stack := NewGameTileStack(append(list, end))

	assert.Equal(t, stack.GetOneFromEnd(), end)
	assert.Equal(t, stack.GetTileList().Len(), len(list))
	assert.Equal(t, stack.GetTileList(), list)
}
