package main

func valid(goal, temp int, others []int) bool {
	if len(others) == 0 {
		return temp == goal
	}

	if valid(goal, temp+others[0], others[1:]) {
		return true
	}

	return valid(goal, temp*others[0], others[1:])
}
