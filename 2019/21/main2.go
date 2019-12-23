package main

func main() {
	done := make(chan struct{})

	vm := NewVM("input")

	go vm.Run()

	go reader(vm.Output, done)

	// ( ¬A ∨ ¬B ∨ ¬C ) ∧ D ∧ (E ∨ H)
	instructions := []string{
		"NOT A J", // J = ¬A
		"NOT B T", // T = ¬B
		"OR T J",  // J = ¬A ∨ ¬B
		"NOT C T", // T = ¬C
		"OR T J",  // J = ¬A ∨ ¬B ∨ ¬C
		"AND D J", // J = ( ¬A ∨ ¬B ∨ ¬C ) ∧ D
		"NOT A T", // T = ?
		"AND A T", // T = false
		"OR E T",  // T = false ∨ E = E
		"OR H T",  // T = E ∨ H
		"AND T J", // J = ( ¬A ∨ ¬B ∨ ¬C ) ∧ D ∧ (E ∨ H)
		"RUN",
	}

	for _, i := range instructions {
		sendString(vm.Input, i)
	}

	<-done
}
