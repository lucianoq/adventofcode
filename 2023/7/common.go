package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	Hand string
	Bid  int
	Type int
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func main() {
	games := parse()

	for i := 0; i < len(games); i++ {
		games[i].Type = Type(games[i].Hand)
	}

	sort.Slice(games, func(i, j int) bool {
		if games[i].Type < games[j].Type {
			return true
		}
		if games[i].Type > games[j].Type {
			return false
		}

		for c := 0; c < 5; c++ {
			card1, card2 := value[games[i].Hand[c]], value[games[j].Hand[c]]
			if card1 < card2 {
				return true
			}
			if card1 > card2 {
				return false
			}
		}
		return false
	})

	winning := 0
	for i, g := range games {
		winning += g.Bid * (i + 1)
	}

	fmt.Println(winning)
}

func parse() []Game {
	scanner := bufio.NewScanner(os.Stdin)

	var games []Game
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		hand := fields[0]
		bid, _ := strconv.Atoi(fields[1])

		games = append(games, Game{
			Hand: hand,
			Bid:  bid,
		})
	}
	return games
}

func getTypeByCount(count map[rune]int) int {
	values := sort.IntSlice{}
	for _, v := range count {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(values))

	switch values[0] {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if values[1] == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if values[1] == 2 {
			return TwoPair
		}
		return OnePair
	case 1:
		return HighCard
	}
	log.Fatal("impossible")
	return 0
}
