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
			list = append(list, list[0])

			sum := 0
			for i := 0; i < len(list)-1; i++ {
				if list[i] == list[i+1] {
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
