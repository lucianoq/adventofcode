package main

import "fmt"

type Satellite struct {
	pos, vel int
}

func main() {
	ssx := [4]*Satellite{{pos: -13}, {pos: -18}, {pos: 0}, {pos: -15}}
	ssy := [4]*Satellite{{pos: 14}, {pos: 9}, {pos: -3}, {pos: 3}}
	ssz := [4]*Satellite{{pos: -7}, {pos: 0}, {pos: -3}, {pos: -13}}

	xPeriod := findPeriod(ssx)
	yPeriod := findPeriod(ssy)
	zPeriod := findPeriod(ssz)

	fmt.Println(lcm(xPeriod, lcm(yPeriod, zPeriod)))
}

func findPeriod(ss [4]*Satellite) int {
	visited := make(map[[4]Satellite]struct{})

	for step := 0; ; step++ {

		s := toStruct(ss)
		if _, ok := visited[s]; ok {
			return step
		}
		visited[s] = struct{}{}

		// for each pair
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				a, b := ss[i], ss[j]
				switch {
				case a.pos < b.pos:
					a.vel++
					b.vel--
				case a.pos > b.pos:
					a.vel--
					b.vel++
				}
			}
		}

		// apply velocity
		for i := 0; i < 4; i++ {
			ss[i].pos += ss[i].vel
		}
	}
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func toStruct(ss [4]*Satellite) [4]Satellite {
	var res [4]Satellite
	for i := 0; i < 4; i++ {
		res[i] = *ss[i]
	}
	return res
}
