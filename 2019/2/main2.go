package main

import "fmt"

const expectedOutput = 19690720

func main() {
	code := parse()

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {

			copiedCode := make([]int, len(code))
			copy(copiedCode, code)

			copiedCode[1], copiedCode[2] = noun, verb

			if run(copiedCode) == expectedOutput {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}
