package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamePlayerAI_SendTile(t *testing.T) {
	ai := newAIPlayer()

	// 初始手牌
	handTileList := TileList{
		NewTile(TileTypeDragon, 1),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeWind, 3),
		NewTile(TileTypeCharacter, 1),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 3),
		NewTile(TileTypeCharacter, 6),
		NewTile(TileTypeCharacter, 6),
		NewTile(TileTypeCharacter, 7),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 6),
		NewTile(TileTypeDot, 2),
		NewTile(TileTypeDot, 6),
	}
	ai.AcceptTiles(handTileList)

	// 抓牌
	ai.AcceptTiles(TileList{
		NewTile(TileTypeBamboo, 2),
	})

	fmt.Println("手牌:", ai.GetHandTiles())

	shouldTile := NewTile(TileTypeDragon, 1)
	tile := ai.SendTile()

	fmt.Println("应该是:", shouldTile, " 实际是:", tile)
	assert.Equal(t, shouldTile, tile)
}

func TestGamePlayerAI_Peng(t *testing.T) {
	ai := newAIPlayer()

	list := TileList{
		NewTile(TileTypeWind, 2),
		NewTile(TileTypeWind, 2),
	}

	hand := TileList{
		NewTile(TileTypeDragon, 2),
		NewTile(TileTypeWind, 1),
		NewTile(TileTypeBamboo, 8),
		NewTile(TileTypeDot, 3),
		NewTile(TileTypeDot, 5),
		NewTile(TileTypeDot, 5),
		NewTile(TileTypeDot, 6),
		NewTile(TileTypeDot, 7),
		NewTile(TileTypeDot, 8),
		NewTile(TileTypeDot, 8),
		NewTile(TileTypeDot, 9),
	}
	tile := NewTile(TileTypeWind, 2)
	ai.AcceptTiles(hand)
	ai.AcceptTiles(list)

	after := ai.Peng(tile)
	should := NewTile(TileTypeDragon, 2)

	assert.Equal(t, after, should)
	assert.Equal(t, ai.FinishedTiles, []TileList{append(list, tile)})
}

func TestGamePlayerAI_Peng2(t *testing.T) {
	ai := newAIPlayer()

	list := TileList{
		NewTile(TileTypeCharacter, 6),
		NewTile(TileTypeCharacter, 6),
	}
	shouldBe := NewTile(TileTypeWind, 3)

	hand := TileList{
		NewTile(TileTypeWind, 4),
		NewTile(TileTypeCharacter, 2),
		NewTile(TileTypeCharacter, 4),
		NewTile(TileTypeCharacter, 8),
		NewTile(TileTypeCharacter, 9),
		NewTile(TileTypeBamboo, 2),
		NewTile(TileTypeBamboo, 3),
		NewTile(TileTypeBamboo, 9),
		NewTile(TileTypeDot, 1),
		NewTile(TileTypeDot, 4),
	}
	tile := NewTile(TileTypeCharacter, 6)
	ai.AcceptTiles(TileList{shouldBe})
	ai.AcceptTiles(hand)
	ai.AcceptTiles(list)

	after := ai.Peng(tile)
	should := NewTile(TileTypeWind, 3)

	fmt.Println("打出:", after)
	fmt.Println("固定牌堆:", ai.FinishedTiles)
	fmt.Println("手牌:", ai.HandTiles)
	assert.Equal(t, after, should)
	assert.Equal(t, ai.FinishedTiles, []TileList{append(list, tile)})
	assert.Equal(t, ai.HandTiles, hand)
}
