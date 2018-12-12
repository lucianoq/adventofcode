package main

import "testing"

func TestCell_PowerLevel0(t *testing.T) {
	if PowerLevel(3, 5, 8) != 4 {
		t.FailNow()
	}
}

func TestCell_PowerLevel1(t *testing.T) {
	if PowerLevel(122, 79, 57) != -5 {
		t.FailNow()
	}
}

func TestCell_PowerLevel2(t *testing.T) {
	if PowerLevel(217, 196, 39) != 0 {
		t.FailNow()
	}
}

func TestCell_PowerLevel3(t *testing.T) {
	if PowerLevel(101, 153, 71) != 4 {
		t.FailNow()
	}
}

func TestFindSquare1(t *testing.T) {
	power, x, y, size := FindSquare(18)
	if power != 113 || x != 90 || y != 269 || size != 16 {
		t.FailNow()
	}
}

func TestFindSquare2(t *testing.T) {
	power, x, y, size := FindSquare(42)
	if power != 119 || x != 232 || y != 251 || size != 12 {
		t.FailNow()
	}
}
