package main

func main() {
	stacks, moves := parse()

	for _, m := range moves {
		first := len(stacks[m.From]) - m.Qty
		last := len(stacks[m.From])
		stacks[m.To] = append(stacks[m.To], stacks[m.From][first:last]...)
		stacks[m.From] = stacks[m.From][:first]
	}

	printTop(stacks)
}
