package main

import "fmt"

const (
	Min = 200000000000000
	Max = 400000000000000
)

func main() {
	hail := parse()

	count := 0
	for i := 0; i < len(hail)-1; i++ {
		for j := i + 1; j < len(hail); j++ {
			p, intersect := findIntersection(hail[i], hail[j])

			if intersect {
				if p.x >= Min && p.x <= Max && p.y >= Min && p.y <= Max {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func findIntersection(h1, h2 HailStone) (P, bool) {
	// (1)  xf = x0 + vx*t   ==>   t = (xf-x0)/vx
	// (2)  yf = y0 + vy*t
	//
	// Combining (1) and (2)
	// yf = y0 + vy * (xf-x0)/vx
	// yf = y0 + (vy/vx)(xf-x0)
	// yf = (vy/vx) * xf + [ y0 - (vy/vx)*x0 ]
	// y  =    m    *  x +        q
	//
	// m = vy/vx
	// q = y0 - m*x0

	m1 := h1.Vel.y / h1.Vel.x
	m2 := h2.Vel.y / h2.Vel.x

	if m1 == m2 {
		// parallel lines
		return P{}, false
	}

	q1 := h1.Pos.y - m1*h1.Pos.x
	q2 := h2.Pos.y - m2*h2.Pos.x

	xf := (q2 - q1) / (m1 - m2)
	yf := m1*xf + q1

	// they meet in (xf,yf)
	// need to check if this was in the past

	// From (1)
	t1 := (xf - h1.Pos.x) / h1.Vel.x
	if t1 < 0 {
		// past
		return P{}, false
	}

	// From (1)
	t2 := (yf - h2.Pos.y) / h2.Vel.y
	if t2 < 0 {
		// past
		return P{}, false
	}

	return P{xf, yf, 0}, true
}
