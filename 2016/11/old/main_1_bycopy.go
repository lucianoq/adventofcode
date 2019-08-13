package old

import (
	"fmt"
	"log"
	"math/bits"
	"os"
)

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

var chips = []uint16{Po, Pm, Tm, Ru, Co}
var generators = []uint16{PoGen, PmGen, TmGen, RuGen, CoGen}
var all = append(chips, generators...)

func main() {
	build := Building{
		Floor(PoGen | TmGen | Tm | PmGen | RuGen | Ru | CoGen | Co),
		Floor(Po | Pm),
		0,
		0,
	}

	//printBuilding(floors)
	for i := 12; ; i++ {
		log.Printf("Trying with %d moves\n", i)
		generate(i, 0, build)
	}
}

func generate(moves int, elev int, build Building) {
	//log.Println("Starting building: ")
	//build.print()

	if moves == 0 {
		return
	}

	if build.final() {
		fmt.Println(moves)
		build.print()
		os.Exit(0)
	}

	for _, delta := range []int{+1, -1} {
		newElev := elev + delta
		if newElev >= 0 && newElev < 4 {

			itemsInFloor := build[elev].all()

			//Move two items
			for _, i := range itemsInFloor {
				for _, j := range itemsInFloor {
					if i != j {
						//log.Printf("Moving %0.13b and %0.13b from floor %d to floor %d", i, j, elev, newElev)
						newBuilding := build.move(elev, newElev, i, j)
						//newBuilding.print()

						if newBuilding[elev].valid() && newBuilding[newElev].valid() {
							//log.Println("VALID. Next move!")
							generate(moves-1, newElev, newBuilding)
						} else {
							//log.Println("NOT VALID")
						}
					}
				}
			}

			// Move only 1 item
			for _, i := range itemsInFloor {

				//log.Printf("Moving %0.13b from floor %d to floor %d", i, elev, newElev)
				newBuilding := build.move(elev, newElev, i)
				//newBuilding.print()

				if newBuilding[elev].valid() && newBuilding[newElev].valid() {
					//log.Println("VALID. Next move!")
					generate(moves-1, newElev, newBuilding)
				} else {
					//log.Println("NOT VALID")
				}
			}
		}
	}
	return
}

type Floor uint16

// Never a chip alone  together with other RTG.
func (f Floor) valid() bool {
	chipsAlone := (f & 0x00FF) &^ (f >> 8) // are chips AND NOT contains gens
	return bits.OnesCount16(uint16(chipsAlone)) <= 0
}
func (f Floor) contains(i uint16) bool {
	return i&uint16(f) != 0
}

func (f Floor) full() bool {
	return uint16(f) == PoGen|TmGen|Tm|PmGen|RuGen|Ru|CoGen|Co|Po|Pm
}

func (f Floor) all() []uint16 {
	list := make([]uint16, 0, 5)
	for i := uint16(0); i < 16; i++ {
		if uint16(f)&(1<<i) != 0 {
			list = append(list, 1<<i)
		}
	}
	return list
}

func (f Floor) print() {
	fmt.Printf("%0.5b %0.5b\n", f>>8, f&0x00FF)
}

type Building [4]Floor

func (b Building) print() {
	for i := 3; i >= 0; i-- {
		b[i].print()
	}
	fmt.Println()
	//time.Sleep(100 * time.Millisecond)
}

func (b Building) move(fromFloor, toFloor int, items ...uint16) Building {
	for _, i := range items {
		// remove item from floor
		b[fromFloor] = Floor(uint16(b[fromFloor]) & ^i)

		// add item to floor
		b[toFloor] = Floor(uint16(b[toFloor]) | i)
	}
	return b
}

func (b Building) final() bool {
	return b[3].full()
}
