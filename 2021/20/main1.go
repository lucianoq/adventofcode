package main

import "fmt"

func main() {
	image := parse()
	image.Enhance()
	image.Enhance()
	fmt.Println(image.CountPixelsOn())
}
