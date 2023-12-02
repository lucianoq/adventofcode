package main

import "fmt"

func main() {
	db := parse()

	var sum int
	for _, exs := range db {
		m := max(exs)
		sum += m.Blue * m.Red * m.Green
	}

	fmt.Println(sum)
}

func max(exs []Extraction) Extraction {
	var maximum Extraction
	for _, ex := range exs {
		if ex.Red > maximum.Red {
			maximum.Red = ex.Red
		}
		if ex.Green > maximum.Green {
			maximum.Green = ex.Green
		}
		if ex.Blue > maximum.Blue {
			maximum.Blue = ex.Blue
		}
	}
	return maximum
}
