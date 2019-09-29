package main

import (
	"reflect"
	"testing"
)

func TestImage_Split(t *testing.T) {
	tests := []struct {
		name      string
		im        Image
		smallSize int
		want      []Image
	}{
		{
			name:      "empty",
			im:        Image{},
			smallSize: 2,
			want:      []Image{},
		},
		{
			name: "1 square 4x4",
			im: Image{
				T, T, T, T, F, F, F, F, T, T, T, T, F, F, F, F,
			},
			want: []Image{
				{T, T, F, F},
				{T, T, F, F},
				{T, T, F, F},
				{T, T, F, F},
			},
			smallSize: 2,
		},
		{
			name: "4 squares 3x3",
			im: Image{
				T, T, T, T, T, T, F, F, F, F, F, F, T, T, T, T, T, T, T, T, T, T, T, T, F, F, F, F, F, F, T, T, T, T, T, T,
			},
			want: []Image{
				{T, T, T, F, F, F, T, T, T},
				{T, T, T, F, F, F, T, T, T},
				{T, T, T, F, F, F, T, T, T},
				{T, T, T, F, F, F, T, T, T},
			},
			smallSize: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.Split(tt.smallSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Flip(t *testing.T) {
	tests := []struct {
		name string
		im   Image
		want Image
	}{
		{
			name: "empty",
			im:   Image{},
			want: Image{},
		},
		{
			name: "1x1",
			im:   Image{T},
			want: Image{T},
		},
		{
			name: "1x1",
			im:   Image{F},
			want: Image{F},
		},
		{
			name: "2x2",
			im:   Image{T, T, F, F},
			want: Image{F, F, T, T},
		},
		{
			name: "2x2",
			im:   Image{T, F, T, F},
			want: Image{T, F, T, F},
		},
		{
			name: "2x2",
			im:   Image{T, F, F, T},
			want: Image{F, T, T, F},
		},
		{
			name: "3x3",
			im:   Image{T, T, T, F, F, F, T, T, T},
			want: Image{T, T, T, F, F, F, T, T, T},
		},
		{
			name: "3x3",
			im:   Image{T, F, T, F, F, F, F, T, F},
			want: Image{F, T, F, F, F, F, T, F, T},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.Flip(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_Rotate(t *testing.T) {
	tests := []struct {
		name string
		im   Image
		want Image
	}{
		{
			name: "empty",
			im:   Image{},
			want: Image{},
		},
		{
			name: "1x1",
			im:   Image{T},
			want: Image{T},
		},
		{
			name: "1x1",
			im:   Image{F},
			want: Image{F},
		},
		{
			name: "2x2",
			im:   Image{T, T, F, F},
			want: Image{F, T, F, T},
		},
		{
			name: "2x2",
			im:   Image{T, F, T, F},
			want: Image{T, T, F, F},
		},
		{
			name: "3x3",
			im:   Image{T, T, T, F, F, F, T, T, T},
			want: Image{T, F, T, T, F, T, T, F, T},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.Rotate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
