package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

const Steps = 26501365

func main() {
	m := parse()

	// Find the 3 points necessary to generate the parabola coefficients

	var points [3]P

	for i, steps := range []int{
		m.Size / 2,          // reach the border, fill one tile
		m.Size/2 + m.Size,   // 5 tiles
		m.Size/2 + 2*m.Size, // 13 tiles
	} {
		cells := Reach(m, steps)
		points[i] = P{steps, cells}
	}

	// Given these 3 points, we can find a,b,c such as
	// y = a*x^2 + b*x + c
	a, b, c := FindParabola(points)

	fmt.Println(int(a*Steps*Steps + b*Steps + c))
}

func FindParabola(p [3]P) (float64, float64, float64) {
	A := mat.NewDense(3, 3, []float64{
		float64(p[0].x * p[0].x), float64(p[0].x), 1,
		float64(p[1].x * p[1].x), float64(p[1].x), 1,
		float64(p[2].x * p[2].x), float64(p[2].x), 1,
	})

	y := mat.NewVecDense(3, []float64{float64(p[0].y), float64(p[1].y), float64(p[2].y)})

	var b mat.VecDense
	_ = b.SolveVec(A, y)

	return b.AtVec(0), b.AtVec(1), b.AtVec(2)
}
