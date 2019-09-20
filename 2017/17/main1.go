package main

import (
	"container/ring"
	"fmt"
)

const stepsForward = 303

func main() {
	r := ring.New(1)
	r.Value = 0

	for i := 1; i <= 2017; i++ {
		for j := 0; j < stepsForward; j++ {
			r = r.Next()
		}

		s := ring.New(1)
		s.Value = i
		r.Link(s)

		r = r.Next()
	}

	fmt.Println(r.Next().Value)
}
