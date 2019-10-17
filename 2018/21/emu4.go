package main

import "fmt"

var input, p, b, c, d, e int

func main() {
	discovered := make(map[int]struct{})

	// looking for the right `a`
	// a = ?

	d = 0
p6:
	for {
		e = d | 65536
		d = 10704114

	p8:
		b = e & 255
		d = (((d + b) & 16777215) * 65899) & 16777215

		if 256 > e {
			if _, ok := discovered[d]; ok {
				return
			}
			discovered[d] = struct{}{}
			fmt.Println(d)
			if d == input {
				return
			} else {
				continue p6
			}
		}

		b = 0

		for {
			c = (b + 1) * 256

			if c > e {
				e = b
				goto p8
			}

			b++
		}
	}
}
