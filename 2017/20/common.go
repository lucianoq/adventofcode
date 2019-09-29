package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Triple struct {
	X, Y, Z int
}

type Particle struct {
	ID  int
	Pos Triple
	Vel Triple
	Acc Triple
}

func (p Particle) PositionAfter(t int) Triple {
	var (
		posX, posY, posZ = p.Pos.X, p.Pos.Y, p.Pos.Z
		velX, velY, velZ = p.Vel.X, p.Vel.Y, p.Vel.Z
	)

	for i := 0; i < t; i++ {
		velX += p.Acc.X
		velY += p.Acc.Y
		velZ += p.Acc.Z
		posX += velX
		posY += velY
		posZ += velZ
	}
	return Triple{posX, posY, posZ}
}

func parse() []Particle {
	particles := make([]Particle, 0)
	scanner := bufio.NewScanner(os.Stdin)

	id := 0
	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Split(line, ", ")

		position := extractTriple(strings.TrimLeft(ff[0], "p="))
		velocity := extractTriple(strings.TrimLeft(ff[1], "v="))
		acceleration := extractTriple(strings.TrimLeft(ff[2], "a="))

		particles = append(particles, Particle{
			ID:  id,
			Pos: position,
			Vel: velocity,
			Acc: acceleration,
		})
		id++
	}

	return particles
}

func extractTriple(s string) Triple {
	s = strings.Trim(s, "<>")
	ff := strings.Split(s, ",")

	x, err := strconv.Atoi(ff[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(ff[1])
	if err != nil {
		log.Fatal(err)
	}
	z, err := strconv.Atoi(ff[2])
	if err != nil {
		log.Fatal(err)
	}

	return Triple{x, y, z}
}

func Manhattan(triple Triple) int {
	return abs(triple.X) + abs(triple.Y) + abs(triple.Z)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
