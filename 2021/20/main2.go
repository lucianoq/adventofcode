package main

import "fmt"

func main() {
	image := parse()

	for i := 0; i < 50; i++ {
		image.Enhance()
	}

	fmt.Println(image.CountPixelsOn())
}
