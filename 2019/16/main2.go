package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	list := parse()
	fmt.Println(part2(list))
}

func part2(list []int8) string {
	// repeat 10,000 times
	list = multiply10k(list)

	// apply 100 phases
	list = run100phases(list)

	// extract the 8 digits considering offset
	return extractDigitsOffset(list)
}

func multiply10k(list []int8) []int8 {
	newList := make([]int8, 0, len(list)*10000)
	for i := 0; i < 10000; i++ {
		newList = append(newList, list...)
	}
	return newList
}

func extractDigitsOffset(list []int8) string {
	offset := 0
	for i, v := range list[:7] {
		offset += int(math.Pow10(6-i)) * int(v)
	}
	log.Printf("offset = %d", offset)

	str := ""
	for _, i := range list[offset : offset+8] {
		str += strconv.Itoa(int(i))
	}
	return str
}
