package main

import "fmt"

var a, b = 10, 20

func main() {
	if max(a, b) == a {
		fmt.Println(a)
	}
}
