package main

import "sort"

type C struct{ X, Y int }

var nullC = C{-1, -1}

func (c C) Near() [4]C {
	return [4]C{
		{c.X - 1, c.Y},
		{c.X, c.Y - 1},
		{c.X, c.Y + 1},
		{c.X + 1, c.Y},
	}
}

func readingOrder(list []C) []C {
	sort.Slice(list, func(i, j int) bool {
		if list[i].X == list[j].X {
			return list[i].Y < list[j].Y
		}
		return list[i].X < list[j].X
	})
	return list
}
