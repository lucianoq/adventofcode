package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	sum := 0
	for _, seq := range strings.Split(scanner.Text(), ",") {
		sum += hash(seq)
	}

	fmt.Println(sum)
}
