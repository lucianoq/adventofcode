package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start, stop, digits int
}

func (r Range) InvalidIDs(chunkLen int) map[int]struct{} {
	repeat := r.digits / chunkLen

	invalids := make(map[int]struct{})

	for i := r.start / pow10(r.digits-chunkLen); i <= r.stop/(pow10(r.digits-chunkLen)+1); i++ {
		id := buildID(i, repeat)
		if id < r.start {
			continue
		}
		if id > r.stop {
			break
		}
		invalids[id] = struct{}{}
	}

	return invalids
}

func parseInput() []Range {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var list []Range
	for _, r := range strings.Split(scanner.Text(), ",") {
		limits := strings.Split(r, "-")
		start, _ := strconv.Atoi(limits[0])
		stop, _ := strconv.Atoi(limits[1])
		startDigits, stopDigits := digits(start), digits(stop)

		// split the range into two ranges
		// e.g., start->999 and 1000->stop
		// to work on single digits range later
		if startDigits == stopDigits {
			list = append(list, Range{start, stop, startDigits})
		} else {
			list = append(list, Range{start, pow10(stopDigits-1) - 1, startDigits})
			list = append(list, Range{pow10(startDigits), stop, stopDigits})
		}
	}
	return list
}

// faster than
// int(math.Pow10(exp))
func pow10(exp int) int {
	base := 10
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}
	return result
}

// faster than
// len(strconv.Itoa(i))
func digits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

// faster than
// strconv.Atoi(strings.Repeat(strconv.Itoa(chunk), repeat))
func buildID(chunk, repeat int) int {
	n := chunk
	for i := 1; i < repeat; i++ {
		n = n*pow10(digits(chunk)) + chunk
	}
	return n
}
