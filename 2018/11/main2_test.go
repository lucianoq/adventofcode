package main

import (
	"log"
	"testing"
)

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
	r := FindSquare(18)
	if r.PowerLevel != 113 || r.X != 90 || r.Y != 269 || r.SquareSize != 16 {
		log.Printf("Expecting %d in %d,%d,%d, obtained %d in %d,%d,%d",
			113, 90, 269, 16, r.PowerLevel, r.X, r.Y, r.SquareSize)
		t.FailNow()
	}
}

func TestFindSquare2(t *testing.T) {
	r := FindSquare(42)
	if r.PowerLevel != 119 || r.X != 232 || r.Y != 251 || r.SquareSize != 12 {
		log.Printf("Expecting %d in %d,%d,%d, obtained %d in %d,%d,%d",
			119, 232, 251, 12, r.PowerLevel, r.X, r.Y, r.SquareSize)
		t.FailNow()
	}
}
