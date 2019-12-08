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
)

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)

	numLayers := len(buf) / Layer
	minLayer := -1
	minZeros := 1<<63 - 1

	for l := 0; l < numLayers; l++ {

		count0 := 0

		for h := 0; h < Height; h++ {
			for w := 0; w < Width; w++ {
				idx := l*Layer + h*Width + w
				pixel := buf[idx]
				if pixel == '0' {
					count0++
				}
			}
		}

		if count0 < minZeros {
			minZeros = count0
			minLayer = l
		}
	}

	layer := buf[minLayer*Layer : (minLayer+1)*Layer]

	var num1, num2 int
	for i := 0; i < len(layer); i++ {
		switch layer[i] {
		case '1':
			num1++
		case '2':
			num2++
		}
	}

	fmt.Println(num1 * num2)
}
