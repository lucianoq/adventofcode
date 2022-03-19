package main

func recursive(model []int, z int) int {

	for digit := 1; digit <= 9; digit++ {

		// add a new digit to the list
		model := append(model, digit)
		digitIdx := len(model) - 1
		z := alu(z, digit, digitIdx)

		if uselessBranches.Contains(Status{digitIdx, z}) {
			// if we already unsuccessfully visited a situation
			// with this z at this level, and we didn't find a valid
			// sequel, there is no point on traversing it again
			continue
		}

		if len(model) == 14 {
			if z == 0 {
				return toInt(model)
			}

			// this is a leaf
			// we should not continue deeper
			continue
		}

		if sol := recursive(model, z); sol != 0 {
			return sol
		}
	}

	// no valid digit was found for this model
	// set it as useless
	uselessBranches.Add(Status{len(model) - 1, z})
	return 0
}
