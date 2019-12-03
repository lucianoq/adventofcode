package main

import "fmt"

func main() {
	fmt.Println(run(sumOfSteps))
}

func sumOfSteps(_ C, step0, step1 int) int {
	return step0 + step1
}
