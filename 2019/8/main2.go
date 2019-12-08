package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	Width  = 25
	Height = 6
	Layer  = Width * Height

	Black       byte = '0'
	White       byte = '1'
	Transparent byte = '2'
)

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)

	finalImage := make([]byte, Layer)

	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {

			pixel := Transparent
			l := 0

			for pixel = buf[l*Layer+h*Width+w]; pixel == Transparent; pixel = buf[l*Layer+h*Width+w] {
				l++
			}

			finalImage[h*Width+w] = pixel
		}
	}

	// print
	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			switch finalImage[h*Width+w] {
			case Black:
				fmt.Print(" ")
			case White:
				fmt.Print("█")
			}
		}
		fmt.Println()
	}
}
