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

		if supportTLS(line) {
			count++
		}
	}

	fmt.Println(count)
}

func supportTLS(ip string) bool {
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

	for _, i := range in {
		if containsABBA(i) {
			return false
		}
	}

	for _, o := range out {
		if containsABBA(o) {
			return true
		}
	}

	return false
}

func containsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i] != s[i+1] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}
