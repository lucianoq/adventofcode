package main

import "fmt"

const LimitSize = 100_000

func main() {
	root := parse()

	fmt.Println(SumSizeSmallFolder(root))
}

func SumSizeSmallFolder(node *Node) int {
	if !node.Dir {
		return 0
	}

	sum := 0
	if node.Size < LimitSize {
		sum += node.Size
	}
	for _, c := range node.Children {
		sum += SumSizeSmallFolder(c)
	}
	return sum
}
