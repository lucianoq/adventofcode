package main

import "fmt"

type G struct {
	ID         int
	NumUnits   int
	HP         int
	AttDamage  int
	AttType    AttackType
	Weaknesses []AttackType
	Immunities []AttackType
	Initiative int
	Team       Team
}

func (g G) String() string {
	return fmt.Sprintf(
		"(id=%d) Units:%5d,\tHP:%6d,\tEP:%6d,\tIni:%4d,\tAD:%5d,\tAT:%s,\timmune:%s,\t\tweak:%s [%s]",
		g.ID,
		g.NumUnits,
		g.HP,
		g.EffectivePower(),
		g.Initiative,
		g.AttDamage,
		g.AttType,
		g.Immunities,
		g.Weaknesses,
		g.Team,
	)
}

func (g G) EffectivePower() int {
	return g.NumUnits * g.AttDamage
}

func (g G) WouldDeal(enemy *G) int {
	if g.AttType.IsIn(enemy.Immunities) {
		return 0
	}

	if g.AttType.IsIn(enemy.Weaknesses) {
		return g.EffectivePower() * 2
	}

	return g.EffectivePower()
}

func (g *G) Attacks(enemy *G) int {
	damage := g.WouldDeal(enemy)
	unitsKilled := damage / enemy.HP
	if unitsKilled > enemy.NumUnits {
		unitsKilled = enemy.NumUnits
	}
	enemy.NumUnits -= unitsKilled
	return unitsKilled
}
