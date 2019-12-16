package main

import (
	"reflect"
	"testing"
)

func Test_apply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		reaction   Reaction
		needed     map[string]int
		excess     map[string]int
		want       map[string]int
		wantExcess map[string]int
	}{
		{
			name: "1",
			reaction: Reaction{
				Product:  Element{"AB", 1},
				Reagents: []Element{{"A", 1}, {"B", 1}},
			},
			needed:     map[string]int{"AB": 1},
			excess:     map[string]int{},
			want:       map[string]int{"A": 1, "B": 1},
			wantExcess: map[string]int{},
		},
		{
			name: "2",
			reaction: Reaction{
				Product:  Element{"AB", 1},
				Reagents: []Element{{"A", 2}, {"B", 2}},
			},
			needed:     map[string]int{"AB": 1},
			excess:     map[string]int{},
			want:       map[string]int{"A": 2, "B": 2},
			wantExcess: map[string]int{},
		},
		{
			name: "3",
			reaction: Reaction{
				Product:  Element{"AB", 2},
				Reagents: []Element{{"A", 2}, {"B", 2}},
			},
			needed:     map[string]int{"AB": 1},
			excess:     map[string]int{},
			want:       map[string]int{"A": 2, "B": 2},
			wantExcess: map[string]int{"AB": 1},
		},
		{
			name: "4",
			reaction: Reaction{
				Product:  Element{"AB", 1},
				Reagents: []Element{{"A", 2}, {"B", 3}},
			},
			needed:     map[string]int{"AB": 3},
			excess:     map[string]int{},
			want:       map[string]int{"A": 6, "B": 9},
			wantExcess: map[string]int{},
		},
		{
			name: "5",
			reaction: Reaction{
				Product:  Element{"AB", 5},
				Reagents: []Element{{"A", 1}, {"B", 1}},
			},
			needed:     map[string]int{"AB": 4},
			excess:     map[string]int{},
			want:       map[string]int{"A": 1, "B": 1},
			wantExcess: map[string]int{"AB": 1},
		},
		{
			name: "6",
			reaction: Reaction{
				Product:  Element{"AB", 5},
				Reagents: []Element{{"A", 1}, {"B", 1}},
			},
			needed:     map[string]int{"AB": 6},
			excess:     map[string]int{},
			want:       map[string]int{"A": 2, "B": 2},
			wantExcess: map[string]int{"AB": 4},
		},
		{
			name: "7",
			reaction: Reaction{
				Product:  Element{"A", 10},
				Reagents: []Element{{"ORE", 10}},
			},
			needed:     map[string]int{"A": 7, "D": 1, "ORE": 10},
			excess:     map[string]int{"A": 3},
			want:       map[string]int{"D": 1, "ORE": 20},
			wantExcess: map[string]int{"A": 6},
		},
	}

	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			apply(tt.reaction, tt.needed, tt.excess)

			if !reflect.DeepEqual(tt.needed, tt.want) {
				t.Errorf("needed got: %+v, wanted: %+v", tt.needed, tt.want)
			}

			if !reflect.DeepEqual(tt.excess, tt.wantExcess) {
				t.Errorf("excess got: %+v, wanted: %+v", tt.excess, tt.wantExcess)
			}
		})
	}
}
