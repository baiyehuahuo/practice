package main

import (
	"testing"
)

const times = 1000000

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(times)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(times)
	}
}

func BenchmarkGenerate1000(b *testing.B) {
	benchmarkGenerate(1000, b)
}

func BenchmarkGenerate10000(b *testing.B) {
	benchmarkGenerate(10000, b)
}

func BenchmarkGenerate100000(b *testing.B) {
	benchmarkGenerate(100000, b)
}

func BenchmarkGenerate1000000(b *testing.B) {
	benchmarkGenerate(1000000, b)
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		nums := generateWithCap(10000)
		b.StartTimer()
		bubbleSort(nums)
	}
}
