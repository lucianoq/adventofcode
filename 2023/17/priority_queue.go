package main

import "container/heap"

type item struct {
	value    Node
	priority int
}

type PriorityQueue []*item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	it := x.(*item)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	it := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return it
}

func NewPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

// typed wrappers around heap package

func (pq *PriorityQueue) Insert(node Node, priority int) {
	heap.Push(pq, &item{
		value:    node,
		priority: priority,
	})
}

func (pq *PriorityQueue) PopMin() (Node, int) {
	it := heap.Pop(pq).(*item)
	return it.value, it.priority
}
