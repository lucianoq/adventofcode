package main

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)
const (
	//input = "hijkl" //example for hashing
	input = "pgflpeqp"
	//input = "ihgpwlah" // example1  DDRRRD
	//input = "kglvqrro" // example2  DDUDRLRRUDRD
	//input = "ihgpwlah" // example3  DRURDRUDDLLDLUURRDULRLDUUDDDRR
)

func isOpen(path string) [4]bool {
	hash := md5.Sum([]byte(input + path))
	str := hex.EncodeToString(hash[:])
	var open [4]bool
	for i := 0; i < 4; i++ {
		open[i] = str[i] > 'a'
	}
	return open
}

type Point struct {
	X, Y int
	Path string
}

func adjacents(p Point) []Point {
	open := isOpen(p.Path)

	var pp []Point
	if p.X > 0 && open[UP] {
		pp = append(pp, Point{p.X - 1, p.Y, p.Path + "U"})
	}
	if p.X < 3 && open[DOWN] {
		pp = append(pp, Point{p.X + 1, p.Y, p.Path + "D"})
	}
	if p.Y > 0 && open[LEFT] {
		pp = append(pp, Point{p.X, p.Y - 1, p.Path + "L"})
	}
	if p.Y < 3 && open[RIGHT] {
		pp = append(pp, Point{p.X, p.Y + 1, p.Path + "R"})
	}
	return pp
}
