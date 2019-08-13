package old

import "math/bits"

const (
	Po uint16 = 1 << iota
	Pm
	Tm
	Ru
	Co
)
const (
	PoGen uint16 = 1 << (iota + 8)
	PmGen
	TmGen
	RuGen
	CoGen
)

type Building uint64

func (b Building) Complete() bool {
	floor := b.Floor(3)
	return floor == Po|Pm|Tm|Ru|Co|PoGen|PmGen|TmGen|RuGen|CoGen
}

// Never a chip alone  together with other RTG.
func (b Building) Valid() bool {
	for i := 0; i < 4; i++ {
		floor := b.Floor(i)
		if !valid(floor) {
			return false
		}
	}
	return true
}

func (b Building) AllMovesResult(elevator int) []Building {
	outputs := make([]Building, 0, 4)

	currentFloor := b.Floor(elevator)

	for _, nextElevator := range []int{elevator + 1, elevator - 1} {
		if nextElevator >= 0 && nextElevator < 4 {
			nextFloor := b.Floor(nextElevator)
			for _, m := range moves(currentFloor) {
				if valid(nextFloor | m) {
					b2 := b.Move(nextElevator, m)
					outputs = append(outputs, b2)
				}
			}
		}
	}
	return outputs
}

func (b Building) Move(to int, items uint16) Building {
	from := b.Elevator()
	fromFloor := b.Floor(from)
	toFloor := b.Floor(to)
	fromFloor &= ^items
	toFloor |= items
	b2 := b
	b2 = b2.SetFloor(from, fromFloor)
	b2 = b2.SetFloor(to, toFloor)
	b2.SetElevator(to)
	return b2
}

func (b Building) Floor(i int) uint16 {
	return uint16(b >> (uint(i) * 16))
}

func (b Building) SetFloor(i int, floor uint16) Building {
	floor64 := uint64(floor) << uint(i)
	mask64 := uint64(0xFFFF) << uint(i)
	return Building((uint64(b) & ^mask64) | floor64)
}

func (b Building) SetElevator(elevator int) Building {
	elevator64 := uint64(elevator) << 62
	mask64 := uint64(0x3) << 62
	return Building((uint64(b) & ^mask64) | elevator64)
}

func (b Building) Elevator() int {
	return int(b >> 62)
}

func valid(floor uint16) bool {
	chipsAlone := (floor & 0x00FF) &^ (floor >> 8)
	return bits.OnesCount16(chipsAlone) == 0
}

var visited = map[uint64]bool{}

func moves(floor uint16) []uint16 {
	moves := make([]uint16, 0, 4)
	for i := uint(0); i < 16; i++ {
		itemI := uint16(1 << i)
		if floor&itemI != 0 {
			if valid(floor & ^itemI) {
				moves = append(moves, itemI)
			}
			for j := i + 1; j < 16; j++ {
				itemJ := uint16(1 << j)
				if floor&itemJ != 0 {
					if valid(floor & ^(itemI | itemJ)) {
						moves = append(moves, itemI|itemJ)
					}
				}
			}
		}
	}
	return moves
}

func main() {
	var start Building = 0
	start.SetFloor(0, PoGen|TmGen|Tm|PmGen|RuGen|Ru|CoGen|Co)
	start.SetFloor(1, Po|Pm)

	toVisit := []Building{start}
	visited := []Building{}

	moves := 0
	elevator := 0

}
