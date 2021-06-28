package main

type G struct {
	ID         int
	NumUnits   int
	HP         int
	AttDamage  int
	AttType    AttackType
	Initiative int
	Weaknesses map[AttackType]struct{}
	Immunities map[AttackType]struct{}
	Team       Team
}

func (g *G) Copy() *G {
	return &G{
		ID:         g.ID,
		NumUnits:   g.NumUnits,
		HP:         g.HP,
		AttDamage:  g.AttDamage,
		AttType:    g.AttType,
		Initiative: g.Initiative,
		Weaknesses: g.Weaknesses,
		Immunities: g.Immunities,
		Team:       g.Team,
	}
}

func (g *G) EffectivePower() int {
	return g.NumUnits * g.AttDamage
}

// WouldDeal
// if the defending group is immune to the attacking group's attack type,
// the defending group instead takes no damage; if the defending group is
// weak to the attacking group's attack type, the defending group instead
// takes double damage.
func (g *G) WouldDeal(enemy *G) int {
	if _, ok := enemy.Immunities[g.AttType]; ok {
		return 0
	}

	if _, ok := enemy.Weaknesses[g.AttType]; ok {
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
