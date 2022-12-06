package main

func findMarker(line string, length int) int {
	seen := map[byte]int{}
	for i := 0; i < len(line); i++ {
		l, ok := seen[line[i]]
		if ok {
			for k, v := range seen {
				if v < l {
					delete(seen, k)
				}
			}
		}

		seen[line[i]] = i
		if len(seen) == length {
			return i + 1
		}
	}
	return -1
}
