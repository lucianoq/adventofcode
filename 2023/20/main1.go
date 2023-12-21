package main

import "fmt"

func main() {
	network := parse()
	var queue []Message

	lowSent, highSent := 0, 0
	count := func(msg Message) {
		if msg.Pulse == LowPulse {
			lowSent++
		} else {
			highSent++
		}
	}

	for i := 0; i < 1000; i++ {
		queue = append(queue, Button)

		var msg Message
		for len(queue) > 0 {
			msg, queue = queue[0], queue[1:]
			count(msg)

			if network[msg.To] != nil {
				queue = append(queue, network[msg.To].Process(msg)...)
			}
		}
	}

	fmt.Println(lowSent * highSent)
}
