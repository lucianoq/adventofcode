package main

import (
	"fmt"
	"strconv"
)

const (
	cupsSize = 1000000
	moves    = 10000000
)

func parse() *Ring {
	r := NewRing()

	for _, c := range input {
		num, _ := strconv.Atoi(string(c))
		r.Insert(num)
		r.Next()
	}

	for i := 10; i <= cupsSize; i++ {
		r.Insert(i)
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

	r.Next()
	cup1Val := r.Value()
	r.Next()
	cup2Val := r.Value()

	fmt.Println(cup1Val * cup2Val)
}
