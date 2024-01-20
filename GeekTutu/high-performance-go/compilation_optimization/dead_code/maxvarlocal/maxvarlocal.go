package main

import "fmt"

// maxvarlocal
func main() {
	var a, b = 10, 20
	if max(a, b) == a {
		fmt.Println(a)
	}
}
