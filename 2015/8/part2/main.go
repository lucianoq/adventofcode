package part2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var totalLen, encodedLen int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if line != "" {
			totalLen += len(line)
			encodedLen += len(encodedString(line))
		}
	}

	fmt.Println(encodedLen - totalLen)
}

func encodedString(s string) string {
	newLine := `"`
	for _, c := range s {
		newLine += encode(c)
	}
	newLine += `"`
	return newLine
}

func encode(c rune) string {
	switch c {
	case '"':
		return `\"`
	case '\\':
		return `\\`
	default:
		return string(c)
	}
}
