package main

const (
	input    = "394618527"
	min, max = 1, cupsSize
)

func move(r *Ring) {
	curr := r.Value()
	pickedUp := r.PickUp3()

	destination := findDestination(r, pickedUp)

	r.MoveTo(destination)

	for i := 0; i < 3; i++ {
		r.Insert(pickedUp[i])
		r.Next()
	}

	r.MoveTo(curr)
	r.Next()
}

func findDestination(r *Ring, cups [3]int) int {
	max := max
	for in(max, cups) {
		max--
	}

	val := r.Value() - 1
	if val < min {
		return max
	}

	for in(val, cups) {
		val--
		if val < min {
			return max
		}
	}

	return val
}

func in(val int, r [3]int) bool {
	for i := 0; i < 3; i++ {
		if r[i] == val {
			return true
		}
	}
	return false
}
