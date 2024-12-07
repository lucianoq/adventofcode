package main

import "strconv"

func valid(goal, temp int, others []int) bool {
	if len(others) == 0 {
		return temp == goal
	}

	if valid(goal, temp+others[0], others[1:]) {
		return true
	}
	if valid(goal, temp*others[0], others[1:]) {
		return true
	}

	temp, _ = strconv.Atoi(strconv.Itoa(temp) + strconv.Itoa(others[0]))
	return valid(goal, temp, others[1:])
}
