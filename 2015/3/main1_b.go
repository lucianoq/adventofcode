package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)
	input := strings.TrimSpace(string(buf))

	space := make(map[int]bool)
	santa := 0

	space[santa] = true
	count := 1

	for _, c := range input {
		switch c {
		case '^':
			santa += 1
		case 'v':
			santa -= 1
		case '<':
			santa -= len(buf)
		case '>':
			santa += len(buf)
		}

		if space[santa] == false {
			count++
		}
		space[santa] = true
	}

	fmt.Println(count)
}
