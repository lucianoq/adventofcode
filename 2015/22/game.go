package main

type Status struct {
	// Player
	HP     int
	Mana   int
	Shield int

	// Timers
	ShieldTimer   int
	PoisonTimer   int
	RechargeTimer int

	// Boss
	BossHP     int
	BossDamage int
}

type Magic struct {
	Spell func(Status) Status
	Cost  int
}

var AllMagics = []Magic{
	{MagicMissile, 53},
	{Drain, 73},
	{Shield, 113},
	{Poison, 173},
	{Recharge, 229},
}

var cheapestSpell = 53

func Game(magics []Magic) bool {
	initialStatus := Status{
		HP:         50,
		Mana:       500,
		BossHP:     58,
		BossDamage: 9,
	}
	return game(initialStatus, magics)
}

func game(s Status, magics []Magic) bool {
	var gameFinished, gameWon bool

	for i := 0; i < len(magics); i++ {
		// Player turn
		s, gameFinished, gameWon = Turn(s, magics[i], true)

		if gameFinished {
			return gameWon
		}

		//fmt.Println()

		// Boss turn
		s, gameFinished, gameWon = Turn(s, Magic{}, false)
		if gameFinished {
			return gameWon
		}

		//fmt.Println()
	}

	return false
}

func Turn(s Status, magic Magic, isPlayer bool) (Status, bool, bool) {
	if isPlayer {
		//fmt.Println("-- Player Turn --")
	} else {
		//fmt.Println("-- Boss Turn --")
	}

	if HardMode {
		s.HP -= 1
		if lost(s) {
			//fmt.Println("You lost")
			return s, true, false
		}
	}

	//fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", s.HP, s.Shield, s.Mana)
	//fmt.Printf("- Boss has %d hit points\n", s.BossHP)
	s = ApplyEffects(s)
	if won(s) {
		//fmt.Println("This kills the boss, and the player wins.")
		return s, true, true
	}

	if isPlayer {
		if s.Mana < cheapestSpell {
			return s, true, false
		}

		s = magic.Spell(s)
	} else {
		s = BossAttack(s)
	}

	if won(s) {
		//fmt.Println("This kills the boss, and the player wins.")
		return s, true, true
	}

	if lost(s) {
		//fmt.Println("You lost")
		return s, true, false
	}

	//fmt.Fprintln(w)
	return s, false, false
}

// Spells

func MagicMissile(s Status) Status {
	s.Mana -= 53
	s.BossHP -= 4
	//fmt.Printf("Player casts Magic Missile, dealing 4 damage.\n")
	return s
}

func Drain(s Status) Status {
	s.Mana -= 73
	s.BossHP -= 2
	s.HP += 2
	//fmt.Printf("Player casts Drain, dealing 2 damage, and healing 2 hit points.\n")
	return s
}

func Shield(s Status) Status {
	s.Mana -= 113
	s.ShieldTimer = 6
	s.Shield += 7
	//fmt.Printf("Player casts Shield, increasing armor by 7.\n")
	return s
}

func Poison(s Status) Status {
	s.Mana -= 173
	s.PoisonTimer = 6
	//fmt.Printf("Player casts Poison.\n")
	return s
}

func Recharge(s Status) Status {
	s.Mana -= 229
	s.RechargeTimer = 5
	//fmt.Printf("Player casts Recharge.\n")
	return s
}

func BossAttack(s Status) Status {
	if s.Shield > 0 {
		hpDealt := dealt(s.BossDamage - s.Shield)
		s.HP -= hpDealt
		//fmt.Printf("Boss attacks for %d - %d = %d damage!.\n", s.BossDamage, s.Shield, hpDealt)
	} else {
		s.HP -= s.BossDamage
		//fmt.Printf("Boss attacks for %d damage.\n", s.BossDamage)
	}

	return s
}

func ApplyEffects(s Status) Status {
	if s.ShieldTimer > 0 {
		s.ShieldTimer--
		//fmt.Printf("Shield's timer is now %d.\n", s.ShieldTimer)

		if s.ShieldTimer == 0 {
			s.Shield -= 7
			//fmt.Printf("Shield wears off, decreasing armor by 7.\n")
		}
	}

	if s.RechargeTimer > 0 {
		s.Mana += 101
		s.RechargeTimer--
		//fmt.Printf("Recharge provides 101 mana; its timer is now %d.\n", s.RechargeTimer)

		if s.RechargeTimer == 0 {
			//fmt.Printf("Recharge wears off.\n")
		}
	}

	if s.PoisonTimer > 0 {
		s.BossHP -= 3
		s.PoisonTimer--
		//fmt.Printf("Poison deals 3 damage; its timer is now %d.\n", s.PoisonTimer)

		if s.PoisonTimer == 0 {
			//fmt.Printf("Poison wears off.\n")
		}
	}

	return s
}

// Utils

func won(s Status) bool {
	if s.BossHP <= 0 {
		return true
	}
	return false
}

func lost(s Status) bool {
	if s.HP <= 0 {
		return true
	}
	return false
}

func dealt(x int) int {
	if x < 1 {
		return 1
	}
	return x
}
