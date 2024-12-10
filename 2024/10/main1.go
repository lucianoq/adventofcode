package main

// find peaks
func score(h map[P]uint8, p P) int {
	peaks := map[P]struct{}{}
	findPeaksRecursive(h, p, peaks)
	return len(peaks)
}

func findPeaksRecursive(h map[P]uint8, p P, peaks map[P]struct{}) {
	if h[p] == 9 {
		peaks[p] = struct{}{}
		return
	}

	for _, adj := range p.Neighbors() {
		if inMap(adj) {
			if h[adj] == h[p]+1 {
				findPeaksRecursive(h, adj, peaks)
			}
		}
	}
}
