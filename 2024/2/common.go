package main

import "strconv"

func toInt(xs []string) ([]int, error) {
	l := make([]int, len(xs))
	for i, x := range xs {
		v, err := strconv.Atoi(x)
		if err != nil {
			return nil, err
		}
		l[i] = v
	}
	return l, nil
}

func safe(xs []int) bool {
	asc := xs[1]-xs[0] > 0

	for i := 1; i < len(xs); i++ {
		incr := xs[i] - xs[i-1]

		if incr > 0 != asc {
			return false
		}

		if incr < 0 {
			incr = -incr
		}

		if incr < 1 || incr > 3 {
			return false
		}
	}
	return true
}
