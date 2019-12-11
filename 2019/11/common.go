package main

type C struct{ X, Y int }

const (
	Up = iota
	Right
	Down
	Left
)

const (
	Black = iota
	White
)

const (
	TurnLeft = iota
	TurnRight
)

type Robot struct {
	Pos       C
	Direction int
}

func NewRobot() *Robot {
	return &Robot{
		Pos:       C{0, 0},
		Direction: Up,
	}
}

func (r *Robot) Turn(where int) {
	switch where {
	case TurnLeft:
		r.Direction = (r.Direction + 3) % 4
	case TurnRight:
		r.Direction = (r.Direction + 1) % 4
	}
}

func (r *Robot) MoveForward() {
	// move forward
	switch r.Direction {
	case Up:
		r.Pos.Y++
	case Right:
		r.Pos.X++
	case Down:
		r.Pos.Y--
	case Left:
		r.Pos.X--
	}
}

func (r *Robot) Paint(grid map[C]int, color int) {
	grid[r.Pos] = color
}

func run(grid map[C]int) {
	// 1 because we should be able to write input
	// although the program terminated and can't read it
	camera := make(chan int, 1)
	output := make(chan int, 0)

	go func() {
		NewVM("input", camera, output).Run()
		close(output)
	}()

	robot := NewRobot()

	for {
		// send color from camera
		camera <- grid[robot.Pos]

		// receive color
		newColor, open := <-output
		exit(!open, grid)

		// receive rotation
		turn := <-output

		robot.Paint(grid, newColor)
		robot.Turn(turn)
		robot.MoveForward()
	}
}
