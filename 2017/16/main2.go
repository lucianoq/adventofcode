package main

import "fmt"

const numIteration = 1e9

func main() {
	cmds := parse()

	s := []byte(input)

	ring := []string{input}
	visited := map[string]struct{}{input: {}}

	for i := 1; i <= numIteration; i++ {
		for _, c := range cmds {
			c.Apply(s)
		}

		if _, ok := visited[string(s)]; ok {
			fmt.Println(ring[numIteration%len(visited)])
			return
		}

		visited[string(s)] = struct{}{}
		ring = append(ring, string(s))
	}
}
