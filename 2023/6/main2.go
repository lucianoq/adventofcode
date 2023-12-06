package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	raceTime, _ := strconv.Atoi(strings.Join(strings.Fields(scanner.Text())[1:], ""))
	scanner.Scan()
	raceDist, _ := strconv.Atoi(strings.Join(strings.Fields(scanner.Text())[1:], ""))

	count := 0
	for t := 0; t <= raceTime; t++ {
		if (raceTime-t)*t > raceDist {
			count++
		}
	}

	fmt.Println(count)
}
