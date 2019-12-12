package main

type Team int

const (
	ImmuneSystem Team = iota
	Infection
)

func (t Team) String() string {
	switch t {
	case ImmuneSystem:
		return "Immune System"
	case Infection:
		return "Infection"
	default:
		return "unknown"
	}
}
