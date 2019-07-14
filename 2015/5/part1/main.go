package part1

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if CountVowels(line) >= 3 && CountainsCoupleEqualChars(line) && NotContainsBadCouples(line) {
			i++
		}
	}
	fmt.Println(i)
}

func CountVowels(s string) (i int) {
	for _, value := range s {
		switch value {
		case 'a', 'e', 'i', 'o', 'u':
			i++
		}
	}
	return
}

func CountainsCoupleEqualChars(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func NotContainsBadCouples(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		switch s[i : i+2] {
		case "ab", "cd", "pq", "xy":
			return false
		}
	}
	return true
}
