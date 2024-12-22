package main

import "fmt"

func main() {
	secrets := parseInput()

	sum := 0
	for _, secret := range secrets {
		for i := 0; i < 2000; i++ {
			secret = compute(secret)
		}
		sum += secret
	}

	fmt.Println(sum)
}
