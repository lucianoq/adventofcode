package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func parse() (int, int, int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()

	r := regexp.MustCompile("^target area: x=([0-9-]+)\\.\\.([0-9-]+), y=([0-9-]+)\\.\\.([0-9-]+)$")
	res := r.FindAllStringSubmatch(line, 1)

	x1, _ := strconv.Atoi(res[0][1])
	x2, _ := strconv.Atoi(res[0][2])
	y1, _ := strconv.Atoi(res[0][3])
	y2, _ := strconv.Atoi(res[0][4])

	return x1, x2, y1, y2
}
