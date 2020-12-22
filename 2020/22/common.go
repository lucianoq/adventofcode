package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() (*Deck, *Deck) {
	scanner := bufio.NewScanner(os.Stdin)

	deck1 := &Deck{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		if strings.HasPrefix(line, "Player") {
			continue
		}

		num, _ := strconv.Atoi(line)
		deck1.Append(num)
	}

	deck2 := &Deck{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Player") {
			continue
		}

		num, _ := strconv.Atoi(line)
		deck2.Append(num)
	}

	return deck1, deck2
}

func main() {
	deck1, deck2 := parse()

	if play(deck1, deck2) {
		fmt.Println(deck1.Score())
	} else {
		fmt.Println(deck2.Score())
	}
}
