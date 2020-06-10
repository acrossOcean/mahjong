package main

func main() {
	gm := NewGameManager()
	p1 := NewAI(gm)
	gm.AddPlayer(p1)

	p2 := NewAI(gm)
	gm.AddPlayer(p2)

	p3 := NewAI(gm)
	gm.AddPlayer(p3)

	p4 := NewAI(gm)
	gm.AddPlayer(p4)

	gm.NewGame()
}
