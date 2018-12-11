package main

import (
	"fmt"
	"log"
)

//const (
//	players = 479
//	marbles = 71035
//)

const (
	players = 9
	marbles = 25
)

var scores []int

func main() {
	fmt.Println(Solve(players, marbles))
}

func Solve(players, marbles int) int {
	scores = make([]int, players)

	c := Circle{
		Current: 0,
		List:    []int{0},
	}

	for i := 1; i <= marbles; i++ {
		player := i % players
		log.Printf("Player %d wants to add marble %d", player, i)
		scores[player] += c.Add(i)
		//Dump(c)
	}

	max := -1
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	//spew.Dump(scores)

	return max
}

type Circle struct {
	Current int
	List    []int
}

func (c *Circle) Add(marble int) int {
	if marble%23 == 0 {
		return c.insertSpecial(marble)
	}

	l := len(c.List)
	newList := make([]int, 0)
	left := c.List[:(c.Current+2)%l]
	//fmt.Println(left)
	right := c.List[(c.Current+2)%l:]
	//fmt.Println(right)

	newList = append(newList, left...)
	newList = append(newList, marble)
	newList = append(newList, right...)
	//fmt.Println(newList)
	c.List = newList
	c.Current = (c.Current + 2) % l
	return 0
}

func (c *Circle) insertSpecial(marble int) int {
	toRemove := Abs((c.Current - 7) % len(c.List))
	toRemoveElem := c.List[toRemove]
	c.List = append(c.List[:toRemove], c.List[toRemove+1:]...)
	c.Current = toRemove % len(c.List)
	return marble + toRemoveElem
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Dump(c Circle) {
	for i, it := range c.List {
		if i == c.Current {
			fmt.Printf("(%d) ", it)
		} else {
			fmt.Printf(" %d  ", it)
		}
	}
	fmt.Println()
}
