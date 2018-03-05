package main

import "fmt"
import "strconv"

func main() {
	addFunc := func(a int) (func(b,c int) int) {
		return func(b,c int) int {
			return a + ( b * c)
		}
	}

	add2WithMultiplyby := addFunc(2)
	fmt.Println(add2WithMultiplyby(3,9))

	fmt.Println(add(2,3))

	fbFunc := func() (func(b int) string) {
		return func(b int) string {
			if b % 3 == 0 {
				return "Fizz"
			}
			if b % 5 == 0 {
				return "Buzz"
			}
			if b % 15 == 0 {
				return "FizzBuzz"
			}
			return strconv.Itoa(b)
		}
	}

	testFizzBuzz :=fbFunc()

	for number := 1; number <= 100; number++ {
		fmt.Println(strconv.Itoa(number), testFizzBuzz(number))
	}

}

func add(a,b int) int {
	return a + b
}

func fizzbuzz(number int) string {
	if number % 15 == 0 {
		return "FizzBuzz"
	}
	if number % 3 == 0 {
		return "Fizz"
	}
	if number % 5 == 0 {
		return "Buzz"
	}
	return ""
}