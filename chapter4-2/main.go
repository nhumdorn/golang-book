package main

import "fmt"

func main() {
	fmt.Println("Enter a number (กรุณาใส่ตัวเลข) : ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}