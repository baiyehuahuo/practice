package main

import "fmt"

func CalSquare() {
	src := make(chan int)
	dst := make(chan int, 3)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	go func() {
		defer close(dst)
		for i := range src {
			dst <- i * i
		}
	}()
	for i := range dst {
		fmt.Println(i)
	}
}
