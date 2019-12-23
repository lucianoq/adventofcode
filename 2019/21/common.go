package main

import "fmt"

func sendString(ch chan<- int, s string) {
	for _, r := range s {
		ch <- int(r)
	}
	ch <- '\n'
}

func reader(input <-chan int, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()

	// var newLine bool
	for c := range input {
		if c > 127 {
			fmt.Println(c)
			return
		}

		// fmt.Print(string(c))
		// if newLine && c == '\n' {
		// 	time.Sleep(1000 * time.Millisecond)
		// 	fmt.Print("\033[2J\033[H")
		// 	newLine = false
		// 	continue
		// }
		// newLine = c == '\n'
	}
}
