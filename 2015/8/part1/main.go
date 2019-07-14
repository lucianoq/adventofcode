package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var reprLen, strLen int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if line != "" {
			if line[0] == '"' && line[len(line)-1] == '"' {
				line = line[1 : len(line)-1]
				reprLen += 2
			}

			var num int
			for len(line) > 0 {
				line, num = getToken(line)
				reprLen += num
				strLen++
			}
		}
	}

	fmt.Println(reprLen - strLen)
}

func getToken(s string) (string, int) {
	// try single
	token := s[0:1]
	if token != "\\" {
		return s[1:], 1
	}

	//try couple
	token = s[0:2]
	if token == `\\` || token == `\"` {
		return s[2:], 2
	}

	//try 4
	token = s[0:2]
	if token == `\x` {
		return s[4:], 4
	}

	log.Fatal("should not be here")
	return "", 0
}
