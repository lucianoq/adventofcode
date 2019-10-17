package main

var r0, r1, r2, r3, r4, r5 int

func main() {
p0:
	// seti 123 0 4
	r4 = 123

p1:
	// bani 4 456 4
	c &= 456

p2:
	// eqri 4 72 4
	if r4 == 72 {
		r4 = 1
	} else {
		r4 = 0
	}

p3:
	// addr 4 1 1
	r1 += r4

p4:
	// seti 0 0 1
	r1 = 0

p5:
	// seti 0 0 4
	r4 = 0

p6:
	// bori 4 65536 5
	r5 = r4 | 65536

p7:
	// seti 10704114 0 4
	r4 = 10704114

p8:
	// bani 5 255 2
	r2 = r5 & 255

p9:
	// addr 4 2 4
	r4 += r2

p10:
	// bani 4 16777215 4
	r4 &= 16777215

p11:
	// muli 4 65899 4
	r4 *= 65899

p12:
	// bani 4 16777215 4
	r4 &= 16777215

p13:
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}

p14:
	// addr 2 1 1
	r1 += r2

p15:
	// addi 1 1 1
	r1++

p16:
	// seti 27 2 1
	r1 = 27

p17:
	// seti 0 4 2
	r2 = 0

p18:
	// addi 2 1 3
	r3 = r2 + 1

p19:
	// muli 3 256 3
	r3 *= 256

p20:
	// gtrr 3 5 3
	if r3 > r5 {
		r3 = 1
	} else {
		r3 = 0
	}

p21:
	// addr 3 1 1
	r1 += r3

p22:
	// addi 1 1 1
	r1++

p23:
	// seti 25 5 1
	r1 = 25

p24:
	// addi 2 1 2
	r2++

p25:
	// seti 17 5 1
	r1 = 17

p26:
	// setr 2 6 5
	r5 = r2

p27:
	// seti 7 8 1
	r1 = 7

p28:
	// eqrr 4 0 2
	if r4 == r0 {
		r2 = 1
	} else {
		r2 = 0
	}

p29:
	// addr 2 1 1
	r1 += r2

p30:
	// seti 5 3 1
	r1 = 5
}
