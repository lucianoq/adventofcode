package main

import "fmt"

type Game struct {
	TurnPlayer1  bool
	Pos1, Score1 int
	Pos2, Score2 int
}

var dice = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

var uni1, uni2 = 0, 0

func main() {
	play(Game{TurnPlayer1: true, Pos1: 5, Pos2: 8}, 1)

	if uni1 > uni2 {
		fmt.Println(uni1)
	} else {
		fmt.Println(uni2)
	}
}

func play(g Game, qty int) {
	if g.Score1 >= 21 {
		uni1 += qty
		return
	}

	if g.Score2 >= 21 {
		uni2 += qty
		return
	}

	for steps, qtyDice := range dice {
		newGame := g
		if g.TurnPlayer1 {
			newPos1 := walk(g.Pos1, steps)
			newGame.TurnPlayer1 = false
			newGame.Pos1 = newPos1
			newGame.Score1 += newPos1
		} else {
			newPos2 := walk(g.Pos2, steps)
			newGame.TurnPlayer1 = true
			newGame.Pos2 = newPos2
			newGame.Score2 += newPos2
		}
		play(newGame, qty*qtyDice)
	}
}

func walk(pos, steps int) int {
	pos = (pos + steps) % 10
	if pos == 0 {
		return 10
	}
	return pos
}
