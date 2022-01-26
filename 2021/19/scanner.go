package main

type Scanner struct {
	id       uint8
	beacons  []Point
	position Point
}

type Shift [3]int32

func (s Scanner) Transform(t Transformation) Scanner {
	beacons := make([]Point, len(s.beacons))

	for i, b := range s.beacons {
		beacons[i] = Point{
			t.Sign(0, b[t.Perm(0)]),
			t.Sign(1, b[t.Perm(1)]),
			t.Sign(2, b[t.Perm(2)]),
		}
	}

	return Scanner{
		id:      s.id,
		beacons: beacons,
	}
}

func (s Scanner) Shift(shift Shift) Scanner {
	newScanner := Scanner{
		id:      s.id,
		beacons: make([]Point, len(s.beacons)),
		position: Point{
			-shift[0],
			-shift[1],
			-shift[2],
		},
	}
	for i, b := range s.beacons {
		newScanner.beacons[i][0] = b[0] + shift[0]
		newScanner.beacons[i][1] = b[1] + shift[1]
		newScanner.beacons[i][2] = b[2] + shift[2]
	}
	return newScanner
}

func (s Scanner) Overlap(s2 Scanner) (bool, Shift) {
	distances0 := map[int32]uint8{}
	distances1 := map[int32]uint8{}
	distances2 := map[int32]uint8{}
	for _, b1 := range s.beacons {
		for _, b2 := range s2.beacons {
			shiftOn0 := b1[0] - b2[0]
			shiftOn1 := b1[1] - b2[1]
			shiftOn2 := b1[2] - b2[2]

			distances0[shiftOn0]++
			distances1[shiftOn1]++
			distances2[shiftOn2]++

			if distances0[shiftOn0] >= 12 && distances1[shiftOn1] >= 12 && distances2[shiftOn2] >= 12 {
				return true, Shift{
					shiftOn0,
					shiftOn1,
					shiftOn2,
				}
			}
		}
	}
	return false, Shift{}
}
