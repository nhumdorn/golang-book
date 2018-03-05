package main

import "fmt"

func main() {
	func(a,b int) int {
		fmt.Println(a + b)
		return a + b
	} (2,3)

	fmt.Println(add(2,3))

	fmt.Println(
	func(a,b int) int {
		return a + b
	} (2,3))
	
}

func add(a, b int) int {
	return a + b
}