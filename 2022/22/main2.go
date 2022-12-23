package main

func (h *Human) Walk() bool {
	dirDelta := Dirs[h.Facing]
	next := P{h.Curr.X + dirDelta.X, h.Curr.Y + dirDelta.Y}
	if wall, ok := Map[next]; ok {
		if wall {
			return false
		}
		h.Curr = next
		return true
	}

	next, nextFacing := crossBorder(next, h.Facing)

	if Map[next] {
		return false
	}

	h.Curr = next
	h.Facing = nextFacing
	return true
}

// Input:
//       +--A--+--B--+
//       |     |     |
//       I     |     N
//       |     |     |
//       +-----+--C--+
//       |     |
//       J     L
//       |     |
// +--D--+-----+
// |     |     |
// G     |     M
// |     |     |
// +-----+--E--+
// |     |
// H     K
// |     |
// +--F--+

// A <--> H  // transpose and shift
// D <--> J  // transpose and shift
// L <--> C  // transpose and shift
// I <--> G  // both vertical, invert on x
// N <--> M  // both vertical, invert on x
// B <--> F  // both horizontal, easy shift
// E <--> K  // transpose and shift

// Transposed
// +-----+-----+--G--+--H--+
// |     |     |     |     |
// |     |     D     |     F
// |     |     |     |     |
// +--I--+--J--+-----+--K--+
// |     |     |     |     |
// A     |     |     E     |
// |     |     |     |     |
// +-----+--L--+--M--+-----+
// |     |     |     |     |
// B     C     |     |     |
// |     |     |     |     |
// +--N--+-----+-----+-----+
// |     |     |     |     |
// |     |     |     |     |
// |     |     |     |     |
// +-----+-----+-----+-----+
func crossBorder(n P, dir Dir) (P, Dir) {
	x, y := n.X, n.Y

	switch {

	// A --> H
	case x == -1 && y < 2*Size:
		return P{y + 2*Size, x + 1}, E

	// B --> F
	case x == -1 && y >= 2*Size:
		return P{x + 4*Size, y - 2*Size}, N

	// C --> L
	case x == Size && dir == S:
		return P{y - Size, x + Size - 1}, W

	// D --> J
	case x == 2*Size-1 && dir == N:
		return P{y + Size, x - Size + 1}, E

	// E --> K
	case x == 3*Size && dir == S:
		return P{y + 2*Size, x - 2*Size - 1}, W

	// F --> B
	case x == 4*Size:
		return P{x - 4*Size, y + 2*Size}, S

	// G --> I
	case y == -1 && x < 3*Size:
		return P{3*Size - 1 - x, y + Size + 1}, E

	// H --> A
	case y == -1 && x >= 3*Size:
		return P{y + 1, x - 2*Size}, S

	// I --> G
	case y == Size-1 && x < Size:
		return P{3*Size - 1 - x, y - Size + 1}, E

	// J --> D
	case y == Size-1 && x >= Size && dir == W:
		return P{y + Size + 1, x - Size}, S

	// K --> E
	case y == Size && dir == E:
		return P{y + 2*Size - 1, x - 2*Size}, N

	// L --> C
	case y == 2*Size && x < 2*Size && dir == E:
		return P{y - Size - 1, x + Size}, N

	// M --> N
	case y == 2*Size && x >= 2*Size:
		return P{3*Size - 1 - x, y + Size - 1}, W

	// N --> M
	case y == 3*Size:
		return P{3*Size - 1 - x, y - Size - 1}, W
	}

	panic("not a border crossing")
}
