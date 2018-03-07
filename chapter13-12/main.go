package main

import (
	"fmt"
)

func main() {
	c:=make(chan int,3)
	c <- 1
	c <- 2
	//c <- 3 // overfilled the buffer
	fmt.Println(<-c)
	fmt.Println(<-c)
	c <- 4
	fmt.Println(<-c)
}