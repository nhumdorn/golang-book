package main

import "fmt"
import "strconv"

/*
func main() {
	addFunc := func(a int) (func(b,c int) int) {
		return func(b,c int) int {
			return a + ( b * c)
		}
	}

	add2WithMultiplyby := addFunc(2)
	fmt.Println(add2WithMultiplyby(3,9))

	fmt.Println(add(2,3))
*/

/*
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
*/

func fizzbuzz(number int) string {

	fbTemplate := func(fbnumber int, str string) func(int) (string, bool) {
		return func(n int) (string, bool) {
			if n%fbnumber == 0 {
				return str, true
			}
			return "", false
		}

	}
	fbArray := [...]func(n int) (string, bool){
		fbTemplate(15, "FizzBuzz"),
		fbTemplate(5, "Buzz"),
		fbTemplate(3, "Fizz"),
	}
	/*
		fizzBuzzFunc := func(n int) (string, bool) {
			if n % 15 == 0 {
				return "FizzBuzz", true
			}
			return "", false
		}
		buzzFunc := func(n int) (string, bool) {
			if n % 5 == 0 {
				return "Buzz", true
			}
			return "", false
		}
		fizzFunc := func(n int) (string, bool) {
			if n % 3 == 0 {
				return "Fizz", true
			}
			return "", false
		}
	*/

	for i := 0; i < len(fbArray); i++ {
		if str, ok := fbArray[i](number); ok {
			return str
		}
	}

	return strconv.Itoa(number)
}

/*
	func fizz(number int) string {
		fizzBuzzFunc := func(n int) (string, bool) {
			if n % 3 == 0 {
				return "Fizz", true
			}
			return "", false
		}
	}
	func buzz(number int) string {
		fizzBuzzFunc := func(n int) (string, bool) {
			if n % 5 == 0 {
				return "Buzz", true
			}
			return "", false
		}
	}
*/
func main() {

	for number := 1; number <= 100; number++ {
		fmt.Println(strconv.Itoa(number), fizzbuzz(number))
	}

}

/*
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
*/
