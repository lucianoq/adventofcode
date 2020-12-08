package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	rules := parse()

	fmt.Println(dfs("shiny gold", rules))
}

func dfs(item string, rules map[string]map[string]int) int {
	count := 0
	for k, v := range rules[item] {
		count += v + v*dfs(k, rules)
	}
	return count
}

func parse() map[string]map[string]int {
	ruleRegex := regexp.MustCompile("^([a-z ]+) bags contain (.+)\\.$")
	contentRegex := regexp.MustCompile("^([0-9]+) ([a-z ]+) bag[s]?$")

	rules := map[string]map[string]int{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		arr := ruleRegex.FindStringSubmatch(scanner.Text())

		subject, content := arr[1], arr[2]

		rules[subject] = make(map[string]int, 0)

		if content == "no other bags" {
			continue
		}

		items := strings.Split(content, ", ")

		for _, it := range items {
			arr = contentRegex.FindStringSubmatch(it)

			// create if not exist
			if _, ok := rules[subject]; !ok {
				rules[subject] = make(map[string]int)
			}
			num, _ := strconv.Atoi(arr[1])

			rules[subject][arr[2]] = num
		}
	}
	return rules
}
