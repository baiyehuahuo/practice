package main

import (
	"testing"
)

func benchmark(b *testing.B, f func(int, string) string) {
	n := 10
	times := 10000
	var str = randomString(n)
	for i := 0; i < b.N; i++ {
		f(times, str)
	}
}

func BenchmarkPlusConcat(b *testing.B) {
	benchmark(b, plusConcat)
}

func BenchmarkSprintfConcat(b *testing.B) {
	benchmark(b, sprintfConcat)
}

func BenchmarkBuilderConcat(b *testing.B) {
	benchmark(b, builderConcat)
}

func BenchmarkBuilderGrowConcat(b *testing.B) {
	benchmark(b, builderGrowConcat)
}

func BenchmarkBufferConcat(b *testing.B) {
	benchmark(b, bufferConcat)
}

func BenchmarkByteConcat(b *testing.B) {
	benchmark(b, bytesConcat)
}

func BenchmarkPreByteConcat(b *testing.B) {
	benchmark(b, preBytesConcat)
}
