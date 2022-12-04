package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		var x1, x2, y1, y2 int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d\n", &x1, &x2, &y1, &y2)

		if (y1-x1)*(x2-y2) >= 0 {
			count++
		}
	}
	fmt.Println(count)
}
