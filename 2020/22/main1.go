package main

func play(deck1, deck2 *Deck) bool {
	for t := 0; ; t++ {
		card1 := deck1.Draw()
		card2 := deck2.Draw()

		if card1 > card2 {
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
