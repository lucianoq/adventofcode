package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type P struct {
	C, R int
}

type Rectangle struct {
	TL, BR P
}

func NewRectangle(a, b P) Rectangle {
	return Rectangle{
		P{min(a.C, b.C), min(a.R, b.R)},
		P{max(a.C, b.C), max(a.R, b.R)},
	}
}

func (r Rectangle) Area() int {
	//return (abs(a.C-b.C) + 1) * (abs(a.R-b.R) + 1)
	return (r.BR.C - r.TL.C + 1) * (r.BR.R - r.TL.R + 1)
}

func parseInput() []P {
	var list []P
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Split(line, ",")
		c, _ := strconv.Atoi(ff[0])
		r, _ := strconv.Atoi(ff[1])
		list = append(list, P{c, r})
	}
	return list
}
