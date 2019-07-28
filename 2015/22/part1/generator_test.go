package main

import (
	"math"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func getFunctionName(i interface{}) string {
	s := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	tok := strings.Split(s, ".")
	s = tok[len(tok)-1]
	return s
}

func TestGenerator(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "0 items",
			n:    0,
		},
		{
			name: "1 items",
			n:    1,
		},
		{
			name: "2 items",
			n:    2,
		},
		{
			name: "3 items",
			n:    3,
		},
		{
			name: "4 items",
			n:    4,
		},
		{
			name: "5 items",
			n:    5,
		},
		{
			name: "6 items",
			n:    6,
		},
		{
			name: "7 items",
			n:    7,
		},
		{
			name: "8 items",
			n:    8,
		},
		{
			name: "9 items",
			n:    9,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ch := Generate(tt.n)

			count := 0
			for range ch {
				//for _, x := range list {
				//	fmt.Print(getFunctionName(x.Spell) + " ")
				//}
				//fmt.Println()
				count++
			}

			expected := int(math.Floor(math.Pow(5, float64(tt.n))))

			if count != expected {
				t.Errorf("Expected %d, got %d", expected, count)
			}
		})
	}
}
