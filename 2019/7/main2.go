package main

import (
	"fmt"
	"sync"
)

func main() {
	max := 0
	for _, phase := range permutations([]int{5, 6, 7, 8, 9}) {
		res := runLoop(phase)
		if res > max {
			max = res
		}
	}
	fmt.Println(max)
}

func runLoop(phase []int) int {
	// ch0       ch1       ch2       ch3        ch4       ch0 (loop)
	// -->  VM0  -->  VM1  -->  VM2  -->   VM3  -->  VM4  -->

	var chs = make([]chan int, 6)

	for i := 0; i < 6; i++ {
		chs[i] = make(chan int, 1)
	}

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		chs[i] <- phase[i]
		go func(i int) {
			NewVM("input", chs[i], chs[(i+1)%5]).Run()
			wg.Done()
		}(i)
	}

	chs[0] <- 0
	wg.Wait()
	return <-chs[0]
}
