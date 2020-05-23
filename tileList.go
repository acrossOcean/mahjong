package main

import (
	"math/rand"
	"time"
)

//TileList 牌堆, 可代表手牌, 待抓牌, 已出牌 等各种组成排列的牌数组
type TileList []Tile

//Remove the n tile from list (start at 0)
func (receiver TileList) Remove(n int) (Tile, TileList) {
	if n < 0 {
		n = 0
	} else if n >= len(receiver) {
		n = len(receiver) - 1
	}

	result := receiver[n]
	r := receiver[:n]
	if n+1 < len(receiver) {
		r = append(r, receiver[n+1:]...)
	}
	return result, r
}

//Shuffle the tileList
func (receiver TileList) Shuffle() TileList {
	result := make([]Tile, receiver.Len())
	rand.Seed(time.Now().UnixNano())

	list := receiver
	for i := 0; i < len(result); i++ {
		result[i], list = list.Remove(rand.Intn(list.Len()))
	}

	return result
}

//GetOneFromHead get No.0 tile from list, and remove it from list
func (receiver TileList) GetOneFromHead() (Tile, TileList) {
	result, list := receiver.GetFromHead(1)
	return result[0], list
}

//GetOneFromEnd get No.(len(TileList) - 1) tile from list, and remove it from list
func (receiver TileList) GetOneFromEnd() (Tile, TileList) {
	result, list := receiver.GetFromEnd(1)
	return result[0], list
}

//GetFromHead get n tiles from list, and remove from list
func (receiver TileList) GetFromHead(n int) (TileList, TileList) {
	return receiver.getTiles(n, true)
}

//GetFromEnd get n tiles from end of this list, and remove from list
func (receiver TileList) GetFromEnd(n int) (TileList, TileList) {
	return receiver.getTiles(n, false)
}

//GetTiles get n tiles from head (if isHead is true) or end of this list, and remove from list
func (receiver TileList) getTiles(n int, isHead bool) (TileList, TileList) {
	if n < 0 {
		n = 0
	} else if n > receiver.Len() {
		n = receiver.Len()
	}

	var other TileList
	var result []Tile
	if isHead {
		result, other = receiver[0:n], receiver[n:]
	} else {
		index := len(receiver) - n
		other, result = receiver[0:index], receiver[index:]
	}

	return result, other
}

// ==============================================================
// implement sort.Interface
// ==============================================================

func (receiver TileList) Len() int { return len(receiver) }

func (receiver TileList) Less(i, j int) bool {
	if receiver[i].tileType != receiver[j].tileType {
		return receiver[i].tileType < receiver[j].tileType
	}

	return receiver[i].tileNum < receiver[j].tileNum
}

func (receiver TileList) Swap(i, j int) { receiver[i], receiver[j] = receiver[j], receiver[i] }
