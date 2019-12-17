package main

import (
	"fmt"
	"log"
	"time"
)

const (
	maxX = 58
	maxY = 51
)

type C struct{ x, y int }

func main() {
	input := make(chan int, 0)
	output := make(chan int, 0)

	go func() {
		vm := NewVM("input", input, output)
		vm.Code[0] = 2
		vm.Run()
		close(output)
	}()

	// Send main
	log.Println("sending main()")
	send(input, "A,A,B,C,B,C,B,C")

	// Send A
	log.Println("sending A()")
	send(input, "10,L,4,R,2,R,R,3")

	// Send B
	log.Println("sending B()")
	send(input, "10,L,4,R,2,R,R,3")

	// Send C
	log.Println("sending C()")
	send(input, "10,L,4,R,2,R,R,3")

	// continuous video feed
	log.Println("sending y")
	send(input, "y")

	pos := C{0, 0}
	for x := range output {
		if x > 127 {
			fmt.Println(x)
			return
		}
		fmt.Print(string(x))
		if x == '\n' {
			pos.y++
			pos.x = 0
		} else {
			pos.x++
		}

		if pos.y == 51 {
			time.Sleep(time.Second)
			fmt.Print("\033[2J\033[H")
			pos.y = 0
			pos.x = 0
		}
	}
}

func send(input chan<- int, s string) {
	if len(s) > 20 {
		log.Fatal("string exceeded limits")
	}

	for _, r := range s {
		log.Printf("sending %d", r)
		input <- int(r)
	}
	log.Printf("sending %d", '\n')
	input <- '\n'
}
