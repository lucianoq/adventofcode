package main

var value = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func Type(hand string) int {
	count := map[rune]int{}

	for _, c := range hand {
		count[c]++
	}

	return getTypeByCount(count)
}
