package main

import "fmt"

func main() {
	list := parse()

	sum := 0
	for _, row := range list {
		p := buildPascal(row)

		p[len(p)-1] = append([]int{0}, p[len(p)-1]...)

		for i := len(p) - 2; i >= 0; i-- {
			p[i] = append([]int{p[i][0] - p[i+1][0]}, p[i]...)
		}

		sum += p[0][0]
	}

	fmt.Println(sum)
}
