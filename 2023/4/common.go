package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	ID   int
	Win  map[int]struct{}
	Have []int
}

const Cards int = 202

func parse() map[int]Card {
	scanner := bufio.NewScanner(os.Stdin)
	cards := map[int]Card{}
	for scanner.Scan() {
		line := scanner.Text()

		c := Card{
			ID:   0,
			Win:  map[int]struct{}{},
			Have: []int{},
		}

		ff := strings.Split(line, ":")
		_, _ = fmt.Sscanf(ff[0], "Card %d", &c.ID)

		ff2 := strings.Split(ff[1], "|")

		for _, w := range strings.Fields(ff2[0]) {
			n, _ := strconv.Atoi(w)
			c.Win[n] = struct{}{}
		}

		for _, h := range strings.Fields(ff2[1]) {
			n, _ := strconv.Atoi(h)
			c.Have = append(c.Have, n)
		}

		cards[c.ID] = c
	}

	return cards
}
