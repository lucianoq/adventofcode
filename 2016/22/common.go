package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Device struct {
	X, Y              int
	Size, Used, Avail int
}

func parse() []Device {
	var devices []Device

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "/dev/grid") {
			ff := strings.Fields(line)

			x, y := extractCoordinates(ff[0])
			size, _ := strconv.Atoi(strings.TrimRight(ff[1], "T"))
			used, _ := strconv.Atoi(strings.TrimRight(ff[2], "T"))
			avail, _ := strconv.Atoi(strings.TrimRight(ff[3], "T"))
			devices = append(devices, Device{
				X:     x,
				Y:     y,
				Size:  size,
				Used:  used,
				Avail: avail,
			})
		}
	}
	return devices
}

func extractCoordinates(s string) (int, int) {
	ff := strings.Split(s, "-")
	x, _ := strconv.Atoi(strings.TrimLeft(ff[1], "x"))
	y, _ := strconv.Atoi(strings.TrimLeft(ff[2], "y"))
	return x, y
}

func Print(devs []Device) {
	type Point struct{ X, Y int }

	m := make(map[Point]Device)

	maxX, maxY := -1, -1
	for _, d := range devs {
		m[Point{X: d.X, Y: d.Y}] = d

		if d.X > maxX {
			maxX = d.X
		}
		if d.Y > maxY {
			maxY = d.Y
		}
	}

	fmt.Print("   ")
	for i := 0; i <= maxX; i++ {
		fmt.Printf("%6d ", i)
	}
	fmt.Println()

	for j := 0; j <= maxY; j++ {
		fmt.Printf("%3d ", j)
		for i := 0; i <= maxX; i++ {
			if d, ok := m[Point{X: i, Y: j}]; ok {
				fmt.Printf("%3d/%2d ", d.Used, d.Size)
			} else {
				fmt.Println("wtf")
			}
		}
		fmt.Println()
	}
}
