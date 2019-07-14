package part2

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var mem map[string]uint16

func main() {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	ops := make(map[string]bool)
	for _, l := range lines {
		if l != "" {

			// I need to remove any assignment to b
			if !strings.HasSuffix(l, " -> b") {
				ops[l] = false
			}
		}
	}

	mem = make(map[string]uint16)

	// override b
	mem["b"] = 46065

	for {

		if _, ok := mem["a"]; ok {
			fmt.Println(mem["a"])
			return
		}

		for line, ok := range ops {
			if !ok {
				success := exec(line)
				ops[line] = success
			}
		}
	}
}

func exec(line string) bool {
	op, args := extract(line)

	switch op {
	case ":=":
		return assignment(args[0], args[1])

	case "NOT":
		return not(args[0], args[1])

	case "AND":
		return and(args[0], args[1], args[2])

	case "OR":
		return or(args[0], args[1], args[2])

	case "LSHIFT":
		return lshift(args[0], args[1], args[2])

	case "RSHIFT":
		return rshift(args[0], args[1], args[2])

	}
	return false
}
