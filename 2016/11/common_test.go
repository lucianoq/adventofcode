package main

import (
	"reflect"
	"testing"
)

func TestFloor_Valid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		atoms []Atom
		want  bool
	}{
		{
			name:  "empty building is valid",
			atoms: nil,
			want:  true,
		},
		{
			name:  "Chip alone is valid",
			atoms: []Atom{PoloniumChip},
			want:  true,
		},
		{
			name:  "Chip and RTG alone is valid",
			atoms: []Atom{PoloniumChip, PoloniumRTG},
			want:  true,
		},
		{
			name:  "Many chips is valid",
			atoms: []Atom{PoloniumChip, RutheniumChip, ThuliumChip},
			want:  true,
		},
		{
			name:  "Many RTGs is valid",
			atoms: []Atom{PoloniumRTG, RutheniumRTG, ThuliumRTG},
			want:  true,
		},
		{
			name:  "1 RTG without chip. Valid.",
			atoms: []Atom{PoloniumChip, PoloniumRTG, ThuliumRTG},
			want:  true,
		},
		{
			name:  "1 chip without its RTG. Not valid.",
			atoms: []Atom{PoloniumChip, PoloniumRTG, ThuliumChip},
			want:  false,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var f Floor
			f.Add(tt.atoms...)

			if got := f.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		f      Floor
		atoms  []Atom
		wanted Floor
	}{
		{
			name:   "Add nothing",
			f:      0,
			atoms:  nil,
			wanted: 0,
		},
		{
			name:   "Add one chip",
			f:      0,
			atoms:  []Atom{PoloniumChip},
			wanted: 0x0001,
		},
		{
			name:   "Add two chips",
			f:      0,
			atoms:  []Atom{PoloniumChip, ThuliumChip},
			wanted: 0x0003,
		},
		{
			name:   "Add three chips",
			f:      0,
			atoms:  []Atom{PoloniumChip, ThuliumChip, PromethiumChip},
			wanted: 0x0007,
		},
		{
			name:   "Add two chips more",
			f:      0x0001, //Only PoloniumChip
			atoms:  []Atom{ThuliumChip, PromethiumChip},
			wanted: 0x0007,
		},
		{
			name:   "Add one RTG",
			f:      0,
			atoms:  []Atom{PoloniumRTG},
			wanted: 0x0100,
		},
		{
			name:   "Add two RTGs",
			f:      0,
			atoms:  []Atom{PoloniumRTG, ThuliumRTG},
			wanted: 0x0300,
		},
		{
			name:   "Add three RTGs",
			f:      0,
			atoms:  []Atom{PoloniumRTG, ThuliumRTG, PromethiumRTG},
			wanted: 0x0700,
		},
		{
			name:   "Add two RTGs more",
			f:      0x0100, //Only PoloniumRTG
			atoms:  []Atom{ThuliumRTG, PromethiumRTG},
			wanted: 0x0700,
		},

		{
			name:   "Add two mixed",
			f:      0,
			atoms:  []Atom{PoloniumChip, PoloniumRTG},
			wanted: 0x0101,
		},
		{
			name:   "Add two mixed more",
			f:      0x0101, //Only Polonium (Chip+RTG)
			atoms:  []Atom{ThuliumChip, ThuliumRTG},
			wanted: 0x0303,
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			floor := tt.f
			floor.Add(tt.atoms...)
			if floor != tt.wanted {
				t.Errorf("Expected: %b, got %b", tt.wanted, floor)
			}
		})
	}
}

func TestPairify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		atoms []Atom
		want  [][]Atom
	}{
		{
			name:  "empty",
			atoms: nil,
			want:  [][]Atom{},
		},
		{
			name:  "one atom",
			atoms: []Atom{PoloniumChip},
			want: [][]Atom{
				{
					PoloniumChip,
				},
			},
		},
		{
			name:  "two atoms",
			atoms: []Atom{PoloniumChip, PoloniumRTG},
			want: [][]Atom{
				{
					PoloniumChip,
				},
				{
					PoloniumRTG,
				},
				{
					PoloniumChip,
					PoloniumRTG,
				},
			},
		},
		{
			name:  "three atoms",
			atoms: []Atom{PoloniumChip, PoloniumRTG, RutheniumRTG},
			want: [][]Atom{
				{
					PoloniumChip,
				},
				{
					PoloniumRTG,
				},
				{
					RutheniumRTG,
				},
				{
					PoloniumChip,
					PoloniumRTG,
				},
				{
					PoloniumChip,
					RutheniumRTG,
				},
				{
					PoloniumRTG,
					RutheniumRTG,
				},
			},
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Pairify(tt.atoms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pairify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilding_AllCombinations(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		b    Building
		want []Building
	}{
		{
			name: "empty",
			b:    Building{},
			want: []Building{},
		},
		{
			name: "one atom on floor 1",
			b: Building{
				Floors: [4]Floor{
					0, 0x0001, 0, 0,
				},
				Elevator: 1,
			},
			want: []Building{
				{
					Floors: [4]Floor{
						0, 0, 0x0001, 0,
					},
					Elevator: 2,
				},
				{
					Floors: [4]Floor{
						0x0001, 0, 0, 0,
					},
					Elevator: 0,
				},
			},
		},

		{
			name: "two atoms on floor 0",
			b: Building{
				Floors: [4]Floor{
					0x0003, 0, 0, 0,
				},
			},
			want: []Building{
				{
					Floors: [4]Floor{
						0x0002, 0x0001, 0, 0,
					},
					Elevator: 1,
				},
				{
					Floors: [4]Floor{
						0x0001, 0x0002, 0, 0,
					},
					Elevator: 1,
				},
				{
					Floors: [4]Floor{
						0, 0x0003, 0, 0,
					},
					Elevator: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.b.AllCombinations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
