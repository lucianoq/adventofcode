package main

import (
	"fmt"
	"strconv"
)

const (
	cupsSize = 9
	moves    = 100
)

func parse() *Ring {
	r := NewRing()

	for _, c := range input {
		num, _ := strconv.Atoi(string(c))
		r.Insert(num)
		r.Next()
	}

	num, _ := strconv.Atoi(string(input[0]))
	r.MoveTo(num)

	return r
}

func main() {
	r := parse()

	for m := 1; m <= moves; m++ {
		move(r)
	}

	r.MoveTo(1)

	s := ""
	r.Do(func(x int) {
		s += strconv.Itoa(x)
	})

	fmt.Println(s[1:]) // remove 1
}
