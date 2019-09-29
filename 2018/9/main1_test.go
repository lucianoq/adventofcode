package main

import (
	"testing"
)

var testInput = []struct {
	name    string
	players int
	marbles int
	score   int
}{
	{"Assignment", 9, 25, 32},
	{"A", 10, 1618, 8317},
	{"B", 13, 7999, 146373},
	{"C", 17, 1104, 2764},
	{"D", 21, 6111, 54718},
	{"E", 30, 5807, 37305},
}

func TestInput(t *testing.T) {
	for _, v := range testInput {
		t.Run(v.name, func(t *testing.T) {
			score := Solve(v.players, v.marbles)
			if score != v.score {
				t.Errorf("wanted %d, got %d", v.score, score)
			}
		})
	}
}
