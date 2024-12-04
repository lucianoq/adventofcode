package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}

	rg := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	res := rg.FindAllStringSubmatch(text, -1)

	sum := 0
	for _, l := range res {
		n1, _ := strconv.Atoi(l[1])
		n2, _ := strconv.Atoi(l[2])
		sum += n1 * n2
	}

	fmt.Println(sum)
}
