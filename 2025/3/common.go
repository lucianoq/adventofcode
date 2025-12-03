package main

func findJolts(bank string, digits int) int {
	output := 0
	idx := -1
	for i := digits; i > 0; i-- {
		var jolts int
		jolts, idx = findMax(bank, idx+1, len(bank)-i+1)
		output = output*10 + jolts
	}
	return output
}

func findMax(list string, start, end int) (int, int) {
	var maxJolts uint8
	var idx int
	for i := start; i < end; i++ {
		if list[i] > maxJolts {
			maxJolts = list[i]
			idx = i
		}
	}
	return int(maxJolts - '0'), idx
}
