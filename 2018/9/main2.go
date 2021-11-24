package main

import "fmt"

var scores []int

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func (n *Node) InsertOnRight(val int) *Node {
	newNode := &Node{Val: val}
	newNode.Prev = n
	newNode.Next = n.Next
	n.Next.Prev = newNode
	n.Next = newNode
	return newNode
}

func (n *Node) Delete() (*Node, int) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	return n.Next, n.Val
}

func (n *Node) GoTo(rotation int) (ret *Node) {
	ret = n
	if rotation >= 0 {
		for i := 0; i < rotation; i++ {
			ret = ret.Next
		}
		return
	}

	for i := 0; i < -rotation; i++ {
		ret = ret.Prev
	}
	return ret
}

func main() {
	fmt.Println(Solve(479, 71035*100))
}

func Solve(players, marbles int) int {
	scores = make([]int, players)
	var score int

	current := &Node{Val: 0}
	current.Prev = current
	current.Next = current

	for m := 1; m <= marbles+1; m++ {
		player := m % players

		if m%23 == 0 {
			current = current.GoTo(-7)
			current, score = current.Delete()
			scores[player] += m + score
		} else {
			current = current.GoTo(+1).InsertOnRight(m)
		}
	}

	max := -1
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	return max
}
