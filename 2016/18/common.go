package main

const input = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"

func count(numLines int) int {
	safes := 0
	line := input

	for i := 0; i < numLines; i++ {
		safes += safe(line)
		line = next(line)
	}
	return safes
}

func next(s string) string {
	newS := ""
	for i := 0; i < len(s); i++ {

		// cases are independent from center
		var left, right bool

		if i-1 >= 0 {
			left = s[i-1] == '^'
		}

		if i+1 < len(s) {
			right = s[i+1] == '^'
		}

		if left != right {
			newS += "^"
		} else {
			newS += "."
		}

	}
	return newS
}

func safe(s string) int {
	count := 0
	for _, c := range s {
		if c == '.' {
			count++
		}
	}
	return count
}
