package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Value    string
	Parent   *Node
	Children []*Node
}

func main() {
	nodes := parse()

	you := nodes["YOU"]
	santa := nodes["SAN"]

	// two lists of nodes
	// from root to the target nodes
	youChain := fullChain(you)
	santaChain := fullChain(santa)

	// find the closer common ancestor
	// root in the worst case
	i, j := len(youChain)-1, len(santaChain)-1
	for youChain[i] == santaChain[j] {
		i--
		j--
	}

	// i is the length from the common ancestor to YOU
	// j is the length from the common ancestor to SAN
	fmt.Println(i + j)
}

func parse() map[string]*Node {
	tree := map[string][]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), ")")
		tree[nodes[0]] = append(tree[nodes[0]], nodes[1])
	}

	nodes := make(map[string]*Node)

	root := &Node{
		Value:    "COM",
		Parent:   nil,
		Children: make([]*Node, 0),
	}
	nodes["COM"] = root
	toDo := []*Node{root}

	var current *Node
	for len(toDo) > 0 {
		current, toDo = toDo[0], toDo[1:]

		for _, child := range tree[current.Value] {
			c := &Node{
				Value:    child,
				Parent:   current,
				Children: make([]*Node, 0),
			}
			nodes[child] = c
			toDo = append(toDo, c)
		}
	}

	return nodes
}

// array of nodes, representing the full path from
// the input node and root (both included)
func fullChain(from *Node) []*Node {
	c := make([]*Node, 0)
	for curr := from; curr != nil; curr = curr.Parent {
		c = append(c, curr)
	}
	return c
}
