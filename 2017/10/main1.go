package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const size = 256

func main() {
	list := createList(size)
	pos := 0
	skipSize := 0
	input := parse()

	for _, l := range input {
		list = hash(list, pos, l)
		pos += l + skipSize
		skipSize++
	}

	fmt.Println(list[0] * list[1])
}

func parse() []int {
	list := make([]int, 0)
	text, _ := ioutil.ReadAll(os.Stdin)
	items := strings.Split(string(text), ",")
	for _, i := range items {
		n, _ := strconv.Atoi(i)
		list = append(list, n)
	}
	return list
}

func createList(length int) []int {
	list := make([]int, length)
	for i := 0; i < length; i++ {
		list[i] = i
	}
	return list
}

func hash(list []int, pos int, l int) []int {
	ls := make([]int, len(list))

	j := pos
	for i := l + pos - 1; j < l+pos; i-- {
		ls[j%len(list)] = list[i%len(list)]
		j++
	}

	for i := l + pos; j < pos+len(list); i++ {
		ls[j%len(list)] = list[i%len(list)]
		j++
	}
	return ls
}
