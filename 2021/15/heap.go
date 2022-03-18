package main

type Node struct {
	Pos  C
	Risk int
}

type NodeHeap []Node

func (n NodeHeap) Len() int {
	return len(n)
}

func (n NodeHeap) Less(i, j int) bool {
	return n[i].Risk < n[j].Risk
}

func (n NodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *NodeHeap) Push(x any) {
	*n = append(*n, x.(Node))
}

func (n *NodeHeap) Pop() any {
	old := *n
	l := len(old)
	x := old[l-1]
	*n = old[0 : l-1]
	return x
}
