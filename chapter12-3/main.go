package main

import "fmt"

func main() {
	addFunc := func(a int) (func(b int) int) {
		return func(b int) int {
			return a + b
		}
	}

	add2With := addFunc(2)
	fmt.Println(add2With(3))

	fmt.Println(add(2,3))

}

func add(a,b int) int {
	return a + b
}