package main

func findMarker(line string, length int) int {
	var seen int32

	var left, right = 0, 0

	// move `right` forward
	for ; ; right++ {

		// if we find a duplicate character
		if seen&(1<<(line[right]-'a')) != 0 {

			// we move `left` forward, up to the next
			// of the first occurrence of the duplicate
			for ; ; left++ {

				if line[left] == line[right] {
					left++
					break
				}
				seen &^= 1 << (line[left] - 'a')
			}

			// no need to add on `seen` (we didn't delete it)
			// no need to check the length cause moving left
			// shorten the string
			continue
		}

		seen |= 1 << (line[right] - 'a')

		// if [left,right) is long enough, that's our goal
		if right-left+1 == length {
			return right + 1
		}
	}
}
