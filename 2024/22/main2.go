package main

import "fmt"

type Changes [4]int

func main() {
	secrets := parseInput()
	priceAfter := map[int]map[Changes]int{}
	discovered := map[Changes]map[int]struct{}{}

	for buyerID, secret := range secrets {

		// Initialise map
		if _, ok := priceAfter[buyerID]; !ok {
			priceAfter[buyerID] = map[Changes]int{}
		}

		// ring data structure, to reduce allocations
		ring, i := Changes{}, 0

		oldPrice := secret % 10
		for t := 1; t <= 2000; t++ {

			// generate a new secret
			secret = compute(secret)

			newPrice := secret % 10
			delta := newPrice - oldPrice
			oldPrice = newPrice

			// add the delta to the ring
			ring[i], i = delta, (i+1)%4

			// skip the first 3 elements, as there are
			// not enough deltas for a full change
			if t >= 4 {
				change := Changes{ring[i], ring[(i+1)%4], ring[(i+2)%4], ring[(i+3)%4]}

				// Initialise map
				if _, ok := discovered[change]; !ok {
					discovered[change] = map[int]struct{}{}
				}
				discovered[change][buyerID] = struct{}{}

				// insert only the first occurrence, ignore the next ones
				if _, ok := priceAfter[buyerID][change]; !ok {
					priceAfter[buyerID][change] = newPrice
				}
			}
		}
	}

	maxBananas := 0
	for change, affectedIDs := range discovered {
		bananas := 0
		for buyerID := range affectedIDs {
			bananas += priceAfter[buyerID][change]
		}
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}

	fmt.Println(maxBananas)
}
