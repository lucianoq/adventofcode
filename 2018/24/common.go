package main

import "sort"

func unitsLeft(groups []*G) int {
	sum := 0
	for _, g := range groups {
		sum += g.NumUnits
	}
	return sum
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

func finished(groups []*G) bool {
	var countImmune, countInfection int
	for _, g := range groups {
		switch g.Team {
		case ImmuneSystem:
			countImmune++
		case Infection:
			countInfection++
		}
	}
	return countImmune == 0 || countInfection == 0
}

func immuneSystemWon(groups []*G) bool {
	for _, g := range groups {
		if g.Team == Infection {
			return false
		}
	}
	return true
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
				attack[g1] = g2
				inf = inf[1:]
			}
		}
	}

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

// AttackPhase
// During the attacking phase, each group deals damage to the target it
// selected, if any. Groups attack in decreasing order of initiative,
// regardless of whether they are part of the infection or the immune system.
// (If a group contains no units, it cannot attack.)
//
// The damage an attacking group deals to a defending group depends on the
// attacking group's attack type and the defending group's immunities and
// weaknesses. By default, an attacking group would deal damage equal to its
// effective power to the defending group.
// The defending group only loses whole units from damage; damage is always
// dealt in such a way that it kills the most units possible, and any remaining
// damage to a unit that does not immediately kill it is ignored. For example,
// if a defending group contains 10 units with 10 hit points each and receives
// 75 damage, it loses exactly 7 units and is left with 3 units at full health.
//
// After the fight is over, if both armies still contain units, a new fight
// begins; combat only ends once one army has lost all of its units.
func AttackPhase(attackPairs map[*G]*G) int {
	attackers := keys(attackPairs)

	// Groups attack in decreasing order of initiative, regardless of whether they
	// are part of the infection or the immune system.
	//   sort by initiative
	sort.Slice(attackers, func(i, j int) bool {
		return attackers[i].Initiative > attackers[j].Initiative
	})

	countKilled := 0
	for _, g1 := range attackers {
		if g1.NumUnits > 0 {
			g2 := attackPairs[g1]
			killed := g1.Attacks(g2)
			countKilled += killed
		}
	}
	return countKilled
}

func keys(m map[*G]*G) []*G {
	ks := make([]*G, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
