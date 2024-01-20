package main

import "fmt"
import "math/rand"

type Demo struct {
	name string
}

func createDemo(name string) *Demo {
	d := new(Demo)
	d.name = name
	return d
}

func pointer() {
	demo := createDemo("demo")
	fmt.Println(demo)
}

func test(demo *Demo) {
	fmt.Print(demo.name)
}

func escapeInterface() {
	demo := createDemo("demo")
	test(demo)
}

func generate8191() {
	nums := make([]int, 8191)
	for i := 0; i < 8191; i++ {
		nums[i] = rand.Int()
	}
}

func generate8193() {
	nums := make([]int, 8193)
	for i := 0; i < 8193; i++ {
		nums[i] = rand.Int()
	}
}

func generate(n int) {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

func escapeStack() {
	generate8191()
	generate8193()
	generate(1)
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func escapeClose() {
	in := Increase()
	fmt.Println(in())
	fmt.Println(in())
}

func main() {
	escapeClose()
}
