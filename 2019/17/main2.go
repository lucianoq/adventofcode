package main

import "fmt"

func main() {
	input := make(chan int, 0)
	output := make(chan int, 0)

	done1 := make(chan struct{})
	go func() {
		vm := NewVM(input, output)
		vm.Code[0] = 2
		vm.Run()
		close(output)
		done1 <- struct{}{}
	}()

	done2 := make(chan struct{})
	go func() {
		defer func() {
			done2 <- struct{}{}
		}()
		for x := range output {
			if x > 127 {
				fmt.Println(x)
				return
			}
		}
	}()

	// Send main
	send(input, "A,B,A,C,A,A,C,B,C,B")

	// Send A
	send(input, "L,12,L,8,R,12")

	// Send B
	send(input, "L,10,L,8,L,12,R,12")

	// Send C
	send(input, "R,12,L,8,L,10")

	// continuous video feed
	send(input, "n")

	<-done1
	<-done2
}

func send(input chan<- int, s string) {
	for _, r := range s {
		input <- int(r)
	}
	input <- '\n'
}
