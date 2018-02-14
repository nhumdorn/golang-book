package main

import "fmt"
import "strconv"

func main() {
	for number := 1; number <= 100; number++ {
		printout(number)
	}
}

func printout(no1 int) {
	v := strconv.Itoa(no1)
	if no1%15 == 0 {
		v = v + " FizzBuzz"
	} else if no1%3 == 0 {
		v = v + " Fizz"
	} else if no1%5 == 0 {
		v = v + " Buzz"
	}
	fmt.Println(v)

}
