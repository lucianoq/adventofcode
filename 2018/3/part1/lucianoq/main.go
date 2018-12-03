package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const size int = 1000

func main() {
	var fabric [size][size]int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var id, posX, posY, w, h int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &posX, &posY, &w, &h)
		for i := posX; i < posX+w; i++ {
			for j := posY; j < posY+h; j++ {
				fabric[i][j] ++
			}
		}
	}

	var count int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if fabric[i][j] >= 2 {
				count ++
			}
		}
	}
	fmt.Println(count)
}
