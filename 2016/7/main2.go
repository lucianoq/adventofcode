package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if supportSSL(line) {
			count++
		}
	}

	fmt.Println(count)
}

func supportSSL(ip string) bool {
	in, out := make([]string, 0), make([]string, 0)
	for len(ip) > 0 {
		i := strings.IndexByte(ip, '[')

		if i == -1 {
			out = append(out, ip)
			break
		}
		out = append(out, ip[:i])

		ip = ip[i+1:]

		i = strings.IndexByte(ip, ']')
		in = append(in, ip[:i])

		ip = ip[i+1:]
	}

	var abas []string

	for _, o := range out {
		abas = append(abas, findABAs(o)...)
	}

	if len(abas) == 0 {
		return false
	}

	for _, i := range in {
		for _, aba := range abas {
			if containsBAB(i, aba) {
				return true
			}
		}
	}

	return false
}

func findABAs(s string) []string {
	list := make([]string, 0)
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			list = append(list, s[i:i+3])
		}
	}
	return list
}

func containsBAB(s, aba string) bool {
	bab := string(append([]byte{}, aba[1], aba[0], aba[1]))
	if strings.Contains(s, bab) {
		return true
	}
	return false
}
