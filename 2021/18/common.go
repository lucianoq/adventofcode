package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Depth  int
	Value  int
}

func Add(a, b *Node) *Node {
	node := &Node{
		Left:  a,
		Right: b,
	}

	node.Left.Parent = node
	node.Right.Parent = node
	node.Left.deepen()
	node.Right.deepen()

	node.Reduce()
	return node
}

// Reduce
// To reduce a snailfish number, you must repeatedly do the first
// action in this list that applies to the snailfish number:
//
// If any pair is nested inside four pairs, the leftmost such pair explodes.
// If any regular number is 10 or greater, the leftmost such regular number splits.
func (n *Node) Reduce() {
	for {
		if exploded := n.Explode(); exploded {
			continue
		}
		if split := n.Split(); split {
			continue
		}
		return
	}
}

// Magnitude
// The magnitude of a pair is 3 times the magnitude
// of its left element plus 2 times the magnitude of
// its right element. The magnitude of a regular number
// is just that number.
func (n *Node) Magnitude() int {
	if n.Leaf() {
		return n.Value
	}
	return 3*n.Left.Magnitude() + 2*n.Right.Magnitude()
}

func (n *Node) String() string {
	if n.Leaf() {
		return fmt.Sprintf("%d", n.Value)
	}
	return fmt.Sprintf("[%s,%s]", n.Left, n.Right)
}

func (n *Node) Leaf() bool {
	return n.Left == nil && n.Right == nil
}

// Explode
// To explode a pair, the pair's left value is added to the first
// regular number to the left of the exploding pair (if any), and
// the pair's right value is added to the first regular number to
// the right of the exploding pair (if any). Exploding pairs will
// always consist of two regular numbers. Then, the entire exploding
// pair is replaced with the regular number 0.
func (n *Node) Explode() bool {
	node := n.firstNodeToExplode()
	if node == nil {
		return false
	}

	left := node.findClosestLeft()
	if left != nil {
		left.Value += node.Left.Value
	}

	right := node.findClosestRight()
	if right != nil {
		right.Value += node.Right.Value
	}

	*node = Node{
		Parent: node.Parent,
		Left:   nil,
		Right:  nil,
		Depth:  node.Depth,
		Value:  0,
	}

	return true
}

// Split
// To split a regular number, replace it with a pair;
// the left element of the pair should be the regular
// number divided by two and rounded down, while the
// right element of the pair should be the regular
// number divided by two and rounded up. For example,
// 10 becomes [5,5], 11 becomes [5,6], 12 becomes [6,6],
// and so on.
func (n *Node) Split() bool {
	node := n.firstNodeToSplit()
	if node == nil {
		return false
	}

	leftValue := node.Value / 2
	rightValue := (node.Value + 1) / 2

	*node = Node{
		Parent: node.Parent,
		Depth:  node.Depth,
		Value:  0,
	}

	node.Left = &Node{
		Parent: node,
		Depth:  node.Depth + 1,
		Value:  leftValue,
	}

	node.Right = &Node{
		Parent: node,
		Depth:  node.Depth + 1,
		Value:  rightValue,
	}

	return true
}

func (n *Node) deepen() {
	n.Depth++
	if !n.Leaf() {
		n.Left.deepen()
		n.Right.deepen()
	}
}

func (n *Node) copy() *Node {
	return parseTree(n.String(), nil, 0)
}

func (n *Node) firstNodeToExplode() *Node {
	if n.Depth == 4 && !n.Leaf() {
		return n
	}
	if n.Left != nil {
		if exp := n.Left.firstNodeToExplode(); exp != nil {
			return exp
		}
	}
	if n.Right != nil {
		if exp := n.Right.firstNodeToExplode(); exp != nil {
			return exp
		}
	}
	return nil
}

func (n *Node) firstNodeToSplit() *Node {
	if n.Value > 9 {
		return n
	}
	if n.Left != nil {
		if spl := n.Left.firstNodeToSplit(); spl != nil {
			return spl
		}
	}
	if n.Right != nil {
		if spl := n.Right.firstNodeToSplit(); spl != nil {
			return spl
		}
	}
	return nil
}

func (n *Node) findClosestLeft() *Node {
	if n.Parent == nil {
		return nil
	}

	// if n is a right child,
	// then we need to find the right most descendant
	// of its left brother
	if n == n.Parent.Right {
		// go down on the right
		curr := n.Parent.Left
		for !curr.Leaf() {
			curr = curr.Right
		}
		return curr
	}

	// if n is a left child,
	// then its closest left is the same
	// of its parent
	if n == n.Parent.Left {
		return n.Parent.findClosestLeft()
	}

	panic("bug")
}

func (n *Node) findClosestRight() *Node {
	if n.Parent == nil {
		return nil
	}

	// if n is a left child,
	// then we need to find the left most descendant
	// of its right brother
	if n == n.Parent.Left {
		// go down on the left
		curr := n.Parent.Right
		for !curr.Leaf() {
			curr = curr.Left
		}
		return curr
	}

	// if n is a right child,
	// then its closest right is the same
	// of its parent
	if n == n.Parent.Right {
		return n.Parent.findClosestRight()
	}

	panic("bug")
}

func parse() []*Node {
	numbers := []*Node{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		numbers = append(numbers, parseTree(line, nil, 0))
	}
	return numbers
}

func parseTree(s string, parent *Node, depth int) *Node {
	// if it's a number, return a leaf node
	if n, err := strconv.Atoi(s); err == nil {
		return &Node{
			Parent: parent,
			Depth:  depth,
			Value:  n,
		}
	}

	// if left side is not a nested node
	if s[1] != '[' {
		i := strings.Index(s, ",")
		node := &Node{
			Parent: parent,
			Depth:  depth,
		}
		node.Left = parseTree(s[1:i], node, depth+1)
		node.Right = parseTree(s[i+1:len(s)-1], node, depth+1)
		return node
	}

	d := 0
	for i := 1; i < len(s)-1; i++ {
		switch s[i] {
		case '[':
			d++
		case ']':
			d--
		}

		// found the correct depth comma at position `i`
		if d == 0 {
			node := &Node{
				Parent: parent,
				Depth:  depth,
			}

			node.Left = parseTree(s[1:i+1], node, depth+1)
			node.Right = parseTree(s[i+2:len(s)-1], node, depth+1)
			return node
		}
	}

	panic("bad input")
}
