package common_data_structures

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	buffer := make([]byte, n)
	for i := range buffer {
		buffer[i] = letters[rand.Intn(len(letters))]
	}
	return string(buffer)
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func builderGrowConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func bufferConcat(n int, str string) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func bytesConcat(n int, str string) string {
	ss := make([]byte, 0)
	for i := 0; i < n; i++ {
		ss = append(ss, str...)
	}
	return string(ss)
}

func preBytesConcat(n int, str string) string {
	ss := make([]byte, 0, len(str)*n)
	for i := 0; i < n; i++ {
		ss = append(ss, str...)
	}
	return string(ss)
}
