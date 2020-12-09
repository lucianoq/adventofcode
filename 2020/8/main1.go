package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Op struct {
	Code string
	Val  int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := []Op{}
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		val, _ := strconv.Atoi(arr[1])

		cmds = append(cmds, Op{
			Code: arr[0],
			Val:  val,
		})
	}

	var (
		accumulator int
		cursor      int
		visited     = map[int]struct{}{}
	)

	for {

		if _, ok := visited[cursor]; ok {
			fmt.Println(accumulator)
			return
		}

		visited[cursor] = struct{}{}
		curr := cmds[cursor]

		log.Printf("executing %s %d", curr.Code, curr.Val)

		switch curr.Code {
		case "acc":

			accumulator += curr.Val
			cursor += 1
		case "jmp":
			cursor += curr.Val
		case "nop":
			cursor += 1
		}
	}
}
