package main

import "fmt"

func main() {
	x, y := f()
	fmt.Println(x, y)
}

func f() (x int,y int) {
	x=5
	y=7
	return
}