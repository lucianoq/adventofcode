package main

import "fmt"

func main() {
	m := parse()

	sum := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {

			c := C{i, j}
			ch := m[c]

			if isDigit(ch) {
				valid := false
				num := 0
				for isDigit(ch) {
					num = 10*num + int(ch-'0')
					valid = valid || isValid(m, c)
					j++
					c = C{i, j}
					ch = m[c]
				}

				if valid {
					sum += num
				}
			}
		}
	}
	fmt.Println(sum)
}

func isValid(m map[C]byte, c C) bool {
	delta := []int{-1, 0, 1}
	for _, di := range delta {
		for _, dj := range delta {
			if di|dj != 0 {
				ch, exists := m[C{c.i + di, c.j + dj}]
				if exists && !isDigit(ch) {
					return true
				}
			}
		}
	}
	return false
}
