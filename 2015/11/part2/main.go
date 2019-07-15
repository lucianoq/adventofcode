package main

import "fmt"

const input = "hepxxyzz"

func main() {
	s := input

	fmt.Println(NextValid(s))
}

func NextValid(s string) string {
	for {
		s = Next(s)
		if Valid(s) {
			return s
		}
	}
}

func Valid(s string) bool {
	return Contains3Straight(s) && NotContainsIOL(s) && Contains2Pairs(s)
}

func Next(s string) string {
	if len(s) == 0 {
		return "a"
	}

	last := s[len(s)-1] + 1
	if last <= 'z' {
		return s[:len(s)-1] + string(last)
	}

	return Next(s[:len(s)-1]) + string('a')
}

// Passwords must include one increasing straight of at least three letters,
// like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd
// doesn't count.
func Contains3Straight(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
			return true
		}
	}
	return false
}

// Passwords may not contain the letters i, o, or l, as these letters can be
// mistaken for other characters and are therefore confusing.
func NotContainsIOL(s string) bool {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'i', 'o', 'l':
			return false
		}
	}
	return true
}

// Passwords must contain at least two different, non-overlapping pairs of
// letters, like aa, bb, or zz.
func Contains2Pairs(s string) bool {
	numPairs := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			numPairs++
			if numPairs >= 2 {
				return true
			}
			i++
			continue
		}
	}
	return false
}
