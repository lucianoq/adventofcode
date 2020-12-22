package main

import (
	"strconv"
	"strings"
)

func play(deck1, deck2 *Deck) bool {
	prevHands := map[string]struct{}{}

	for t := 1; ; t++ {
		st := status(deck1, deck2)
		if _, ok := prevHands[st]; ok {
			return true
		}
		prevHands[st] = struct{}{}

		card1 := deck1.Draw()
		card2 := deck2.Draw()

		var p1won bool

		if deck1.Len() >= card1 && deck2.Len() >= card2 {
			// Start a subgame
			p1won = play(
				deck1.CopyN(card1),
				deck2.CopyN(card2),
			)
		} else {
			// game continues normally
			p1won = card1 > card2
		}

		if p1won {
			deck1.Append(card1)
			deck1.Append(card2)
		} else {
			deck2.Append(card2)
			deck2.Append(card1)
		}

		if deck1.Len() == 0 {
			return false
		}

		if deck2.Len() == 0 {
			return true
		}
	}
}

func status(deck1, deck2 *Deck) string {
	sb := strings.Builder{}
	sb.Grow((deck1.Len()+deck2.Len())*3 + 10)
	for _, x := range deck1.ToList() {
		sb.WriteString(strconv.Itoa(x) + ",")
	}
	sb.WriteString("/")
	for _, x := range deck2.ToList() {
		sb.WriteString(strconv.Itoa(x) + ",")
	}
	return sb.String()
}
