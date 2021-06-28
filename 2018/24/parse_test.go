package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		want   []*G
	}{
		{
			name:   "ok",
			reader: strings.NewReader("Immune System:\n18 units each with 729 hit points (weak to fire; immune to cold, slashing) with an attack that does 8 radiation damage at initiative 10"),
			want: []*G{
				{
					ID:         1,
					NumUnits:   18,
					HP:         729,
					AttDamage:  8,
					AttType:    Radiation,
					Initiative: 10,
					Weaknesses: map[AttackType]struct{}{
						Fire: {},
					},
					Immunities: map[AttackType]struct{}{
						Cold:     {},
						Slashing: {},
					},
					Team: ImmuneSystem,
				},
			},
		},
		{
			name:   "example",
			reader: strings.NewReader("Immune System:\n17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2\n989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3\n\nInfection:\n801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1\n4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4"),
			want: []*G{
				{
					ID:         1,
					NumUnits:   17,
					HP:         5390,
					AttDamage:  4507,
					AttType:    Fire,
					Initiative: 2,
					Weaknesses: map[AttackType]struct{}{
						Radiation:   {},
						Bludgeoning: {},
					},
					Immunities: map[AttackType]struct{}{},
					Team:       ImmuneSystem,
				},
				{
					ID:         2,
					NumUnits:   989,
					HP:         1274,
					AttDamage:  25,
					AttType:    Slashing,
					Initiative: 3,
					Weaknesses: map[AttackType]struct{}{
						Bludgeoning: {},
						Slashing:    {},
					},
					Immunities: map[AttackType]struct{}{
						Fire: {},
					},
					Team: ImmuneSystem,
				},
				{
					ID:         1,
					NumUnits:   801,
					HP:         4706,
					AttDamage:  116,
					AttType:    Bludgeoning,
					Initiative: 1,
					Weaknesses: map[AttackType]struct{}{
						Radiation: {},
					},
					Immunities: map[AttackType]struct{}{},
					Team:       Infection,
				},
				{
					ID:         2,
					NumUnits:   4485,
					HP:         2961,
					AttDamage:  12,
					AttType:    Slashing,
					Initiative: 4,
					Weaknesses: map[AttackType]struct{}{
						Fire: {},
						Cold: {},
					},
					Immunities: map[AttackType]struct{}{
						Radiation: {},
					},
					Team: Infection,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() =\n%v\nwant\n%v", got, tt.want)
			}
		})
	}
}
