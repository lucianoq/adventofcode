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
	res := decompress(str)
	fmt.Println(len(res))
}

func decompress(str []byte) []byte {
	res := make([]byte, 0, len(str)*10)

	for i := 0; i < len(str); i++ {

		// ignore spaces, newline, tabs, etc.
		if unicode.IsSpace(rune(str[i])) {
			continue
		}

		if str[i] != '(' {
			res = append(res, str[i])
			continue
		}

		closeIdx := strings.IndexByte(string(str[i+1:]), ')') + i + 1
		nums := strings.Split(string(str[i+1:closeIdx]), "x")
		howMany, _ := strconv.Atoi(nums[0])
		times, _ := strconv.Atoi(nums[1])

		chunk := str[closeIdx+1 : closeIdx+1+howMany]
		for j := 0; j < times; j++ {
			res = append(res, chunk...)
		}

		i = closeIdx + howMany
	}
	return res
}
