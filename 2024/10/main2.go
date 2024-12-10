package main

// find paths
func score(h map[P]uint8, p P) int {
	if h[p] == 9 {
		return 1
	}

	paths := 0
	for _, adj := range p.Neighbors() {
		if inMap(adj) {
			if h[adj] == h[p]+1 {
				paths += score(h, adj)
			}
		}
	}
	return paths
}
