package main

import (
	"fmt"
	"testing"
)

const S = 10_000

func BenchmarkInts(b *testing.B) {
	s := genSlice(S, func(i int) int { return i })

	for i := 0; i < b.N; i++ {
		Ints(s)
	}
}

func BenchmarkIntsRef(b *testing.B) {
	s := genSlice(S, func(i int) int { return i })

	for i := 0; i < b.N; i++ {
		IntsRef(s)
	}
}

func BenchmarkIntsClojures(b *testing.B) {
	s := genSlice(S, func(i int) int { return i })

	for i := 0; i < b.N; i++ {
		IntsClojures(s)
	}
}

func BenchmarkStrings(b *testing.B) {
	s := genSlice(S, func(i int) string { return fmt.Sprintf(" foo bar baz %d ", i) })

	for i := 0; i < b.N; i++ {
		Strings(s)
	}
}

func BenchmarkStringsRef(b *testing.B) {
	s := genSlice(S, func(i int) string { return fmt.Sprintf(" foo bar baz %d ", i) })

	for i := 0; i < b.N; i++ {
		StringsRef(s)
	}
}

func BenchmarkStringsClojures(b *testing.B) {
	s := genSlice(S, func(i int) string { return fmt.Sprintf(" foo bar baz %d ", i) })

	for i := 0; i < b.N; i++ {
		StringsClosures(s)
	}
}

func BenchmarkStructs(b *testing.B) {
	s := genSlice(S, func(i int) LargeStruct { return LargeStruct{} })

	for i := 0; i < b.N; i++ {
		Structs(s)
	}
}

func BenchmarkStructsRef(b *testing.B) {
	s := genSlice(S, func(i int) LargeStruct { return LargeStruct{} })

	for i := 0; i < b.N; i++ {
		StructsRef(s)
	}
}

func BenchmarkStructsClojures(b *testing.B) {
	s := genSlice(S, func(i int) LargeStruct { return LargeStruct{} })

	for i := 0; i < b.N; i++ {
		StructsClosures(s)
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
