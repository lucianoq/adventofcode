package main

func main() {
	stacks, moves := parse()

	for _, m := range moves {
		for i := 0; i < m.Qty; i++ {
			// pop
			last := len(stacks[m.From]) - 1
			elem := stacks[m.From][last]
			stacks[m.From] = stacks[m.From][:last]

			//push
			stacks[m.To] = append(stacks[m.To], elem)
		}
	}

	printTop(stacks)
}
