package main

import (
	"fmt"
	"strconv"
)

const (
	Oxygen = iota
	CO2
)

func main() {
	list := parse()
	oxygen := find(Oxygen, list)
	co2 := find(CO2, list)
	fmt.Println(oxygen * co2)
}

func find(gas int, list []string) uint64 {
	for bit := 0; ; bit++ {
		freq0, freq1 := findFrequencies(list, bit)
		keep := whatToKeep(gas, freq0, freq1)
		list = filterOut(list, bit, keep)

		if len(list) == 1 {
			val, _ := strconv.ParseUint(list[0], 2, 64)
			return val
		}
	}
}

func whatToKeep(gas int, freq0, freq1 int) uint8 {
	switch gas {
	case Oxygen:
		if freq1 >= freq0 {
			return '1'
		}
		return '0'

	case CO2:
		if freq1 >= freq0 {
			return '0'
		}
		return '1'
	}
	return 0
}

func filterOut(list []string, bit int, keep uint8) []string {
	newList := []string{}
	for _, item := range list {
		if item[bit] == keep {
			newList = append(newList, item)
		}
	}
	return newList
}

func findFrequencies(list []string, bit int) (int, int) {
	freq0, freq1 := 0, 0
	for _, item := range list {
		switch item[bit] {
		case '0':
			freq0++
		case '1':
			freq1++
		}
	}

	return freq0, freq1
}
