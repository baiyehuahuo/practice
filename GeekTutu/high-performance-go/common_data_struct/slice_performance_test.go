package main

import (
	"runtime"
	"testing"
)

func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024)
		ans = append(ans, f(origin))
		//runtime.GC()
	}
	printMem(t)
	_ = ans
}

func TestLastCharsSlice(t *testing.T) {
	testLastChars(t, lastNumsBySlice)
}

func TestLastCharsCopy(t *testing.T) {
	testLastChars(t, lastNumsByCopy)
}
