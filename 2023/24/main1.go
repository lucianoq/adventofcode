package main

import (
	"fmt"
	"log"
)

// const (
//
//	xMin, xMax = 200000000000000, 400000000000000
//	yMin, yMax = 200000000000000, 400000000000000
//
// )
const (
	xMin, xMax = 7, 27
	yMin, yMax = 7, 27
)

func main() {
	hail := parseInput()

	within := 0
	for i := 0; i < len(hail)-1; i++ {
		for j := i + 1; j < len(hail); j++ {
			h1, h2 := hail[i], hail[j]
			if collide(h1, h2) {
				within++
			}
		}
	}

	fmt.Println(within)

	//25860 too high
}

func inside(axis byte, point float64) bool {
	switch axis {
	case 'x':
		return point >= xMin && point <= xMax

	case 'y':
		return point >= yMin && point <= yMax

	case 'z':
		return true
	}

	return false
}

func collide(h1, h2 Hail) bool {

	fmt.Println()
	fmt.Println(h1)
	fmt.Println(h2)

	whereX, cX := collideAxis('x', h1, h2)
	if cX {
		if inside('x', whereX) {
			log.Printf("x intersects at %.2f, inside", whereX)
			whereY, cY := collideAxis('y', h1, h2)
			if cY {
				if inside('y', whereY) {
					log.Printf("y intersects at %.2f, inside", whereY)
					return true
				} else {
					log.Printf("y intersects at %.2f, outside", whereY)
					return false
				}
			} else {
				log.Println("y does not intersect")
				return false
			}
		} else {
			log.Printf("x intersects at %.2f, outside", whereX)
			return false
		}
	} else {
		log.Println("x does not intersect")
		return false
	}
}

//// h1 func
//// x = x0 + vx*t
//// y = y0 + vy*t
//
//// h2 func
//// x = x0 + vx*t
//// y = y0 + vy*t
//
//// { x = h1.vx*t + h1.x0
//// { x = h2.vx*t + h2.x0
//// ---------------------
////   0 = t * (h1.vx-h2.vx) + h1.x0 - h2.x0
////   t = h2.x0 - h1.x0
////       -------------
////       h1.vx - h2.vx
//
//var tx, ty int
//
//if h1.Vel.x == h2.Vel.x {
//	if h1.Pos.x != h2.Pos.x {
//		log.Printf("On x, Same velocity, different starting position: %d@%d vs %d@%d", h1.Pos.x, h1.Vel.x, h2.Pos.x, h2.Vel.x)
//		return P{}, false
//	}
//	log.Printf("On x, Same velocity, same starting position: %d@%d vs %d@%d", h1.Pos.x, h1.Vel.x, h2.Pos.x, h2.Vel.x)
//	log.Fatal("1")
//} else {
//	tx = (h2.Pos.x - h1.Pos.x) / (h1.Vel.x - h2.Vel.x)
//	if tx < 0 {
//		return P{}, false
//	}
//}
//
//if h1.Vel.y == h2.Vel.y {
//	if h1.Pos.y != h2.Pos.y {
//		log.Printf("On y, Same velocity, different starting position: %d@%d vs %d@%d", h1.Pos.y, h1.Vel.y, h2.Pos.y, h2.Vel.y)
//		return P{}, false
//	}
//	log.Printf("On y, Same velocity, same starting position: %d@%d vs %d@%d", h1.Pos.y, h1.Vel.y, h2.Pos.y, h2.Vel.y)
//	log.Fatal("2")
//} else {
//	ty = (h2.Pos.y - h1.Pos.y) / (h1.Vel.y - h2.Vel.y)
//	if ty < 0 {
//		return P{}, false
//	}
//}
//
//return tx == ty

func collideAxis(axis byte, h1 Hail, h2 Hail) (float64, bool) {
	var p1, p2 int
	var v1, v2 int

	switch axis {
	case 'x':
		p1, v1 = h1.Pos.x, h1.Vel.x
		p2, v2 = h2.Pos.x, h2.Vel.x

	case 'y':
		p1, v1 = h1.Pos.y, h1.Vel.y
		p2, v2 = h2.Pos.y, h2.Vel.y

	case 'z':
		p1, v1 = h1.Pos.z, h1.Vel.z
		p2, v2 = h2.Pos.z, h2.Vel.z
	}

	// x = p1+v1*t
	// x = p2+v2*t
	// p1 +v1*t = p2 +v2*t
	// t(v1-v2) = p2-p1
	// t = (p2-p1) / (v1-v2)
	//

	// x
	// x = p1 + v1*t
	// x = p2 + v2*t
	// 0 = p1-p2 + (v1-v2)t
	// t = (p2-p1) / (v1-v2)
	if v1 == v2 {
		if p1 == p2 {
			log.Printf("parallel, always intersecting")
			return float64(p1), true
		}
		log.Printf("parallel, never intersecting")
		return 0, false
	}

	t := float64(p2-p1) / float64(v1-v2)
	log.Printf("collide %s at t=%.2f", string(axis), t)
	return float64(p1) + float64(v1)*t, true
}
