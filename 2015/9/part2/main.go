package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	g := NewGraph()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line != "" {
			var n1, n2 string
			var d int

			n, err := fmt.Sscanf(line, "%s to %s = %d", &n1, &n2, &d)
			if err != nil || n != 3 {
				log.Fatal(err)
			}

			g.AddEdge(n1, n2, d)
		}
	}

	fmt.Println(g.MaxPath())
}
