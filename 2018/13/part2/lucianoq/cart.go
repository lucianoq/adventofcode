package main

type Direction int

const (
	null Direction = iota
	N
	E
	S
	W
)

type Intent int

const (
	Left Intent = iota
	Straight
	Right
)

func (i Intent) Next() Intent {
	return Intent((int(i) + 1) % 3)
}

type Cart struct {
	X         int
	Y         int
	Direction Direction
	NextTurn  Intent
}

func (c *Cart) Step() {
	switch c.Direction {
	case N:
		c.Y--
	case E:
		c.X++
	case S:
		c.Y++
	case W:
		c.X--
	}

	newCell := grid[c.X][c.Y]
	if newCell.Char == "/" {
		switch c.Direction {
		case N:
			c.Direction = E
		case W:
			c.Direction = S
		case S:
			c.Direction = W
		case E:
			c.Direction = N
		}
	}
	if newCell.Char == `\` {
		switch c.Direction {
		case N:
			c.Direction = W
		case W:
			c.Direction = N
		case S:
			c.Direction = E
		case E:
			c.Direction = S
		}
	}
	if newCell.Char == "+" {
		switch c.Direction {
		case N:
			switch c.NextTurn {
			case Left:
				c.Direction = W
			case Right:
				c.Direction = E
			}
		case E:
			switch c.NextTurn {
			case Left:
				c.Direction = N
			case Right:
				c.Direction = S
			}
		case W:
			switch c.NextTurn {
			case Left:
				c.Direction = S
			case Right:
				c.Direction = N
			}
		case S:
			switch c.NextTurn {
			case Left:
				c.Direction = E
			case Right:
				c.Direction = W
			}
		}
		c.NextTurn = c.NextTurn.Next()
	}
}

func (c *Cart) String() string {
	switch c.Direction {
	case N:
		return "^"
	case S:
		return "v"
	case W:
		return "<"
	case E:
		return ">"
	}
	return "?"
}
