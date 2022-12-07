package main

import "fmt"

const (
	DiskSize   = 70_000_000
	NeededSize = 30_000_000
)

func main() {
	root := parse()

	currentFreeSpace := DiskSize - root.Size
	minToFree := NeededSize - currentFreeSpace

	fmt.Println(FindSmallestDirToDelete(root, minToFree))
}

func FindSmallestDirToDelete(node *Node, necessary int) int {
	if !node.Dir || node.Size < necessary {
		return 1<<63 - 1
	}

	min := node.Size
	for _, c := range node.Children {
		newMin := FindSmallestDirToDelete(c, necessary)
		if newMin < min {
			min = newMin
		}
	}
	return min
}
