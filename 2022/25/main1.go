package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		sum += toDecimal(scanner.Text())
	}

	fmt.Println(toSNAFU(sum))
}

func toDecimal(snafu string) int {
	num := 0
	for pos, char := range snafu {
		var digit int
		switch char {
		case '0', '1', '2':
			digit = int(char - '0')
		case '-':
			digit = -1
		case '=':
			digit = -2
		}
		num += digit * pow(5, len(snafu)-1-pos)
	}
	return num
}

func toSNAFU(n int) string {
	const Digits = "012=-0"

	list := append([]string{"0"}, strings.Split(strconv.FormatInt(int64(n), 5), "")...)

	for i := len(list) - 1; i > 0; i-- {
		switch list[i] {
		case "0", "1", "2":
			continue
		case "3", "4", "5":
			list[i] = string(Digits[list[i][0]-'0'])
			list[i-1] = string(list[i-1][0] + 1)
		}
	}
	return strings.TrimLeft(strings.Join(list, ""), "0")
}

func pow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
