package main

func main() {
	gm := NewGameManager()
	p1 := new(CommonGamePlayer)
	gm.AddPlayer(p1)

	p2 := new(CommonGamePlayer)
	gm.AddPlayer(p2)

	p3 := new(CommonGamePlayer)
	gm.AddPlayer(p3)

	p4 := new(CommonGamePlayer)
	gm.AddPlayer(p4)

	gm.NewGame()
}
