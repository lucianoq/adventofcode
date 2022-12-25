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
			notThreeLeft := i == 0 || str[i] != str[i-1]
			notThreeRight := i == len(str)-2 || str[i] != str[i+2]
			if notThreeLeft && notThreeRight {
				double = true
			}
		}
	}

	return double
}
