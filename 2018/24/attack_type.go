package main

type AttackType int

const (
	Bludgeoning AttackType = iota
	Fire
	Radiation
	Slashing
	Cold
)

func (at AttackType) String() string {
	switch at {
	case Bludgeoning:
		return "bludgeoning"
	case Fire:
		return "fire"
	case Radiation:
		return "radiation"
	case Slashing:
		return "slashing"
	case Cold:
		return "cold"
	default:
		return "unknown"
	}
}

func (at AttackType) IsIn(list []AttackType) bool {
	for _, t := range list {
		if at == t {
			return true
		}
	}
	return false
}
