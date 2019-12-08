package main

import "fmt"

func main() {
	max := 0
	for _, phase := range permutations([]int{0, 1, 2, 3, 4}) {
		res := run5(phase)
		if res > max {
			max = res
		}
	}
	fmt.Println(max)
}

func run5(phase []int) int {
	// ch0       ch1       ch2       ch3        ch4       ch5
	// -->  VM0  -->  VM1  -->  VM2  -->   VM3  -->  VM4  -->

	var chs = make([]chan int, 6)

	for i := 0; i < 6; i++ {
		chs[i] = make(chan int, 1)
	}

	for i := 0; i < 5; i++ {
		chs[i] <- phase[i]
		go NewVM("input", chs[i], chs[i+1]).Run()
	}

	chs[0] <- 0
	return <-chs[5]
}
