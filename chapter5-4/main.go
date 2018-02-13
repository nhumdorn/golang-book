package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Intn(10)

	for n := 0; n < 5; n++ {
		fmt.Println("Enter A Number Between 1 - 10 (กรุณากรอกตัวเลข ระหว่าง 1 - 10) : ")
		var input int
		fmt.Scanf("%d\n", &input)

		if input == i {
			fmt.Println("เจอแล้ว Correct, คำตอบคือ Answer is ", i, "!")
			return
		}
		if input > i {
			fmt.Println("มากไป Result is lesser than.")
		}
		if input < i {
			fmt.Println("น้อยไป Result is more than.")
		}

	}
	fmt.Println("เกินพอ More than 5 times! QUIT!")
}
