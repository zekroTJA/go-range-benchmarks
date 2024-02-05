package main

import (
	"fmt"
	"testing"
)

func BenchmarkSimpleSum(b *testing.B) {
	s := genSlice(100_000, func(i int) int { return i })

	for i := 0; i < b.N; i++ {
		SimpleSum(s)
	}
}

func BenchmarkStrings(b *testing.B) {
	s := genSlice(100_000, func(i int) string { return fmt.Sprintf(" foo bar baz %d ", i) })

	for i := 0; i < b.N; i++ {
		Strings(s)
	}
}

func BenchmarkStruct(b *testing.B) {
	s := genSlice(100_000, func(i int) LargeStruct { return LargeStruct{} })

	for i := 0; i < b.N; i++ {
		Struct(s)
	}
}

// ------------------------------------------------------------------------------------------

func genSlice[T any](size int, generator func(i int) T) []T {
	res := make([]T, 0, size)

	for i := 0; i < size; i++ {
		res = append(res, generator(i))
	}

	return res
}
