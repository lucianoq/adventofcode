package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

const salt = "yjdafjpo"

func main() {
	var approved []int
	candidates := make(map[string][]int)

	for i := 0; ; i++ {
		hex := hash(i)

		for c := range quintuples(hex) {
			for _, candidateID := range candidates[c] {
				if i <= candidateID+1000 {
					log.Printf("Candidate [%d] approved thanks to [%d]%s", candidateID, i, hex)
					approved = append(approved, candidateID)
					log.Printf("Now we have %d approved hex", len(approved))
					if len(approved) == 64 {
						fmt.Printf("last approved was %d", candidateID)
						return
					}
				} else {
					log.Printf("Candidate [%d] discarded, because too old", candidateID)
				}
			}
			candidates[c] = []int{}
		}

		if c, ok := triples(hex); ok {
			log.Printf("[%d] Found a candidate: %s", i, hex)
			candidates[c] = append(candidates[c], i)
		}
	}
}

func hash(i int) string {
	sum := md5.Sum([]byte(salt + strconv.Itoa(i)))
	return hex.EncodeToString(sum[:])
}

func triples(s string) (string, bool) {
	for i := 0; i < len(s)-2; i++ {
		if same(s[i : i+3]) {
			return string(s[i]), true
		}
	}
	return "", false
}

func quintuples(s string) map[string]bool {
	res := make(map[string]bool)
	for i := 0; i < len(s)-4; i++ {
		if same(s[i : i+5]) {
			res[string(s[i])] = true
			i += 4
		}
	}
	return res
}

func same(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}
