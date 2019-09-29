package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const size int = 1000

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

var fabric [size][size]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	claims := make([]*claim, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		c := &claim{}
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.width, &c.height)
		claims = append(claims, c)
	}

	for _, c := range claims {
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				fabric[i][j]++
			}
		}
	}

	for _, c := range claims {
		if alone(c) {
			fmt.Println(c.id)
			os.Exit(0)
		}
	}
	os.Exit(-1)
}

func alone(c *claim) bool {
	for i := c.x; i < c.x+c.width; i++ {
		for j := c.y; j < c.y+c.height; j++ {
			if fabric[i][j] != 1 {
				return false
			}
		}
	}
	return true
}
