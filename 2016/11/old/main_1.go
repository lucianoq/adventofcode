package old

import (
	"fmt"
	"log"
	"math/bits"
	"os"
)

var visited = map[uint64]bool{}

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

//var chips = []uint16{Po, Pm, Tm, Ru, Co}
//var generators = []uint16{PoGen, PmGen, TmGen, RuGen, CoGen}
//var all = append(chips, generators...)

func main() {
	build := &Building{
		Floor(PoGen | TmGen | Tm | PmGen | RuGen | Ru | CoGen | Co),
		Floor(Po | Pm),
		0,
		0,
	}

	//printBuilding(floors)
	generate(0, build)
}

func generate(elev int, build *Building) {
	build.visit()

	log.Println("Starting building: ")
	build.print()

	if build.final() {
		build.print()
		os.Exit(0)
	}

	for _, delta := range []int{+1, -1} {
		newElev := elev + delta
		if newElev >= 0 && newElev < 4 {

			itemsInFloor := build[elev].all()

			//Move two items
			for _, p := range build[elev].allPairs() {
				newBuilding := build.move(elev, newElev, p.i, p.j)

				if newBuilding[elev].valid() && newBuilding[newElev].valid() {
					//log.Printf("Moving %0.13b and %0.13b from floor %d to floor %d", p.i, p.j, elev, newElev)
					//log.Println("VALID. Next move!")
					//newBuilding.print()
					if !newBuilding.visited() {
						//log.Println("never visited")
						generate(newElev, newBuilding)
					}
				}
			}

			// Move only 1 item
			for _, i := range itemsInFloor {

				newBuilding := build.move(elev, newElev, i)

				if newBuilding[elev].valid() && newBuilding[newElev].valid() {
					//log.Printf("Moving %0.13b from floor %d to floor %d", i, elev, newElev)
					//log.Println("VALID. Next move!")
					//newBuilding.print()
					if !newBuilding.visited() {
						//log.Println("never visited")
						generate(newElev, newBuilding)
					}
				}
			}
		}
	}
}

type Floor uint16

// Never a chip alone  together with other RTG.
func (f *Floor) valid() bool {
	chipsAlone := (*f & 0x00FF) &^ (*f >> 8) // are chips AND NOT contains gens
	return bits.OnesCount16(uint16(chipsAlone)) <= 0
}
func (f *Floor) contains(i uint16) bool {
	return i&uint16(*f) != 0
}

func (f *Floor) full() bool {
	return uint16(*f) == PoGen|TmGen|Tm|PmGen|RuGen|Ru|CoGen|Co|Po|Pm
}

func (f *Floor) all() []uint16 {
	list := make([]uint16, 0, 5)
	for i := uint16(0); i < 16; i++ {
		if uint16(*f)&(1<<i) != 0 {
			list = append(list, 1<<i)
		}
	}
	return list
}

type pair struct {
	i, j uint16
}

func (f *Floor) allPairs() []pair {
	list := make([]pair, 0, 5)

	row := uint16(*f)

	for i := uint16(0); i < 15; i++ {
		if row&(1<<i) != 0 {
			for j := i + 1; j < 16; j++ {
				if row&(1<<j) != 0 {
					list = append(list, pair{1 << i, 1 << j})
				}
			}
		}
	}
	return list
}

func (f *Floor) print() {
	fmt.Printf("%0.5b %0.5b\n", *f>>8, *f&0x00FF)
}

type Building [4]Floor

func (b *Building) print() {
	for i := 3; i >= 0; i-- {
		b[i].print()
	}
	fmt.Println()
}

func (b *Building) move(fromFloor, toFloor int, items ...uint16) *Building {
	newB := *b
	for _, i := range items {
		// remove item from floor
		newB[fromFloor] = Floor(uint16(newB[fromFloor]) & ^i)

		// add item to floor
		newB[toFloor] = Floor(uint16(newB[toFloor]) | i)
	}
	return &newB
}

func (b *Building) final() bool {
	return b[3].full()
}

func (b *Building) visit() {
	var x uint64
	for i := uint16(0); i < 4; i++ {
		x |= uint64(uint16(b[i]) << i)
	}
	visited[x] = true
}

func (b *Building) visited() bool {
	var x uint64
	for i := uint16(0); i < 4; i++ {
		x |= uint64(uint16(b[i]) << i)
	}
	return visited[x]
}
