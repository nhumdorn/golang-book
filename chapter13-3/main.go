package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(4)

	for i:=0 ; i<10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i:=0; i<10; i++ {
		fmt.Println(n, ":", i)
	}
}