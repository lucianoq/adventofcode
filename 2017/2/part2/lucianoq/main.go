package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var err error
	sum := 0
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) != "" {
			list := strings.Fields(scanner.Text())
			l := len(list)
			n := make([]int, l)
			for i := 0; i < l; i++ {
				n[i], err = strconv.Atoi(list[i])
				if err != nil {
					log.Fatal(err)
				}
			}

			for i := 0; i < l; i++ {
				for j := 0; j < l; j++ {
					if i == j {
						continue
					}
					if n[i]%n[j] == 0 {
						sum += n[i] / n[j]
					}
				}
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
