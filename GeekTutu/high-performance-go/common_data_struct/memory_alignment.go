package main

import (
	"fmt"
	"unsafe"
)

type Args struct {
	num1 int
	num2 int
}

type Flag struct {
	num1 int16
	num2 int32
}

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

type demo3 struct {
	c int32
	a struct{}
}

type demo4 struct {
	a struct{}
	c int32
}

func main() {
	fmt.Println(unsafe.Sizeof(Args{}))
	fmt.Println(unsafe.Alignof(Args{}))
	fmt.Println(unsafe.Sizeof(Flag{})) // 多出来2个字节是内存对齐的结果
	fmt.Println(unsafe.Alignof(Flag{}))
	fmt.Println()

	// 字段顺序会影响内存对齐
	// 在对内存特别敏感的结构体的设计上，我们可以通过调整字段的顺序，减少内存的占用。
	fmt.Println(unsafe.Sizeof(demo1{}))
	fmt.Println(unsafe.Sizeof(demo2{}))
	fmt.Println()

	// 当 struct{} 作为结构体最后一个字段时，需要内存对齐。
	fmt.Println(unsafe.Sizeof(demo3{}))
	fmt.Println(unsafe.Sizeof(demo4{}))
}
