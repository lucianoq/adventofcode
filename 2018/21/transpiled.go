package main

var a, p, b, c, d, e int

func main() {
	a = 0 // TODO

	// zero:
	// begin

	// d =     0b1111011 // 0
	// d &=  0b111001000 // 1
	// if d == 0b1001000 {
	// 	d = 1
	// } // 2
	// p += d // 3
	//
	//
	// goto zero
	// p = 0 // 4

	// if 0b1111011 & 0b111001000 == 0b1001000 {
	// 	goto begin
	// } else {
	// 	goto zero
	// }

	// begin:
	// begin real code
	p = 5

	d = 0 // 5
	p++

	for {
		e = d | 65536 //  10000000000000000 // 6
		d = 10704114  //  101000110101010011110010 // 7

		for {
			b = e & 255 // 11111111 // 8

		p9:
			d += b // 9

		p10:
			d &= 16777215 // 111111111111111111111111 // 10

		p11:
			d *= 65899 // 10000000101101011 // 11

		p12:
			d &= 16777215 // 111111111111111111111111 // 12

		p13:
			if e <= 256 { // 100000000
				// b = 1

				// p = 16
				// goto p16

				goto p28

			} else {
				// b = 0

				goto p17

			} // 13

			// p14:
			// 	p += b // 14
			// 	p++
			//
			// p15:
			// 	p++ // 15
			// 	p++
			//
			// p16:
			// 	p = 28
			// 	goto p28
			// p = 27 // 16
			// p++

		p17:
			b = 0 // 17

		p18:
			c = b + 1 // 18

		p19:
			c *= 256 // 19

		p20:
			if c > e {
				e = b // 26
			} else {
				b++ // 24
				goto p18
			} // 20
		}

	p28:
		if d == a {
			return
			// b = 1
		}
	}
}
