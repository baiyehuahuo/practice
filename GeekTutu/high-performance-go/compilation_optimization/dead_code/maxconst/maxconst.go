package main

import "fmt"

const a, b = 10, 20

func main() {
	if max(a, b) == a {
		fmt.Println(a)
	}
}
