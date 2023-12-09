package main

import "fmt"

func main() {
	list := parse()

	sum := 0
	for _, row := range list {
		p := buildPascal(row)

		p[len(p)-1] = append(p[len(p)-1], 0)

		for i := len(p) - 2; i >= 0; i-- {
			p[i] = append(p[i], p[i][len(p[i])-1]+p[i+1][len(p[i+1])-1])
		}

		sum += p[0][len(p[0])-1]
	}

	fmt.Println(sum)
}
