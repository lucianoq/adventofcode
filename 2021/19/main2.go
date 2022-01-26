package main

import "fmt"

func main() {
	aligned := alignAll(parse())

	var maxDist int32
	for _, s1 := range aligned {
		for _, s2 := range aligned {
			dist := manhattan(s1.position, s2.position)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	fmt.Println(maxDist)
}

func manhattan(b1, b2 Point) int32 {
	return abs(b1[0]-b2[0]) + abs(b1[1]-b2[1]) + abs(b1[2]-b2[2])
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
