package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		items := strings.Fields(line)
		switch items[0] {
		case "rect":
			ll := strings.Split(items[1], "x")
			w, _ := strconv.Atoi(ll[0])
			h, _ := strconv.Atoi(ll[1])
			rect(w, h)
		case "rotate":
			switch items[1] {
			case "row":
				y, _ := strconv.Atoi(strings.Split(items[2], "=")[1])
				by, _ := strconv.Atoi(items[4])
				rotateRow(y, by)
			case "column":
				x, _ := strconv.Atoi(strings.Split(items[2], "=")[1])
				by, _ := strconv.Atoi(items[4])
				rotateColumn(x, by)
			}
		}

		printDisplay(true)
		time.Sleep(50 * time.Millisecond)
	}
}
