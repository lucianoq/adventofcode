package main

import (
	"fmt"
	"strconv"
)

func main() {
	list := parse()
	list = run100phases(list)
	fmt.Println(extractFirst8(list))
}

func extractFirst8(list []int8) string {
	str := ""
	for _, i := range list[:8] {
		str += strconv.Itoa(int(i))
	}
	return str
}
