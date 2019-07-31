package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var dirs = "NESW"

func parse() []string {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal()
	}
	return strings.Split(strings.TrimSpace(string(buf)), ", ")
}

