package main

import "fmt"

func main() {
	m := parse()

	sum := 0
	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			c := C{i, j}
			ch := m[c]

			if ch == '*' {
				sum += gearRatio(m, c)
			}
		}
	}
	fmt.Println(sum)
}

func (c C) Left() C   { return C{c.i, c.j - 1} }
func (c C) Right() C  { return C{c.i, c.j + 1} }
func (c C) Top() C    { return C{c.i - 1, c.j} }
func (c C) Bottom() C { return C{c.i + 1, c.j} }

func gearRatio(m map[C]byte, c C) int {

	numbers := []int{}

	if isDigit(m[c.Left()]) {
		numbers = append(numbers, expandNumber(m, c.Left()))
	}

	if isDigit(m[c.Right()]) {
		numbers = append(numbers, expandNumber(m, c.Right()))
	}

	if isDigit(m[c.Top()]) {
		numbers = append(numbers, expandNumber(m, c.Top()))
	} else {
		if isDigit(m[c.Top().Left()]) {
			numbers = append(numbers, expandNumber(m, c.Top().Left()))
		}
		if isDigit(m[c.Top().Right()]) {
			numbers = append(numbers, expandNumber(m, c.Top().Right()))
		}
	}

	if isDigit(m[c.Bottom()]) {
		numbers = append(numbers, expandNumber(m, c.Bottom()))
	} else {
		if isDigit(m[c.Bottom().Left()]) {
			numbers = append(numbers, expandNumber(m, c.Bottom().Left()))
		}
		if isDigit(m[c.Bottom().Right()]) {
			numbers = append(numbers, expandNumber(m, c.Bottom().Right()))
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}

	return 0
}

func expandNumber(m map[C]byte, c C) int {
	for isDigit(m[c.Left()]) {
		c.j--
	}
	num := 0
	for ; isDigit(m[C{c.i, c.j}]); c.j++ {
		num = 10*num + int(m[c]-'0')
	}
	return num
}
