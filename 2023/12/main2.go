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

		conditions := strings.Repeat(fields[0]+"?", 5)
		conditions = conditions[:len(conditions)-1]

		groupsStr := strings.Repeat(fields[1]+",", 5)
		groupsStr = groupsStr[:len(groupsStr)-1]

		sum += Cache{}.cachedArrangements(conditions, toList(groupsStr))
	}
	fmt.Println(sum)
}

type Cache map[string]int

func (c Cache) cachedArrangements(s string, num []int) int {
	s = strings.Trim(s, ".")

	key := s + fmt.Sprintf("%v", num)

	if val, ok := c[key]; ok {
		return val
	}

	result := c.arrangements(s, num)

	c[key] = result
	return result
}

func (c Cache) arrangements(s string, num []int) int {
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
		sum += c.cachedArrangements(s[1:], num)

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

	return sum + c.cachedArrangements(s[num[0]:], num[1:])
}
