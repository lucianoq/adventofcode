package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	ls := parse()
	ls = append(ls, 0) // adding jolts source
	sort.Ints(ls)

	count1 := 0
	count3 := 0

	for i := 0; i < len(ls)-1; i++ {
		switch ls[i+1] - ls[i] {
		case 1:
			count1++
		case 2:
		case 3:
			count3++
		default:
			log.Fatal("more than 3")
		}
	}

	// because of the built-in adapter that is always 3 more.
	count3++

	fmt.Println(count1 * count3)
}

func parse() []int {
	ls := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		ls = append(ls, n)
	}
	return ls
}
