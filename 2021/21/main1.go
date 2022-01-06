package main

import "fmt"

type Player struct {
	Position int
	Score    int
}

func (p *Player) Forward(steps int) {
	p.Position = (p.Position + steps) % 10
	if p.Position == 0 {
		p.Position = 10
	}
	p.Score += p.Position
}

func main() {
	p1 := Player{Position: 5}
	p2 := Player{Position: 8}

	rolls := 0
	dice := func() int {
		rolls++

		face := rolls % 100
		if face == 0 {
			return 100
		}
		return face
	}

	for {
		steps := dice() + dice() + dice()
		p1.Forward(steps)

		if p1.Score >= 1000 {
			fmt.Println(rolls * p2.Score)
			return
		}

		steps = dice() + dice() + dice()
		p2.Forward(steps)

		if p2.Score >= 1000 {
			fmt.Println(rolls * p1.Score)
			return
		}
	}
}
