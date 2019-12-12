package main

import (
	"fmt"
	"math"
	"sort"
	"time"
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
	Draw(asteroids, asteroid, Laser, true)
	time.Sleep(50 * time.Millisecond)
	delete(asteroids, asteroid)
	Draw(asteroids, asteroid, Laser, false)
	time.Sleep(10 * time.Millisecond)
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

func Draw(asteroids map[C]struct{}, target, laser C, loading bool) {
	fmt.Print("\033[2J\033[H")
	for j := 0; j < 21; j++ {
		for i := 0; i < 21; i++ {

			if i == laser.x && j == laser.y {
				fmt.Print(green("🔫 "))
				continue
			}

			if i == target.x && j == target.y {
				if loading {
					fmt.Print(red("★ "))
				} else {
					fmt.Print(red("  "))
				}
				continue
			}

			if _, ok := asteroids[C{i, j}]; ok {
				fmt.Print("★ ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func red(s string) string {
	return "\033[1;91m" + s + "\033[0m"
}

func green(s string) string {
	return "\033[1;92m" + s + "\033[0m"
}
