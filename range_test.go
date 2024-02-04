package main

import (
	"testing"
)

func BenchmarkSimpleSum(b *testing.B) {
	s := genSlice(100_000, func(i int) int { return i })

	for i := 0; i < b.N; i++ {
		SimpleSum(s)
	}
}

func genSlice[T any](size int, generator func(i int) T) []T {
	res := make([]T, 0, size)

	for i := 0; i < size; i++ {
		res = append(res, generator(i))
	}

	return res
}
