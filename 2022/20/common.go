package main

import (
	"bufio"
	"os"
	"strconv"
)

type Node struct {
	Val        int
	Prev, Next *Node
}

func parse() ([]*Node, *Node) {
	index := make([]*Node, 0)
	var zero *Node
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		newNode := &Node{Val: num}
		index = append(index, newNode)
		if num == 0 {
			zero = newNode
		}
	}

	// link items each other
	for i := 0; i < len(index); i++ {
		index[i].Prev = index[(i+len(index)-1)%len(index)]
		index[i].Next = index[(i+1)%len(index)]
	}

	return index, zero
}

func groveCoordinates(curr *Node) int {
	sum := 0
	for i := 0; ; i++ {
		switch i {
		case 1000, 2000:
			sum += curr.Val
		case 3000:
			return sum + curr.Val
		}
		curr = curr.Next
	}
}

func move(curr *Node, delta int) *Node {
	for ; delta > 0; delta-- {
		curr = curr.Next
	}
	for ; delta < 0; delta++ {
		curr = curr.Prev
	}
	return curr
}

func mix(index []*Node) {
	for _, p := range index {

		// Remove p from the chain
		p.Prev.Next = p.Next
		p.Next.Prev = p.Prev

		curr := move(
			p.Prev,               // left side element in order to add p to the right.
			p.Val%(len(index)-1), // -1 because I'm rotating on the list without p
		)

		// insert p to the right of curr
		p.Prev = curr
		p.Next = curr.Next
		curr.Next.Prev = p
		curr.Next = p
	}
}
