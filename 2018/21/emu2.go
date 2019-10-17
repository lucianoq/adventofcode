package main

var a, p, b, c, d, e int

func main() {

	// looking for the right `a`
	// a = ?

p0:
	d = 123

p1:
	c &= 456

p2:
	if d == 72 {
		goto p5
	} else {
		goto p1
	}

p5:
	d = 0

p6:
	e = d | 65536

p7:
	d = 10704114

p8:
	b = e & 255

p9:
	d += b

p10:
	d &= 16777215

p11:
	d *= 65899

p12:
	d &= 16777215

p13:
	if 256 > e {
		goto p28
	} else {
		goto p17
	}

p17:
	b = 0

p18:
	c = b + 1

p19:
	c *= 256

p20:
	if c > e {
		goto p26
	}

p24:
	b++

p25:
	goto p18

p26:
	e = b

p27:
	goto p8

p28:
	if d == a {
		return
	}
	goto p6

}
