package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Cuboid struct {
	state      bool
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
}

func parse() []Cuboid {
	exp := regexp.MustCompile("^([onf]+) x=([0-9-]+)\\.\\.([0-9-]+),y=([0-9-]+)\\.\\.([0-9-]+),z=([0-9-]+)\\.\\.([0-9-]+)$")

	cubes := []Cuboid{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if !exp.MatchString(line) {
			panic("wrong input")
		}

		ff := exp.FindAllStringSubmatch(line, 1)

		state := ff[0][1] == "on"
		minX, _ := strconv.Atoi(ff[0][2])
		maxX, _ := strconv.Atoi(ff[0][3])
		minY, _ := strconv.Atoi(ff[0][4])
		maxY, _ := strconv.Atoi(ff[0][5])
		minZ, _ := strconv.Atoi(ff[0][6])
		maxZ, _ := strconv.Atoi(ff[0][7])

		cubes = append(cubes, Cuboid{
			state: state,
			minX:  minX,
			maxX:  maxX,
			minY:  minY,
			maxY:  maxY,
			minZ:  minZ,
			maxZ:  maxZ,
		})
	}

	return cubes
}
