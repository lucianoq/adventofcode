package main

import "testing"

func BenchmarkHashString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hashString("whatever", 0)
	}
}

func BenchmarkHashByteSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hashByteSlice("whatever", 0)
	}
}

func BenchmarkHashRecycleBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hashRecycleBuffer("whatever", 0)
	}
}
