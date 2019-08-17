package main

import "fmt"

const input = "10001001100000001"

func run(diskSize int) {
	x := make([]bool, len(input), diskSize*2)
	for i, c := range input {
		x[i] = c == '1'
	}

	x = fillMemory(x, diskSize)
	x = checkSum(x)

	printStr(x)
}

func printStr(x []bool) {
	for _, b := range x {
		if b {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
}

func fillMemory(x []bool, diskSize int) []bool {
	for len(x) < diskSize {
		b := make([]bool, len(x))
		for i := 0; i < len(b); i++ {
			b[i] = !x[len(x)-1-i]
		}
		x = append(x, false)
		x = append(x, b...)
	}
	return x[:diskSize]
}

func checkSum(x []bool) []bool {
	for len(x)%2 == 0 {
		for i := 0; i < len(x)/2; i++ {
			x[i] = x[2*i] == x[2*i+1]
		}
		x = x[:len(x)/2]
	}
	return x
}
