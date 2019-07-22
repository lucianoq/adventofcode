package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Replacement struct {
	From string
	To   string
}

func main() {
	atoms, replacements, input := parse()

	// set of found molecules
	molecules := make(map[string]bool)

	// loop on the input string
	for i := 0; i < len(input); i++ {

		// Try to see if the substring starting in `i` is a known atom
		for _, a := range atoms {
			if strings.HasPrefix(input[i:], a) {

				// ok, we found an atom `a` starting at `i`
				// Apply all the right replacements
				for _, r := range replacements {
					if r.From == a {

						// substring before `i` is untouched
						// `r.To` is the replacement string
						// substring starting in `i+len(a)` is untouched
						newMolecule := input[:i] + r.To + input[i+len(a):]
						molecules[newMolecule] = true
					}
				}
			}
		}
	}

	fmt.Println(len(molecules))
}

func parse() (atoms []string, replacements []Replacement, input string) {

	scanner := bufio.NewScanner(os.Stdin)

	atomSet := make(map[string]bool)
	replacements = make([]Replacement, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break
		}

		var from, to string
		n, _ := fmt.Sscanf(line, "%s => %s", &from, &to)
		if n != 2 {
			log.Fatal("parsing failed")
		}

		atomSet[from] = true
		replacements = append(replacements, Replacement{From: from, To: to})
	}

	atoms = make([]string, 0, len(atomSet))
	for k := range atomSet {
		atoms = append(atoms, k)
	}

	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())

	return
}
