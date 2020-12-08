package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	revRules := reverse(parse())
	visited := make(map[string]struct{})
	dfs("shiny gold", revRules, visited)

	// -1 to not count shiny gold itself
	fmt.Println(len(visited) - 1)
}

func dfs(item string, rules map[string][]string, visited map[string]struct{}) {
	if _, ok := visited[item]; ok {
		return
	}

	visited[item] = struct{}{}

	for _, c := range rules[item] {
		dfs(c, rules, visited)
	}
}

// reverse the graph from "contains" to "is contained"
func reverse(in map[string][]string) map[string][]string {
	out := make(map[string][]string)
	for k, v := range in {
		for _, it := range v {
			if out[it] == nil {
				out[it] = make([]string, 0, 1)
			}
			out[it] = append(out[it], k)
		}
	}
	return out
}

// create a graph and returns it as adjacency list of strings
func parse() map[string][]string {
	ruleRegex := regexp.MustCompile("^([a-z ]+) bags contain (.+)\\.$")
	contentRegex := regexp.MustCompile("^[0-9]+ ([a-z ]+) bag[s]?$")

	rules := map[string][]string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := ruleRegex.FindStringSubmatch(scanner.Text())

		subject, content := arr[1], arr[2]

		rules[subject] = make([]string, 0)

		if content == "no other bags" {
			continue
		}

		items := strings.Split(content, ", ")

		for _, it := range items {
			arr = contentRegex.FindStringSubmatch(it)
			rules[subject] = append(rules[subject], arr[1])
		}
	}
	return rules
}
