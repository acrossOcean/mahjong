package main

import (
	"errors"
	"fmt"
)

type GameManager struct {
	// 上次牌堆引用
	lastTileStack GameTileStack
	// 本次牌堆引用
	tileStack *GameTileStack
	// 玩家引用
	playerList []IGamePlayer
	// 当前庄家
	currentDealerIndex int
	// 下次抓牌人
	nextPlayerCount int
	// 牌组剩余数量(实际剩余数量)
	realTileCount map[Tile]int
	// 牌面上剩余数量 (不包括手牌, 暗杠的牌等)
	tileCount map[Tile]int
}

func NewGameManager() *GameManager {
	result := new(GameManager)
	result.currentDealerIndex = -1
	result.tileCount = make(map[Tile]int)
	result.realTileCount = make(map[Tile]int)

	allTile := DefaultAllTiles()
	allTile = append(allTile, allTile...)
	allTile = append(allTile, allTile...)
	result.tileStack = NewGameTileStack(allTile.Shuffle())

	for _, tile := range allTile {
		result.tileCount[tile]++
		result.realTileCount[tile]++
	}

	return result
}

func NewGameManagerWithStack(stack GameTileStack) *GameManager {
	result := new(GameManager)
	result.currentDealerIndex = -1
	result.tileCount = make(map[Tile]int)
	result.realTileCount = make(map[Tile]int)

	result.lastTileStack = stack
	result.tileStack = NewGameTileStack(stack.GetTileList().Shuffle())

	for _, tile := range result.tileStack.GetTileList() {
		result.tileCount[tile]++
		result.realTileCount[tile]++
	}

	return result
}

func (receiver *GameManager) AddPlayer(player IGamePlayer) error {
	if len(receiver.playerList) >= 4 {
		// 最多4个玩家, 多余的不算
		return errors.New(ErrPlayerIsFull)
	}

	receiver.playerList = append(receiver.playerList, player)
	return nil
}

// 刷新游戏数据, 开始新的一局游戏
func (receiver GameManager) NewGame() {
	fmt.Println("新游戏")
	// 如果当前没有牌, 那么生成新的牌局
	if receiver.tileStack == nil {
		fmt.Println("生成牌堆")
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
		receiver.lastTileStack = *receiver.tileStack
	}

	// 如果当前没有庄家, 抓牌决定庄家, 返回新的玩家列表, 按照抓风顺序来排列
	if receiver.currentDealerIndex < 0 {
		receiver.ShufflePlayer()
	}

	if len(receiver.playerList) <= 0 {
		fmt.Println("没有玩家, 结束")
		return
	}

	// 确定庄家
	receiver.currentDealerIndex = receiver.nextPlayerCount

	// 开始最初的13张牌抓牌
	receiver.beginDraw()

	// 开始流程,直到有玩家胡牌或者牌全部抓完
	receiver.startGame()
}

// 玩家抓风, 随机一个风字牌的排列, 按照顺序安排给玩家
func (receiver *GameManager) ShufflePlayer() {
	var tile Tile
	list := DefaultAllWinds().Shuffle()

	newPlayerList := make([]IGamePlayer, len(receiver.playerList))
	for i, player := range receiver.playerList {
		tile, list = list.GetOneFromHead()
		player.SetLocation(tile)
		newPlayerList[i] = player
	}

	// 按照顺序排列玩家
	receiver.playerList = newPlayerList
}

func (receiver *GameManager) beginDraw() {
	fmt.Println("初始牌堆:", receiver.tileStack.list)

	// 先一人抓 4 张, 抓 3 轮
	for i := 0; i < 3; i++ {
		for _, player := range receiver.playerList {
			list := receiver.tileStack.GetFromHead(4)
			player.AcceptTiles(list)

			// 从计数中去除
			for _, tile := range list {
				receiver.realTileCount[tile]--
			}

			fmt.Println("玩家", player, "得到:", list)
		}
	}

	// 然后每人抓 1 张
	for _, player := range receiver.playerList {
		list := receiver.tileStack.GetFromHead(1)
		player.AcceptTiles(list)

		// 从计数中去除
		for _, tile := range list {
			receiver.realTileCount[tile]--
		}

		fmt.Println("玩家", player, "得到:", list)
	}

	fmt.Println("牌堆剩余:", receiver.tileStack.list)
	for _, player := range receiver.playerList {
		fmt.Println("玩家", player, "手牌:", player.GetAllTiles())
		//fmt.Println("玩家", player, "打出:", player.SendTile())
	}
}

