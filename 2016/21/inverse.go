package main

import "bytes"

func InverseSwapPosition(s []byte, x, y int) []byte {
	return SwapPosition(s, y, x)
}

func InverseSwapLetters(s []byte, x, y byte) []byte {
	return SwapLetters(s, y, x)
}

func InverseMove(s []byte, x, y int) []byte {
	return Move(s, y, x)
}

func InverseReverse(s []byte, x, y int) []byte {
	return Reverse(s, x, y)
}

// This function works only on 8-letters strings
// But the assignment itself use 4 as threshold
func InverseRotateBasedOnLetter(s []byte, x byte) []byte {
	// cancel the +1 on all rotation
	s = RotateLeft(s, 1)

	idx := bytes.IndexByte(s, x)

	// all numbers with idx >= 4 are moved to idx+idx+1, so 2*idx+1, so odd
	// all numbers with idx<4 are moved to idx+idx, so 2*idx, so even
	if idx%2 != 0 {
		// this cancel the "+1" on idx>=4
		idx++

		// (idx + len) because we finished a loop and started again)
		// /2 cancel the operation 2*
		return RotateLeft(s, (idx+len(s))/2)

	} else {

		// idx/2 cancel the operation 2*
		return RotateLeft(s, idx/2)
	}

}

func InverseRotateLeft(s []byte, x int) []byte {
	return RotateLeft(s, -x)
}
