package main

func (s Status) Complete(numPoI, startX, startY int) bool {

	// All bits 1 for all the PoI found in the map
	// i.e. 00001110 if the map contains PoI 1,2,3
	var mask uint64 = 1<<uint(numPoI+1) - 2

	return s.PoICollected == mask && s.X == startX && s.Y == startY
}
