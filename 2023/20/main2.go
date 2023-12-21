package main

import (
	"fmt"
	"strings"
)

func main() {
	network := parse()

	//graphviz(network)

	// 4 high pulses from these 4 to `xn` means 1 low pulse
	// from `xn` to `rx`, our goal
	firstOccurrence := map[string]int{
		"fz": 0,
		"hn": 0,
		"xf": 0,
		"mp": 0,
	}
	count := len(firstOccurrence)

	var queue []Message

	for i := 1; ; i++ {
		queue = append(queue, Button)

		var msg Message
		for len(queue) > 0 {
			msg, queue = queue[0], queue[1:]

			// real goal
			if msg.To == "rx" && msg.Pulse == LowPulse {
				return
			}

			first, exists := firstOccurrence[msg.From]
			if exists && msg.Pulse == HighPulse && first == 0 {
				firstOccurrence[msg.From] = i
				count--
			}

			if count == 0 {
				var values []int
				for _, v := range firstOccurrence {
					values = append(values, v)
				}
				fmt.Println(lcmm(values))
				return
			}

			if network[msg.To] != nil {
				queue = append(queue, network[msg.To].Process(msg)...)
			}
		}
	}
}

func lcmm(xs []int) int {
	lcm := func(a, b int) int { return a * b / gcd(a, b) }

	result := 1
	for _, n := range xs {
		result = lcm(result, n)
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func graphviz(network Network) {
	fmt.Println("digraph {")
	fmt.Println("  broadcaster [color=blue]")
	fmt.Println("  rx [color=red]")

	for k, v := range network {
		fmt.Print(k)
		fmt.Print(" -> {")

		switch m := v.(type) {
		case *Conjunction:
			fmt.Print(strings.Join(m.Destinations, " "))
		case *FlipFlop:
			fmt.Print(strings.Join(m.Destinations, " "))
		case *Broadcast:
			fmt.Print(strings.Join(m.Destinations, " "))
		}

		fmt.Println("}")
	}

	fmt.Println("}")
}
