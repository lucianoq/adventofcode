package main

import (
	"testing"
)

func TestGame1(t *testing.T) {
	s := Status{
		HP:         10,
		Mana:       250,
		BossHP:     13,
		BossDamage: 8,
	}
	magics := []Magic{{Poison, 173}, {MagicMissile, 53}}

	got := game(s, magics)
	if got == false {
		t.Errorf("Player should win.")
	}
}

func TestGame2(t *testing.T) {
	s := Status{
		HP:         10,
		Mana:       250,
		BossHP:     14,
		BossDamage: 8,
	}
	magics := []Magic{{Recharge, 229}, {Shield, 113}, {Drain, 73}, {Poison, 173}, {MagicMissile, 53}}

	got := game(s, magics)
	if got == false {
		t.Errorf("Player should win.")
	}
}
