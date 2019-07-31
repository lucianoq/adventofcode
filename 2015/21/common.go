package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

type Player struct {
	HP     int
	Damage int
	Armor  int
}

var weapons = []Item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armors = []Item{
	{"-", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []Item{
	{"-", 0, 0, 0},
	{"-", 0, 0, 0},
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

func win(a, b Player) bool {
	for {
		b.HP -= dealt(a.Damage - b.Armor)
		if b.HP <= 0 {
			return true
		}

		a.HP -= dealt(b.Damage - a.Armor)
		if a.HP <= 0 {
			return false
		}
	}
}

func dealt(x int) int {
	if x < 1 {
		return 1
	}
	return x
}

func parse() (int, int, int) {
	var hp, damage, armor int
	var line string
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line = strings.TrimSpace(scanner.Text())
	n, _ := fmt.Sscanf(line, "Hit Points: %d", &hp)
	if n != 1 {
		log.Fatal("parse failed")
	}

	scanner.Scan()
	line = strings.TrimSpace(scanner.Text())
	n, _ = fmt.Sscanf(line, "Damage: %d", &damage)
	if n != 1 {
		log.Fatal("parse failed")
	}

	scanner.Scan()
	line = strings.TrimSpace(scanner.Text())
	n, _ = fmt.Sscanf(line, "Armor: %d", &armor)
	if n != 1 {
		log.Fatal("parse failed")
	}

	return hp, damage, armor
}
