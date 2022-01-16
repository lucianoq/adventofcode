package main

type Cell uint8

const (
	Wall Cell = iota
	Empty
	Amber
	Bronze
	Copper
	Desert
	Out
)

func (c Cell) String() string {
	switch c {
	case Empty:
		return "."
	case Wall:
		return "#"
	case Amber:
		return "A"
	case Bronze:
		return "B"
	case Copper:
		return "C"
	case Desert:
		return "D"
	case Out:
		return " "
	}
	panic("wrong cell")
}
