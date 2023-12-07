package main

var value = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1, //Joker
}

func Type(hand string) int {
	count := map[rune]int{}

	for _, c := range hand {
		count[c]++
	}

	assignJokers(count)

	return getTypeByCount(count)
}

func assignJokers(count map[rune]int) {
	for count['J'] > 0 {
		vMaxNotJoker := -1
		var card rune
		for k, v := range count {
			if k == 'J' {
				continue
			}
			if v > vMaxNotJoker {
				vMaxNotJoker = v
				card = k
			}
		}
		count['J']--
		count[card]++
	}
}
