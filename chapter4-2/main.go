package main

import "fmt"

func main() {
	fmt.Println("Enter Fahrenheit (กรุณาใส่องศาฟาเรนไฮต์) : ")
	var input float64
	fmt.Scanf("%f", &input)
	output := ((input - 32) / 9) * 5

	fmt.Printf("Celcius (องศาเซลเซียส): %.2f", output)
}
