package main

func main() {
	done := make(chan struct{})

	vm := NewVM("input")

	go vm.Run()

	go reader(vm.Output, done)

	// ( ¬A ∨ ¬B ∨ ¬C ) ∧ D
	instructions := []string{
		"NOT A J", // J = ¬A
		"NOT B T", // T = ¬B
		"OR T J",  // J = ¬A ∨ ¬B
		"NOT C T", // T = ¬C
		"OR T J",  // J = ¬A ∨ ¬B ∨ ¬C
		"AND D J", // J = ( ¬A ∨ ¬B ∨ ¬C ) ∧ D
		"WALK",
	}

	for _, i := range instructions {
		sendString(vm.Input, i)
	}

	<-done
}
