package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

var (
	Laser = C{11, 11}
	Null  = C{-1, -1}
)

func main() {
	asteroids := parse()
	delete(asteroids, Laser)

	i := 1
	for len(asteroids) > 0 {
		left, right := split(asteroids)

		if len(right) > 0 {
			sortByAngleAndDistance(right)
			for target := right[0]; target != Null; target = pickOne(right, target) {
				disintegrate(asteroids, i, target)
				i++
			}
		}

		if len(left) > 0 {
			sortByAngleAndDistance(left)
			for target := left[0]; target != Null; target = pickOne(left, target) {
				disintegrate(asteroids, i, target)
				i++
			}
		}
	}
}

func (c C) angle() float64 {
	if c.x == Laser.x {
		return math.Inf(-1)
	}
	return float64(c.y-Laser.y) / float64(c.x-Laser.x)
}

// check if it is the 200th, otherwise remove from space
func disintegrate(asteroids map[C]struct{}, i int, asteroid C) {
	if i == 200 {
		fmt.Println(100*asteroid.x + asteroid.y)
		os.Exit(0)
	}
	delete(asteroids, asteroid)
}

// manhattan distance between two points
func manhattan(a, b C) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

// find the first asteroid with an angle greater than
// the one in input
func pickOne(asteroids []C, prev C) C {
	for _, a := range asteroids {
		if a.angle() > prev.angle() {
			return a
		}
	}
	return Null
}

func split(asteroids map[C]struct{}) ([]C, []C) {
	left := make([]C, 0)
	right := make([]C, 0)

	for a := range asteroids {
		switch {
		case a.x > Laser.x:
			right = append(right, a)
		case a.x < Laser.x:
			left = append(left, a)
		default:
			if a.y < Laser.y {
				right = append(right, a)
			} else {
				left = append(left, a)
			}
		}
	}
	return left, right
}

// ascending order by angle
// if equal, prefer closest to laser
func sortByAngleAndDistance(list []C) {
	sort.Slice(list, func(i, j int) bool {
		if list[i].angle() == list[j].angle() {
			return manhattan(list[i], Laser) < manhattan(list[j], Laser)
		}
		return list[i].angle() < list[j].angle()
	})
}
