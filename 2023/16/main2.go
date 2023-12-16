package main

import "fmt"

func main() {
	m := parse()

	maxEnergy := 0

	for r := 0; r < m.MaxR; r++ {
		if nrg := energy(m, Beam{r, -1, E}); nrg > maxEnergy {
			maxEnergy = nrg
		}

		if nrg := energy(m, Beam{r, m.MaxC, W}); nrg > maxEnergy {
			maxEnergy = nrg
		}
	}

	for c := 0; c < m.MaxC; c++ {
		if nrg := energy(m, Beam{-1, c, S}); nrg > maxEnergy {
			maxEnergy = nrg
		}

		if nrg := energy(m, Beam{m.MaxR, c, N}); nrg > maxEnergy {
			maxEnergy = nrg
		}
	}

	fmt.Println(maxEnergy)
}
