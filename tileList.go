package main

import (
	"math/rand"
	"time"
)

//TileList 牌堆, 可代表手牌, 待抓牌, 已出牌 等各种组成排列的牌数组
type TileList []Tile

//Remove the n tile from list
func (receiver TileList) Remove(n int) Tile {
	result := receiver[n]
	r := receiver[:n-1]
	r = append(r, receiver[n:]...)
	receiver = r
	return result
}

//Shuffle the tileList
func (receiver TileList) Shuffle() TileList {
	result := make([]Tile, receiver.Len())
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(result); i++ {
		result[i] = receiver.Remove(rand.Intn(receiver.Len()))
	}

	return result
}

//GetOneFromHead get No.0 tile from list, and remove it from list
func (receiver TileList) GetOneFromHead() Tile {
	return receiver.GetFromHead(1)[0]
}

//GetOneFromEnd get No.(len(TileList) - 1) tile from list, and remove it from list
func (receiver TileList) GetOneFromEnd() Tile {
	return receiver.GetFromEnd(1)[0]
}

//GetFromHead get n tiles from list, and remove from list
func (receiver TileList) GetFromHead(n int) []Tile {
	return receiver.getTiles(n, true)
}

//GetFromEnd get n tiles from end of this list, and remove from list
func (receiver TileList) GetFromEnd(n int) []Tile {
	return receiver.getTiles(n, false)
}

//GetTiles get n tiles from head (if isHead is true) or end of this list, and remove from list
func (receiver TileList) getTiles(n int, isHead bool) []Tile {
	if n > receiver.Len() {
		n = receiver.Len()
	}

	var result []Tile
	if isHead {
		result, receiver = receiver[0:n], receiver[n:]
	} else {
		receiver, result = receiver[0:n], receiver[n:]
	}

	return result
}

// ==============================================================
// implement sort.Interface
// ==============================================================

func (receiver TileList) Len() int { return len(receiver) }

func (receiver TileList) Less(i, j int) bool {
	if receiver[i].tileType != receiver[j].tileType {
		return receiver[i].tileType > receiver[j].tileType
	}

	return receiver[i].tileNum > receiver[j].tileNum
}

func (receiver TileList) Swap(i, j int) { receiver[i], receiver[j] = receiver[j], receiver[i] }
