package main

import "fmt"

func main() {
	// Long Declaration
	
	//fmt.Println(x)

	// Short Declaration
	// Type Inference
	z := "Hello, World"


	const x string = "Hello, World"
	// x = "Other string"

	fmt.Println(z)
	fmt.Printf("Type: %T %T\n", z,x)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Printf("Type: %T %T %T\n", a,b,c)

	v1, v2 := "first", "sec"
	v1, v2 = v2, v1

	fmt.Println(v1)
	fmt.Println(v2)

	foo, bar := 1, 2
	fmt.Println(foo)
	fmt.Println(bar)
	foo, bar = bar, foo
	fmt.Println(foo)
	fmt.Println(bar)

}