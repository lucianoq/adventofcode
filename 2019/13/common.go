package main

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
