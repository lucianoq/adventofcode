package main

import (
	"fmt"
)

var scores []int

func main() {
	fmt.Println(Solve(479, 71035))
}

func Solve(players, marbles int) int {
	scores = make([]int, players)

	c := Circle{
		Current: 0,
		List:    []int{0},
	}

	for i := 1; i <= marbles+1; i++ {
		player := i % players
		scores[player] += c.Add(i)
	}

	max := -1
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	return max
}

type Circle struct {
	Current int
	List    []int
}

func (c *Circle) Add(marble int) int {
	if marble%23 == 0 {
		toRemove := (c.Current - 7 + len(c.List)) % len(c.List)
		toRemoveElem := c.List[toRemove]
		c.List = append(c.List[:toRemove], c.List[toRemove+1:]...)
		c.Current = toRemove % len(c.List)
		return marble + toRemoveElem
	}

	l := len(c.List)
	newList := make([]int, 0)
	left := c.List[:(c.Current+2)%l]
	right := c.List[(c.Current+2)%l:]

	newList = append(newList, left...)
	newList = append(newList, marble)
	newList = append(newList, right...)
	c.List = newList
	c.Current = (c.Current + 2) % l
	return 0
}
