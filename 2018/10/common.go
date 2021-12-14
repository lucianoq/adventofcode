package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Star struct {
	X    int
	Y    int
	VelX int
	VelY int
}

func (s *Star) Current(time int) (x, y int) {
	return s.X + time*s.VelX, s.Y + time*s.VelY
}

var stars []*Star

func main() {
	stars = make([]*Star, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			line = strings.Replace(line, " ", "", -1)
			var x, y, velX, velY int
			n, err := fmt.Sscanf(line, "position=<%d,%d>velocity=<%d,%d>", &x, &y, &velX, &velY)
			if err != nil || n != 4 {
				log.Fatal(err)
			}
			stars = append(stars, &Star{x, y, velX, velY})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; ; i++ {
		Sky(i)
	}
}
