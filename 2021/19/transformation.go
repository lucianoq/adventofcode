package main

type Transformation [3]int8

func (t Transformation) Perm(x int8) int8 {
	if t[x] < 0 {
		return -t[x] - 1
	}
	return t[x] - 1
}

func (t Transformation) Sign(x int8, val int32) int32 {
	if t[x] < 0 {
		return -val
	}
	return val
}

var (
	Transformations = [][3]int8{
		// 0 permutations, 0 or 2 minus
		{+1, +2, +3},
		{-1, -2, +3},
		{+1, -2, -3},
		{-1, +2, -3},

		// 1 permutations, 1 or 3 minus
		{-1, +3, +2},
		{+1, -3, +2},
		{+1, +3, -2},
		{-1, -3, -2},

		// 1 permutations, 1 or 3 minus
		{-2, +1, +3},
		{+2, -1, +3},
		{+2, +1, -3},
		{-2, -1, -3},

		// 2 permutations, 0 or 2 minus
		{+2, +3, +1},
		{-2, -3, +1},
		{+2, -3, -1},
		{-2, +3, -1},

		// 2 permutations, 0 or 2 minus
		{+3, +1, +2},
		{-3, -1, +2},
		{+3, -1, -2},
		{-3, +1, -2},

		// 1 permutations, 1 or 3 minus
		{-3, +2, +1},
		{+3, -2, +1},
		{+3, +2, -1},
		{-3, -2, -1},
	}
)
