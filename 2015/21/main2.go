package main

import (
	"fmt"
)

func main() {
	boss := Player{}
	boss.HP, boss.Damage, boss.Armor = parse()

	maxCost := 0

	for _, w := range weapons {
		for _, a := range armors {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r1 != r2 {
						player := Player{
							HP:     100,
							Damage: w.Damage + a.Damage + r1.Damage + r2.Damage,
							Armor:  w.Armor + a.Armor + r1.Armor + r2.Armor,
						}

						if !win(player, boss) {
							cost := w.Cost + a.Cost + r1.Cost + r2.Cost
							if cost > maxCost {
								maxCost = cost
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(maxCost)
}
