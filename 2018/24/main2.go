package main

import (
	"fmt"
	"os"
)

func main() {
	groups := parse(os.Stdin)

Boost:
	for i := 0; ; i++ {
		groups := boost(groups, i)

		for !finished(groups) {
			attackPairs := TargetSelectionPhase(groups)
			killed := AttackPhase(attackPairs)
			if killed == 0 {
				// we found a loop. No one is able to improve their situation
				// we need to skip this cycle
				continue Boost
			}
			groups = filterOutDead(groups)
		}

		if immuneSystemWon(groups) {
			fmt.Println(unitsLeft(groups))
			return
		}
	}
}

func boost(groups []*G, b int) []*G {
	newGroups := make([]*G, len(groups))
	for i, g := range groups {
		newGroups[i] = g.Copy()
		if newGroups[i].Team == ImmuneSystem {
			newGroups[i].AttDamage += b
		}
	}
	return newGroups
}
