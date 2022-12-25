package main

import "strconv"

func valid(x int) bool {
	str := strconv.Itoa(x)

	double := false

	for i := 0; i < len(str)-1; i++ {
		if str[i] > str[i+1] {
			return false
		}
		if str[i] == str[i+1] {
			double = true
		}
	}
	return double
}
