package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(computeSize(str))
}

func computeSize(str []byte) int {
	count := 0
	for i := 0; i < len(str); i++ {

		// ignore spaces, newline, tabs, etc.
		if unicode.IsSpace(rune(str[i])) {
			continue
		}

		if str[i] != '(' {
			count++
			continue
		}

		closeIdx := strings.IndexByte(string(str[i+1:]), ')') + i + 1
		nums := strings.Split(string(str[i+1:closeIdx]), "x")
		howMany, _ := strconv.Atoi(nums[0])
		times, _ := strconv.Atoi(nums[1])

		chunk := str[closeIdx+1 : closeIdx+1+howMany]

		count += computeSize(chunk) * times

		i = closeIdx + howMany
	}
	return count
}
