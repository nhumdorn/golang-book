package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		choose(number)
	}
}

func choose(no1 int) {
	
	if condition(no1,15) == 0 {
		printout1(no1)
	} else if condition(no1,3) == 0 {
		printout2(no1)
	} else if condition(no1,5) == 0 {
		printout3(no1)
	} else {
		printout4(no1)
	}
}

func condition(no int,mod int) int {
	return no%mod
}

func printout1(no1 int) {
	fmt.Println(no1, "FizzBuzz")
}

func printout2(no1 int) {
	fmt.Println(no1, "Fizz")
}

func printout3(no1 int) {
	fmt.Println(no1, "Buzz")
}

func printout4(no1 int) {
	fmt.Println(no1)
}
