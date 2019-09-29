package main

import (
	"fmt"
	"os"
)

var a, b, c, d, e, f, g, h int

func main() {
	a = 1 // assignment

	b = 79
	c = b
	if a != 0 {
		b = b*100 + 100000
		c = b + 17000
	}

	// The code above generates these 2 numbers
	// b = 107900 // This will be incremented by 17
	// c = 124900 // This will never change

	// not in the code,
	// I need it to make it run the first time
	g = b - c

	// This is a for loop with (b=107900; b<=124900; b+=17)
	for {

		// This is a flag found
		f = 1

		// for loop with (d=2; d!=b; d++)
		d = 2
		for g != 0 {

			// for loop with (e=2; d*e!=b; e++)
			e = 2
			for g != 0 {

				// if (d*e == b) set the flag
				g = d*e - b
				if g == 0 {
					f = 0
				}

				// see for loop above
				e++
				g = e - b
			}

			// see for loop above
			d++
			g = d - b
		}

		// count if what we were looking has been found
		if f == 0 {
			h++
		}

		// Exit if the difference between b and c is 0
		// So, as we never change c, but we increment b by 17,
		// and there are 17000 between, it's a 1000 iteration loop
		g = b - c
		if g == 0 {
			fmt.Println(h)
			os.Exit(0)
		}

		// see for loop above
		b += 17
	}
}
