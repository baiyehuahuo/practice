package main

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	// time.Sleep(time.Second * 3)
	// b.ResetTimer()
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}
