package main

import "fmt"

type Satellite struct {
	x, y, z    int
	vx, vy, vz int
}

func main() {
	ss := [4]*Satellite{
		{x: -13, y: 14, z: -7},
		{x: -18, y: 9, z: 0},
		{x: 0, y: -3, z: -3},
		{x: -15, y: 3, z: -13},
	}

	for step := 1; step <= 1000; step++ {

		// for each pair
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				a, b := ss[i], ss[j]
				switch {
				case a.x < b.x:
					a.vx++
					b.vx--
				case a.x > b.x:
					a.vx--
					b.vx++
				}
				switch {
				case a.y < b.y:
					a.vy++
					b.vy--
				case a.y > b.y:
					a.vy--
					b.vy++
				}
				switch {
				case a.z < b.z:
					a.vz++
					b.vz--
				case a.z > b.z:
					a.vz--
					b.vz++
				}
			}
		}

		// apply velocity
		for i := 0; i < 4; i++ {
			ss[i].x += ss[i].vx
			ss[i].y += ss[i].vy
			ss[i].z += ss[i].vz
		}
	}

	var energy int
	for i := 0; i < 4; i++ {
		energy += (abs(ss[i].x) + abs(ss[i].y) + abs(ss[i].z)) * (abs(ss[i].vx) + abs(ss[i].vy) + abs(ss[i].vz))
	}
	fmt.Println(energy)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
