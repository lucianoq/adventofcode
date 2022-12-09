package main

type P struct{ X, Y int }

func (p P) Follow(h P) P {
	diffX, diffY := h.X-p.X, h.Y-p.Y
	if abs(diffX) <= 1 && abs(diffY) <= 1 {
		return p
	}
	return P{p.X + sign(diffX), p.Y + sign(diffY)}
}

func sign(i int) int {
	if i == 0 {
		return 0
	}
	if i > 0 {
		return 1
	}
	return -1
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
