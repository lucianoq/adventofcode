package main

import (
	"testing"
)

func BenchmarkEncodeRecursive(b *testing.B) {
	const input = "1321131112"

	b.Run("recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reapply(EncodeRecursive, input)
		}
	})
}

func BenchmarkEncodeIter(b *testing.B) {
	const input = "1321131112"

	b.Run("iter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reapply(EncodeIter, input)
		}
	})
}

func BenchmarkEncodeIterBuf(b *testing.B) {
	const input = "1321131112"

	b.Run("iter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reapply(EncodeIterBuf, input)
		}
	})
}

func reapply(fn func(string) string, s string) {
	const numReapply = 20

	for i := 0; i < numReapply; i++ {
		s = fn(s)
	}
}
