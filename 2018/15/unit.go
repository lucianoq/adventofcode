package main

import (
	"sort"
)

const (
	Elf = iota
	Goblin
)

type Unit struct {
	Type int
	AP   int
	HP   int
	Pos  C
	Game *Game
}

func (u *Unit) IsCloseEnemy() bool {
	for _, c := range u.Pos.Near() {
		if unit, ok := u.Game.Units[c]; ok {
			if u.Type != unit.Type {
				return true
			}
		}
	}
	return false
}

func (u *Unit) OnlyEnemies(list []*Unit) []*Unit {
	res := make([]*Unit, 0)
	for _, u2 := range list {
		if u.Type != u2.Type {
			res = append(res, u2)
		}
	}
	return res
}

func (u *Unit) ThereIsEnemy(list []*Unit) bool {
	for _, u2 := range list {
		if u.Type != u2.Type {
			return true
		}
	}
	return false
}

func (u *Unit) targetsInRange() []C {
	targets := make([]C, 0)
	// all units
	for _, enemy := range u.Game.Units {
		// only enemies
		if enemy.Type != u.Type {
			// adjacent cells
			for _, adj := range enemy.Pos.Near() {
				// not a wall
				if u.Game.World[adj] {
					// not another unit
					if _, ok := u.Game.Units[adj]; !ok {
						targets = append(targets, adj)
					}
				}
			}
		}
	}
	return targets
}

func (u *Unit) nearestReachable() []C {
	distances := make(map[C]int)

	for _, c := range u.targetsInRange() {
		dist := u.distance(c)

		// unreachable
		if dist == -1 {
			continue
		}

		distances[c] = dist
	}

	minDistance := 1<<63 - 1
	for _, dist := range distances {
		if dist < minDistance {
			minDistance = dist
		}
	}

	nearest := make([]C, 0)
	for c, dist := range distances {
		if dist == minDistance {
			nearest = append(nearest, c)
		}
	}

	return nearest
}

// distance returns -1 if not reachable
// min steps otherwise
func (u *Unit) distance(c C) int {
	dist, _ := u.Game.bfs(u.Pos, c)
	return dist
}

func (u *Unit) Move() {
	nearest := u.nearestReachable()
	if len(nearest) == 0 {
		return
	}
	targetCell := readingOrder(nearest)[0]

	_, step := u.Game.bfs(u.Pos, targetCell)

	u.Game.Units[step] = u
	delete(u.Game.Units, u.Pos)
	u.Pos = step
}

func (u *Unit) Attack() {
	toAttack := unityToAttack(u.OnlyEnemies(u.Game.UnitsAround(u.Pos)))

	toAttack.HP -= u.AP

	if toAttack.HP <= 0 {
		delete(u.Game.Units, toAttack.Pos)
		switch toAttack.Type {
		case Elf:
			u.Game.NumElves--
		case Goblin:
			u.Game.NumGoblins--
		}
	}
}

func unityToAttack(list []*Unit) *Unit {
	sort.Slice(list, func(i, j int) bool {
		if list[i].HP == list[j].HP {
			if list[i].Pos.X == list[j].Pos.X {
				return list[i].Pos.Y < list[j].Pos.Y
			}
			return list[i].Pos.X < list[j].Pos.X
		}
		return list[i].HP < list[j].HP
	})

	return list[0]
}
