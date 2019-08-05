package main

import "fmt"

func main() {
	freq := parse()
	str := ""
	for _, m := range freq {
		maxC, maxF := ' ', 0
		for k, v := range m {
			if v > maxF {
				maxC, maxF = k, v
			}
		}
		str += string(maxC)
	}
	fmt.Println(str)
}
