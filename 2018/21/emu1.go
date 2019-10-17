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
		d = 1
	} else {
		d = 0
	}

p3:
	p += d

p4:
	p = 0

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
		b = 1
	} else {
		b = 0
	}

p14:
	p += b

p15:
	p++

p16:
	p = 27

p17:
	b = 0

p18:
	c = b + 1

p19:
	c *= 256

p20:
	if c > e {
		c = 1
	} else {
		c = 0
	}

p21:
	p += c

p22:
	p++

p23:
	p = 25

p24:
	b++

p25:
	p = 17

p26:
	e = b

p27:
	p = 7

p28:
	if d == a {
		b = 1
	} else {
		b = 0
	}

p29:
	p += b

p30:
	p = 5
}
