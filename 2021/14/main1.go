package main

import "fmt"

func main() {
	start, rules := parse()

	polymer := start
	for step := 0; step < 10; step++ {
		newPolymer := polymer[:1]
		for i := 1; i < len(polymer); i++ {
			pair := polymer[i-1 : i+1]
			newPolymer += rules[pair] + polymer[i:i+1]
		}
		polymer = newPolymer
	}

	letterCount := map[uint8]int{}
	for _, r := range polymer {
		letterCount[uint8(r)]++
	}

	min, max := minMax(letterCount)

	fmt.Println(max - min)
}
