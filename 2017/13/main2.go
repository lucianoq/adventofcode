package main

import "fmt"

func main() {
	firewall, max := parse()
	for delay := 0; ; delay++ {
		if !caught(firewall, max, delay) {
			fmt.Println(delay)
			return
		}
	}
}

func caught(firewall map[int]Scanner, max, delay int) bool {
	for i := 0; i <= max; i++ {
		if s, ok := firewall[i]; ok {
			if s.IsZero(i + delay) {
				return true
			}
		}
	}
	return false
}
