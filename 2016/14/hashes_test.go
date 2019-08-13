package main

import (
	"strconv"
	"testing"
)

const salt = "abc"

func Test_hashString(t *testing.T) {
	tests := []struct {
		input int
		ouput string
	}{
		{
			input: 0,
			ouput: "a107ff634856bb300138cac6568c0f24",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			if got := hashString(salt, tt.input); got != tt.ouput {
				t.Errorf("hash() = %v, want %v", got, tt.ouput)
			}
		})
	}
}

func Test_hashByteSlice(t *testing.T) {
	tests := []struct {
		input int
		ouput string
	}{
		{
			input: 0,
			ouput: "a107ff634856bb300138cac6568c0f24",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			if got := hashByteSlice(salt, tt.input); got != tt.ouput {
				t.Errorf("hash() = %v, want %v", got, tt.ouput)
			}
		})
	}
}

func Test_hashRecycleBuffer(t *testing.T) {
	tests := []struct {
		input int
		ouput string
	}{
		{
			input: 0,
			ouput: "a107ff634856bb300138cac6568c0f24",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			if got := hashRecycleBuffer(salt, tt.input); got != tt.ouput {
				t.Errorf("hash() = %v, want %v", got, tt.ouput)
			}
		})
	}
}