func (receiver *GameManager) startGame() {
	needDraw := true
	var tile Tile
	var nextOper PlayerOperation
	hasOper := false

	for {
		if receiver.tileStack.GetTileList().Len() <= 0 {
			fmt.Println("牌堆无牌, 流局")
			break
		}

		// 获取当前玩家
		player := receiver.playerList[receiver.nextPlayerCount]
		playerWin := false
		var newTile Tile

		fmt.Println("轮到玩家:", player, "当前手牌:", player.GetHandTiles())

		if needDraw {
			// 抓牌
			tile = receiver.tileStack.GetFromHead(1)[0]
			receiver.realTileCount[tile]--
			fmt.Println("抓到:", tile)
		}
		needDraw = true

		// 询问当前玩家操作
		for {
			if nextOper == OperationPeng {
				newTile = player.Peng(tile)
				fmt.Println("玩家", player, "碰:", tile)
				nextOper = OperationNull
				hasOper = true
			} else {
				oper := player.IsNeed(tile)
				if oper == OperationWin {
					fmt.Println("玩家", player, "赢:", tile)
					playerWin = true
					hasOper = true
					break
				} else if oper == OperationGang {
					lastTile := receiver.tileStack.GetOneFromEnd()
					tile = player.Gang(tile, lastTile)
					newTile = tile
					hasOper = true
					fmt.Println("玩家", player, "杠:", tile)
				} else {
					player.AcceptTiles(TileList{tile})
					fmt.Println("当前手牌:", player.GetAllTiles())
					newTile = player.SendTile()
					break
				}
			}
		}

		fmt.Println("玩家:", player, "打出:", newTile)
		fmt.Println("打牌后:", player.GetAllTiles())

		//break
		if playerWin {
			fmt.Println("游戏结束, 玩家:", player, "获胜")
			break
		}

		fmt.Println("玩家:", player, "手牌为:", player.GetHandTiles())

		if hasOper {
			break
		}

		winPlayerCounts := make([]int, 0)
		operPlayer := -1

		// 轮询, 看有没有需要的玩家
		for index := (receiver.nextPlayerCount + 1) % len(receiver.playerList); index != receiver.nextPlayerCount; index = (index + 1) % len(receiver.playerList) {
			player := receiver.playerList[index]

			operation := player.IsNeed(newTile)
			if operation == OperationWin {
				winPlayerCounts = append(winPlayerCounts, index)
			} else if (operPlayer == -1) && (operation == OperationGang || operation == OperationPeng) {
				operPlayer = index
				nextOper = operation
			}
		}

		if len(winPlayerCounts) > 0 {
			// 有玩家可以胜利, 结束
			fmt.Println("游戏结束")
			fmt.Print("获胜玩家:")
			for _, count := range winPlayerCounts {
				fmt.Print(receiver.playerList[count], ";")
			}
			break
		} else if operPlayer != -1 {
			// 有玩家选择了操作, 下个玩家直接选
			tile = newTile
			receiver.nextPlayerCount = operPlayer
			needDraw = false
			fmt.Println("玩家:", receiver.playerList[receiver.nextPlayerCount], "发生操作:", nextOper)
		} else {
			receiver.nextPlayerCount = (receiver.nextPlayerCount + 1) % len(receiver.playerList)
		}
	}

	receiver.currentDealerIndex = receiver.NextDealer()
}

func (receiver *GameManager) DrawPlayerAllInfo() {
	fmt.Println("当前牌堆:")
	fmt.Println(receiver.tileStack.GetTileList())

	for i, player := range receiver.playerList {
		fmt.Println("玩家:", i+1)
		fmt.Println(player.GetAllTiles())
	}
}

func (receiver *GameManager) NextDealer() int {
	return (receiver.currentDealerIndex + 1) % len(receiver.playerList)
}

func (receiver GameManager) getRestNum(tile Tile) int {
	return receiver.tileCount[tile]
}

func (receiver GameManager) getRestNumWithoutList(tile Tile, list TileList) int {
	count := receiver.getRestNum(tile)

	for _, t := range list {
		if t == tile {
			count--
		}
	}

	return count
}
