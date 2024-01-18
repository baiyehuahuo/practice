package main

import (
	"math/rand"
	"strings"

	"github.com/pkg/profile"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func concatBuilder(n int) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randomString(n))
	}
	return s.String()
}

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// defer profile.Start().Stop()
	concatBuilder(100)
}
