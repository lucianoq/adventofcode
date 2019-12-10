package main

import "fmt"

func main() {
	asteroids := parse()

	max := 0
	for k := range asteroids {
		visible := countVisible(asteroids, k)
		if visible > max {
			max = visible
		}
	}

	fmt.Println(max)
}

func countVisible(area map[C]struct{}, subj C) int {
	visibleSet := make(map[C]struct{})
	for k := range area {
		if k != subj {
			dx := k.x - subj.x
			dy := k.y - subj.y
			visibleSet[irreducibleFraction(C{dx, dy})] = struct{}{}
		}
	}
	return len(visibleSet)
}

func irreducibleFraction(c C) C {
	// abs needed because otherwise lose
	// the difference between points on same line
	// but different directions.
	d := abs(gcd(c.x, c.y))
	return C{c.x / d, c.y / d}
}

// Euclidean algorithm for gcd
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
