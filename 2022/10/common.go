package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse() <-chan string {
	ch := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch
}

func run(startCycleFunc func(cycle, x int)) {
	var (
		x       = 1
		cmdCh   = parse()
		halfAdd func()
	)

	for cycle := 1; ; cycle++ {

		startCycleFunc(cycle, x)

		if halfAdd != nil {
			halfAdd()
			halfAdd = nil
			continue
		}

		cmd, ok := <-cmdCh
		if !ok {
			break
		}
		ff := strings.Fields(cmd)

		switch ff[0] {

		case "noop":

		case "addx":
			halfAdd = func() {
				num, _ := strconv.Atoi(ff[1])
				x += num
			}
		}
	}
}
