package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Replacement struct {
	From string
	To   string
}

func main() {
	input, replacements := parse()

	//for k := range replacements {
	//	sort.Slice(replacements[k], func(i, j int) bool { return len(replacements[k][i]) > len(replacements[k][j]) })
	//}
	//
	//sort.Slice(repl, func(i, j int) bool { return len(replacements[repl[i]][0]) > len(replacements[repl[j]][0]) })

	//walk(input, 0)

	rand.Seed(time.Now().UnixNano())
	var steps = 0
	same := 0

	res := input
	for {
		fmt.Println(res)

		randomReplacement := replacements[rand.Intn(len(replacements))]

		newRes := strings.Replace(res, randomReplacement.To, randomReplacement.From, 1)
		if newRes == "e" {
			fmt.Println(steps)
			return
		}

		if res == newRes {
			same++
			if same == 30 {
				//restart
				res = input
				same = 0
				steps = 0
				continue
			}
		}

		if res != newRes {
			steps++
			same = 0
			res = newRes
		}
	}

	//// set of found molecules
	//molecules := map[string]bool{"e": true}
	//
	//for numStep := 1; ; numStep++ {
	//	children := make(map[string]bool)
	//
	//	for m := range molecules {
	//		for i := 0; i < len(m); i++ {
	//			for _, a := range atoms {
	//				//if i<=len(m) {
	//				if strings.HasPrefix(m[i:], a) {
	//					for r := range replacements {
	//						if replacements[r].From == a {
	//							newMolecule := applyChange(m, i, replacements[r])
	//							if newMolecule == input {
	//								fmt.Println(numStep)
	//								return
	//							}
	//							children[newMolecule] = true
	//						}
	//					}
	//				}
	//			}
	//		}
	//	}
	//
	//	molecules = children
	//	log.Printf("step %d finished. New set of molecules: %d\n", numStep, len(molecules))
	//}
}

func applyChange(old string, i int, to string) string {
	// case "=" is the same for both return. Including here for saving one
	// subslicing op
	if i+len(to) >= len(old) {
		return old[:i] + to
	}
	return old[:i] + to + old[i+len(to):]
}

func parse() (string, []Replacement) {
	scanner := bufio.NewScanner(os.Stdin)

	replacements := make([]Replacement, 0)

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

		replacements = append(replacements, Replacement{From: from, To: to})
	}

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	return input, replacements
}

//func walk(s string, steps int) {
//	//log.Println(s)
//
//	if s == "e" {
//		fmt.Println(steps)
//		os.Exit(0)
//	}
//
//	for _, r := range repl {
//		for i := 0; i < len(s); i++ {
//			if strings.HasPrefix(s[i:], r) {
//				for _, to := range replacements[r] {
//					newS := applyChange(s, i, to)
//					walk(newS, steps+1)
//				}
//			}
//		}
//	}
//}
