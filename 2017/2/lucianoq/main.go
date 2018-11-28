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
	for scanner.Scan() {
		if scanner.Text() != "" {
			list := strings.Split(scanner.Text(), "")
			l := len(list)

			sum := 0
			for i:=0; i<l; i++ {
				if list[i] == list[(i+l/2)%l] {
					n, err := strconv.Atoi(list[i])
					if err != nil {
						log.Fatal(err)
					}
					sum += n
				}
			}

			fmt.Println(sum)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
