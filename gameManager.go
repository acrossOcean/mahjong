package main

import "errors"

type GameManager struct {
	// 上次牌堆引用
	lastTileStack GameTileStack
	// 本次牌堆引用
	tileStack *GameTileStack
	// 玩家引用
	playerList []GamePlayer
	// 当前庄家
	currentDealer GamePlayer
	// 下次抓牌人
	nextPlayerCount int
	// 牌组剩余数量
	tileCount map[Tile]int
}

func NewGameManager() *GameManager {
	result := new(GameManager)

	allTile := DefaultAllTiles()
	allTile = append(allTile, allTile...)
	allTile = append(allTile, allTile...)
	result.tileStack = NewGameTileStack(allTile.Shuffle())

	for _, tile := range allTile {
		result.tileCount[tile]++
	}

	return result
}

func NewGameManagerWithStack(stack GameTileStack) *GameManager {
	result := new(GameManager)

	result.lastTileStack = stack
	result.tileStack = NewGameTileStack(stack.GetTileList().Shuffle())

	for _, tile := range result.tileStack.GetTileList() {
		result.tileCount[tile]++
	}

	return result
}

func (receiver *GameManager) AddPlayer(player GamePlayer) error {
	if len(receiver.playerList) >= 4 {
		// 最多4个玩家, 多余的不算
		return errors.New(ErrPlayerIsFull)
	}

	receiver.playerList = append(receiver.playerList, player)
	return nil
}

// 刷新游戏数据, 开始新的一局游戏
func (receiver GameManager) NewGame() {
	// 如果当前没有牌, 那么生成新的牌局
	if receiver.tileStack == nil {
		var tileList TileList
		if receiver.lastTileStack.GetTileList().Len() == 0 {
			receiver.tileStack = NewEmptyGameTileStack()
			allTile := DefaultAllTiles()
			allTile = append(allTile, allTile...)
			allTile = append(allTile, allTile...)
			tileList = allTile.Shuffle()
		} else {
			tileList = receiver.lastTileStack.GetTileList()
		}

		receiver.tileStack = NewGameTileStack(tileList.Shuffle())
	}

	// 如果当前没有庄家, 抓牌决定庄家, 返回新的玩家列表, 按照抓风顺序来排列
	if receiver.currentDealer == nil {
		receiver.ShufflePlayer()
	}

	// 确定庄家
	receiver.currentDealer = receiver.playerList[receiver.nextPlayerCount]

	currentGameStartTileStack := receiver.tileStack
	// 开始最初的13张牌抓牌
	receiver.beginDraw()

	// 开始流程,直到有玩家胡牌或者牌全部抓完
	receiver.startGame()

	receiver.lastTileStack = *currentGameStartTileStack
}

//TODO: 玩家抓风, 根据玩家数量判断
func (receiver *GameManager) ShufflePlayer() {
}

//TODO: 把当前牌堆中的牌分发给玩家
func (receiver *GameManager) beginDraw() {

}

//TODO: 开始流程, 直到有玩家胡牌或者牌全部抓完退出
func (receiver *GameManager) startGame() {

}
