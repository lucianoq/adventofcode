package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	total := 0
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		total += findJolts(scanner.Text(), 2)
	}
	fmt.Println(total)
}
