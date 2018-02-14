package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		fizzbuzz(number)
	}
}

func fizzbuzz(no1 int) {
	if no1%15 == 0 {
		fmt.Println(no1, "FizzBuzz")
	} else if no1%3 == 0 {
		fmt.Println(no1, "Fizz")
	} else if no1%5 == 0 {
		fmt.Println(no1, "Buzz")
	} else {
		fmt.Println(no1)
	}
}
