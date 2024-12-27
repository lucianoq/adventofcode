package main

type pqNode struct {
	Node
	Heat int
}

type priorityQueue []pqNode

func (n priorityQueue) Len() int {
	return len(n)
}

func (n priorityQueue) Less(i, j int) bool {
	return n[i].Heat < n[j].Heat
}

func (n priorityQueue) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *priorityQueue) Push(x any) {
	*n = append(*n, x.(pqNode))
}

func (n *priorityQueue) Pop() any {
	old := *n
	l := len(old)
	x := old[l-1]
	*n = old[0 : l-1]
	return x
}
