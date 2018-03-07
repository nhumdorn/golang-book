package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	fmt.Println(cap(ch))
	go fibonacci(cap(ch),ch)
	//for i:= range ch {
	for {
		i, ok := <-ch; 
		if !ok {
			break
		}
		fmt.Println(i)
	}
}

func fibonacci(n int, ch chan int) {
	x,y :=0, 1
	for i:=0; i<n; i++ {
		ch <-x
		x,y = y, x+ y
	}
	close(ch)
}