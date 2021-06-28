package main

import (
	"fmt"
	"os"
)

func main() {
	groups := parse(os.Stdin)

	for !finished(groups) {
		attackPairs := TargetSelectionPhase(groups)
		AttackPhase(attackPairs)
		groups = filterOutDead(groups)
	}

	fmt.Println(unitsLeft(groups))
}
