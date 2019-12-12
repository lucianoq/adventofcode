package main

import (
	"fmt"
	"sort"
)

func main() {
	groups := parse()

	for _, g := range groups {
		fmt.Println(g)
	}

	for !finished(groups) {
		StatusPhase(groups)

		attackPairs := TargetSelectionPhase(groups)

		AttackPhase(attackPairs)

		groups = filterOutDead(groups)

		fmt.Println()
		fmt.Println()
	}

	StatusPhase(groups)

	sum := 0
	for _, g := range groups {
		sum += g.NumUnits
	}
	fmt.Println(sum)
}

func filterOutDead(gs []*G) []*G {
	newGs := make([]*G, 0)
	for _, g := range gs {
		if g.NumUnits > 0 {
			newGs = append(newGs, g)
		}
	}
	return newGs
}

func StatusPhase(gs []*G) {
	fmt.Println("Immune System")
	for _, g := range gs {
		if g.Team == ImmuneSystem {
			fmt.Printf("Group %d contains %d units\n", g.ID, g.NumUnits)
		}
	}

	fmt.Println("Infection")
	for _, g := range gs {
		if g.Team == Infection {
			fmt.Printf("Group %d contains %d units\n", g.ID, g.NumUnits)
		}
	}

	fmt.Println()
}

func finished(groups []*G) bool {
	var countIS, countInf int
	for _, g := range groups {
		switch g.Team {
		case ImmuneSystem:
			countIS++
		case Infection:
			countInf++
		}
	}
	return countIS == 0 || countInf == 0
}

func TargetSelectionPhase(groups []*G) map[*G]*G {
	attack := make(map[*G]*G)

	is, inf := splitTeams(groups)

	sortByChoosingPriority(inf)

	for _, g1 := range inf {
		if len(is) > 0 {
			sortByTargetPriority(g1, is)

			g2 := is[0]

			if g1.WouldDeal(g2) > 0 {
				fmt.Printf("%s group %d would deal defending group %d %d damage\n", g1.Team, g1.ID, g2.ID, g1.WouldDeal(g2))
				attack[g1] = g2
				is = is[1:]
			}
		}
	}

	is, inf = splitTeams(groups)

	sortByChoosingPriority(is)

	for _, g1 := range is {
		if len(inf) > 0 {
			sortByTargetPriority(g1, inf)

			g2 := inf[0]

			if g1.WouldDeal(g2) > 0 {
				fmt.Printf("%s group %d would deal defending group %d %d damage\n", g1.Team, g1.ID, g2.ID, g1.WouldDeal(g2))
				attack[g1] = g2
				is = is[1:]
			}
		}
	}

	fmt.Println()

	return attack
}

func splitTeams(groups []*G) ([]*G, []*G) {
	is, inf := make([]*G, 0), make([]*G, 0)
	for _, g := range groups {
		switch g.Team {
		case ImmuneSystem:
			is = append(is, g)
		case Infection:
			inf = append(inf, g)
		}
	}
	return is, inf
}

// In decreasing order of effective power, groups choose their targets;
// in a tie, the group with the higher initiative chooses first.
func sortByChoosingPriority(g []*G) {
	// sort slice by choose priority
	sort.Slice(g, func(i, j int) bool {
		if g[i].EffectivePower() == g[j].EffectivePower() {
			return g[i].Initiative > g[j].Initiative
		}
		return g[i].EffectivePower() > g[j].EffectivePower()
	})
}

// 1. most dealt damage
// 2. defending largest effective power
// 3. defending highest initiative.
// If it cannot deal any defending groups damage, it does not choose a target.
func sortByTargetPriority(by *G, g []*G) {
	sort.Slice(g, func(i, j int) bool {
		dealI := by.WouldDeal(g[i])
		dealJ := by.WouldDeal(g[j])

		if dealI == dealJ {
			if g[i].EffectivePower() == g[j].EffectivePower() {
				return g[i].Initiative > g[j].Initiative
			}
			return g[i].EffectivePower() > g[j].EffectivePower()
		}

		return dealI > dealJ
	})
}

func AttackPhase(attackPairs map[*G]*G) {
	attackers := keys(attackPairs)
	sortByAttackingPriority(attackers)

	for _, g1 := range attackers {
		if g1.NumUnits > 0 {
			g2 := attackPairs[g1]
			killed := g1.Attacks(g2)
			fmt.Printf("%s group %d attacks defending group %d, killing %d units\n", g1.Team, g1.ID, g2.ID, killed)
		}
	}
	fmt.Println()
}

// Groups attack in decreasing order of initiative, regardless of whether they
// are part of the infection or the immune system.
func sortByAttackingPriority(g []*G) {
	sort.Slice(g, func(i, j int) bool {
		return g[i].Initiative > g[j].Initiative
	})
}

func keys(m map[*G]*G) []*G {
	ks := make([]*G, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
