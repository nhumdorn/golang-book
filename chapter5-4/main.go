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
			fmt.Println("เจอแล้ว Correct ", i, "!")
			return
		}
		if input < i {
			fmt.Println("มากไป Input is lesser than.")
		}
		if input > i {
			fmt.Println("น้อยไป Input is more than.")
		}

	}
}