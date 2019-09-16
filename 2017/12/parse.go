package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse() *Graph {
	scanner := bufio.NewScanner(os.Stdin)

	g := NewGraph()

	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Fields(line)

		from, err := strconv.Atoi(ff[0])
		if err != nil {
			log.Fatal(err)
		}
		g.AddNode(&Node{from})

		for _, f := range ff[2:] {
			f = strings.TrimRight(f, ",")
			recipient, err := strconv.Atoi(f)
			if err != nil {
				log.Fatal(err)
			}
			g.AddEdge(&Node{from}, &Node{recipient})
		}
	}
	return g
}
