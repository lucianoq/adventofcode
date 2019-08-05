package main

import "fmt"

func main() {
	freq := parse()
	str := ""
	for _, m := range freq {
		minC, minF := ' ', 1<<63-1
		for k, v := range m {
			if v < minF {
				minC, minF = k, v
			}
		}
		str += string(minC)
	}
	fmt.Println(str)
}
