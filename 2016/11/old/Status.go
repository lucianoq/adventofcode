package old

type Status struct {
	ID       string
	Elevator int
	Floors   [4]map[string]bool
}

var (
	startStatus = Status{
		ID: ""
		Elevator: 0,
		Floors: [4]map[string]bool{
			{
				"PoGen": true,
				"TmGen": true, "Tm": true,
				"PmGen": true,
				"RuGen": true, "Ru": true,
				"CoGen": true, "Co": true,
			},
			{
				"Po": true,
				"Pm": true,
			},
			{},
			{},
		},
	}
	finalStatus = Status{
		Elevator: 3,
		Floors: [4]map[string]bool{
			{},
			{},
			{},
			{
				"PoGen": true, "Po": true,
				"TmGen": true, "Tm": true,
				"PmGen": true, "Pm": true,
				"RuGen": true, "Ru": true,
				"CoGen": true, "Co": true,
			},
		},
	}
)

func main() {

	visited := map[Status]bool
	s := startStatus

}

func Equal(a, b Status) bool {
	if a.Elevator != b.Elevator {
		return false
	}

	for i := 0; i < 4; i++ {
		for k, v := range a.Floors[i] {
			if b.Floors[i][k] != v {
				return false
			}
		}
	}
	return true
}

func PossibleValidMoves(s Status) []Status {
	keys := make([]string, 0, len(s.Floors[s.Elevator]))
	for k := range s.Floors[s.Elevator] {
		keys = append(keys, k)
	}

	var moves [][]string

	for i := 0; i < len(keys); i++ {
		moves = append(moves, []string{keys[i]})
		for j := i + 1; j < len(keys); j++ {
			moves = append(moves, []string{keys[i], keys[j]})
		}
	}

	var next []Status

	for _, i := range []int{+1, -1} {
		oldFloor, newFloor := s.Elevator, s.Elevator+i
		if newFloor >= 0 && newFloor < 4 {
			for _, move := range moves {
				tmpStatus := s
				for _, item := range move {
					delete(tmpStatus.Floors[oldFloor], item)
					tmpStatus.Floors[newFloor][item] = true
				}
				if valid(tmpStatus.Floors[oldFloor]) && valid(tmpStatus.Floors[newFloor]) {
					next = append(next, tmpStatus)
				}
			}
		}
	}

	return next
}

func valid(floor map[string]bool) bool {
	var (
		thereIsGenerator bool
		chipAlone        bool
	)
	for k := range floor {
		if len(k) == 5 {
			thereIsGenerator = true
		}

		if len(k) == 2 {
			if !floor[k+"Gen"] {
				chipAlone = true
			}
		}
	}
	return !thereIsGenerator || !chipAlone
}

func Success(s Status) bool {
	return Equal(s, finalStatus)
}
