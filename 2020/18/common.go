package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isNum = regexp.MustCompile("^-?[0-9]+$")

type Stack []string

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Top() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() string {
	top, stack := (*s)[len(*s)-1], (*s)[:len(*s)-1]
	*s = stack
	return top
}

func (s *Stack) Push(x string) {
	*s = append(*s, x)
}

type Queue []string

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(x string) {
	*q = append(*q, x)
}

func (q *Queue) Dequeue() string {
	head, tail := (*q)[0], (*q)[1:]
	*q = tail
	return head
}

func solveReversePolish(q Queue) int {
	stack := Stack{}

	for !q.Empty() {
		item := q.Dequeue()

		switch item {
		case "+":
			t1, _ := strconv.Atoi(stack.Pop())
			t2, _ := strconv.Atoi(stack.Pop())

			stack.Push(strconv.Itoa(t1 + t2))
		case "*":
			t1, _ := strconv.Atoi(stack.Pop())
			t2, _ := strconv.Atoi(stack.Pop())

			stack.Push(strconv.Itoa(t1 * t2))
		default:
			stack.Push(item)
		}
	}

	res, _ := strconv.Atoi(stack.Pop())
	return res
}

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// adding spaces around parenthesis to tokenize them
		line = strings.ReplaceAll(line, "(", " ( ")
		line = strings.ReplaceAll(line, ")", " ) ")

		fields := strings.Fields(line)

		sum += solveReversePolish(ShuntingYard(fields))
	}

	fmt.Println(sum)
}
