package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		sum += arrangements(fields[0], toList(fields[1]))
	}

	fmt.Println(sum)
}

func arrangements(s string, num []int) int {
	s = strings.Trim(s, ".")

	if s == "" {
		if len(num) == 0 {
			return 1
		}
		return 0
	}

	if len(num) == 0 {
		if strings.Contains(s, "#") {
			return 0
		}
		return 1
	}

	sum := 0

	if s[0] == '?' {
		// check recursively whether the first is '.'
		sum += arrangements(s[1:], num)

		// and now consider whether the first is '#'
		s = "#" + s[1:]
	}

	if len(s) < num[0] {
		return sum
	}

	if strings.ContainsRune(s[:num[0]], '.') {
		return sum
	}

	if len(s) > num[0] {
		switch s[num[0]] {
		case '#':
			return sum
		case '?':
			s = s[:num[0]] + "." + s[num[0]+1:]
		}
	}

	return sum + arrangements(s[num[0]:], num[1:])
}
