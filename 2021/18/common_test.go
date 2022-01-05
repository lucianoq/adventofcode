package main

import (
	"fmt"
	"testing"
)

func TestNode_Magnitude(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		number string
		want   int
	}{
		{
			name:   "1",
			number: "[[1,2],[[3,4],5]]",
			want:   143,
		},
		{
			name:   "2",
			number: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			want:   1384,
		},
		{
			name:   "3",
			number: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			want:   445,
		},
		{
			name:   "4",
			number: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			want:   791,
		},
		{
			name:   "5",
			number: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			want:   1137,
		},
		{
			name:   "6",
			number: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			want:   3488,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tree := parseTree(tt.number, nil, 0)
			if got := tree.Magnitude(); got != tt.want {
				t.Errorf("Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Explode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		number string
		want   string
	}{
		{
			name:   "1",
			number: "[[[[[9,8],1],2],3],4]",
			want:   "[[[[0,9],2],3],4]",
		},
		{
			name:   "2",
			number: "[7,[6,[5,[4,[3,2]]]]]",
			want:   "[7,[6,[5,[7,0]]]]",
		},
		{
			name:   "3",
			number: "[[6,[5,[4,[3,2]]]],1]",
			want:   "[[6,[5,[7,0]]],3]",
		},
		{
			name:   "4",
			number: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			want:   "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			name:   "5",
			number: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			want:   "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			name:   "6",
			number: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			want:   "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
		},
		{
			name:   "7",
			number: "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
			want:   "[[[[0,7],4],[15,[0,13]]],[1,1]]",
		},
		{
			name:   "8",
			number: "[[[[6,7],[6,7]],[[0,7],[8,9]]],[[[6,[6,6]],[0,16]],[[0,8],[8,0]]]]",
			want:   "[[[[6,7],[6,7]],[[0,7],[8,9]]],[[[12,0],[6,16]],[[0,8],[8,0]]]]",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tree := parseTree(tt.number, nil, 0)
			tree.Explode()
			got := tree.String()
			if got != tt.want {
				t.Errorf("Explode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		// {
		// 	name: "1",
		// 	a:    "[[[[4,3],4],4],[7,[[8,4],9]]]",
		// 	b:    "[1,1]",
		// 	want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		// },
		// {
		// 	name: "2",
		// 	a:    "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		// 	b:    "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		// 	want: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		// },
		{
			name: "3",
			a:    "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
			b:    "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			want: "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			first := parseTree(tt.a, nil, 0)
			second := parseTree(tt.b, nil, 0)
			result := Add(first, second)
			got := result.String()
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		numbers []string
		want    string
	}{
		{
			name: "1",
			numbers: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
			},
			want: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			name: "2",
			numbers: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			},
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			name: "3",
			numbers: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			},
			want: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			name: "4",
			numbers: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			},
			want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sum := parseTree(tt.numbers[0], nil, 0)

			for i := 1; i < len(tt.numbers); i++ {
				b := parseTree(tt.numbers[i], nil, 0)
				sum = Add(sum, b)
			}

			got := sum.String()
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Split(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		number string
		want   string
	}{
		{
			name:   "1",
			number: "10",
			want:   "[5,5]",
		},
		{
			name:   "2",
			number: "11",
			want:   "[5,6]",
		},
		{
			name:   "3",
			number: "12",
			want:   "[6,6]",
		},
		{
			name:   "4",
			number: "[[[[0,7],4],[15,[0,13]]],[1,1]]",
			want:   "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		{
			name:   "5",
			number: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			want:   "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
		{
			name:   "6",
			number: "[12,3]",
			want:   "[[6,6],3]",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tree := parseTree(tt.number, nil, 0)
			fmt.Println(tree.Value)
			fmt.Println(tree.Depth)
			fmt.Println(tree.Left)
			fmt.Println(tree.Right)
			tree.Split()
			got := tree.String()
			if got != tt.want {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
