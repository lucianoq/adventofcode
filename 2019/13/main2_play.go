package main

import (
	"fmt"
	"time"
)

type C struct{ x, y int }

type Tile int

// 0 is an empty tile. No game object appears in this tile.
// 1 is a wall tile. Walls are indestructible barriers.
// 2 is a block tile. Blocks can be broken by the ball.
// 3 is a horizontal paddle tile. The paddle is indestructible.
// 4 is a ball tile. The ball moves diagonally and bounces off objects.
const (
	Empty Tile = iota
	Wall
	Block
	Paddle
	Ball
)

const (
	Noop  = 0
	Left  = -1
	Right = 1
)

func (t Tile) String() string {
	switch t {
	case Empty:
		return "  "
	case Wall:
		return "██"
	case Block:
		return "🎄" // 🛸" // "▒"
	case Paddle:
		return "🎅" // ▬▬" // __" // "▂"
	case Ball:
		return "⚽" // "+"
	}
	return ""
}

func main() {
	done := make(chan struct{})

	input := make(chan int, 0)
	output := make(chan int, 0)
	vm := NewVM("input", input, output)

	go func() {
		// play for free!
		vm.Code[0] = 2

		vm.Run()
		close(output)
	}()

	// output
	grid := make(map[C]Tile)

	var ballPos, paddlePos C

	// load grid
	for i := 0; i < 21*36; i++ {
		x, y, tile := <-output, <-output, Tile(<-output)
		grid[C{x, y}] = tile
		if tile == Ball {
			ballPos = C{x, y}
		}
		if tile == Paddle {
			paddlePos = C{x, y}
		}
	}

	go func() {

		var maxScore int

		// read changes
		for {
			redraw(grid)

			x, open := <-output
			if !open {
				fmt.Println(maxScore)
				done <- struct{}{}
				return
			}
			y := <-output

			if x == -1 && y == 0 {
				score := <-output
				if score > maxScore {
					maxScore = score
				}
				continue
			}

			tile := Tile(<-output)
			grid[C{x, y}] = tile

			if tile == Ball {
				ballPos = C{x, y}
			}
			if tile == Paddle {
				paddlePos = C{x, y}
			}
		}
	}()

	// input
	go func() {
		for {
			switch {
			case ballPos.x < paddlePos.x:
				input <- Left
			case ballPos.x > paddlePos.x:
				input <- Right
			case ballPos.x == paddlePos.x:
				input <- Noop
			}

			time.Sleep(40 * time.Millisecond)
		}
	}()

	<-done
}

func redraw(grid map[C]Tile) {
	fmt.Print("\033[2J\033[H")
	for j := 0; j < 21; j++ {
		for i := 0; i < 36; i++ {
			fmt.Print(grid[C{i, j}])
		}
		fmt.Print("\n\r")
	}
}
