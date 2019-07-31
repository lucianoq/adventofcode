package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var MFCSAM = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	db := parse()

	for i := 1; i < 501; i++ {
		if compatible(db[i]) {
			fmt.Println(i)
		}
	}
}

func compatible(test map[string]int) bool {
	for k, v := range test {
		switch k {
		case "cats", "trees":
			if v <= MFCSAM[k] {
				return false
			}
		case "pomeranians", "goldfish":
			if v >= MFCSAM[k] {
				return false
			}
		default:
			if v != MFCSAM[k] {
				return false
			}
		}
	}
	return true
}

func parse() [501]map[string]int {
	db := [501]map[string]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		r, err := regexp.Compile("^Sue (\\d+): (.*)$")
		if err != nil {
			log.Fatal(err)
		}

		arr := r.FindStringSubmatch(line)
		id, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}
		list := arr[2]

		db[id] = make(map[string]int)

		kvs := strings.Split(list, ", ")
		for _, kv := range kvs {
			xs := strings.Split(kv, ": ")
			n, err := strconv.Atoi(xs[1])
			if err != nil {
				log.Fatal(err)
			}
			db[id][xs[0]] = n
		}
	}

	return db
}
