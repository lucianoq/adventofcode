package main

import "fmt"

func main() {
	firewall, max := parse()
	fmt.Println(severity(firewall, max))
}

func severity(firewall map[int]Scanner, max int) int {
	var sum = 0
	for i := 0; i <= max; i++ {
		if s, ok := firewall[i]; ok {
			if s.IsZero(i) {
				sum += s.Range() * i
			}
		}
	}
	return sum
}
