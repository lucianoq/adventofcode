package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func parse() []string {
	buf, _ := ioutil.ReadAll(os.Stdin)
	s := strings.TrimSpace(string(buf))
	return strings.Split(s, ",")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
