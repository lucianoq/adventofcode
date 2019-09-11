package main

type Direction uint8

func (d *Direction) RotateLeft90() {
	*d = (*d + 1) % 4
}

const (
	East Direction = iota
	North
	West
	South
)

type Coord struct {
	X, Y int
}

func (c *Coord) Move(d Direction) {
	switch d {
	case North:
		c.Y++
	case East:
		c.X++
	case West:
		c.X--
	case South:
		c.Y--
	}
}
