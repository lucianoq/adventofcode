package main

func (h *Human) Walk() bool {
	dir := Dirs[h.Facing]

	next := P{h.Curr.X + dir.X, h.Curr.Y + dir.Y}
	if wall, ok := Map[next]; ok {
		if wall {
			return false
		}
		h.Curr = next
		return true
	}

	// fast-forward in opposite direction
	oppDir := P{-dir.X, -dir.Y}
	for {
		lookAhead := P{next.X + oppDir.X, next.Y + oppDir.Y}
		if _, ok := Map[lookAhead]; !ok {

			if Map[next] {
				return false
			}
			h.Curr = next
			return true
		}
		next = lookAhead
	}
}
