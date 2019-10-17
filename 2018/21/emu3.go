package main

var a, p, b, c, d, e int

func main() {

	// looking for the right `a`
	// a = ?

	d = 123

	for {
		c &= 456

		if d == 72 {
			break
		}
	}

	d = 0

p6:
	for {
		e = d | 65536
		d = 10704114

	p8:
		b = e & 255
		d = (((d + b) & 16777215) * 65899) & 16777215

		if 256 > e {
			if d == a {
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
