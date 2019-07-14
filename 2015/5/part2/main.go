package part2

import (
	"bufio"
	"fmt"
	"os"
)

//qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a
// letter that repeats with exactly one letter between them (zxz).
//xxyxx is nice because it has a pair that appears twice and a letter that
// repeats with one between, even though the letters used by each rule overlap.
//uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a
// single letter between them.
//ieodomkazucvgmuy is naughty because it has a repeating letter with one
// between (odo), but no pair that appears twice.

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 3 {
			if A(line) && B(line) {
				i++
			}
		}

	}
	fmt.Println(i)
}

// It contains a pair of any two letters that appears at least twice in the
// string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like
// aaa (aa, but it overlaps).
func A(s string) bool {
	// -1 because I'm considering pairs
	// -2 because first index stops when there is room for another pair with j.
	for i := 0; i < len(s)-1-2; i++ {
		for j := i + 2; j < len(s)-1; j++ {
			if s[i:i+2] == s[j:j+2] {
				return true
			}
		}
	}
	return false
}

// It contains at least one letter which repeats with exactly one letter between
// them, like xyx, abcdefeghi (efe), or even aaa.
func B(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

