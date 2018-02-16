package main

import "fmt"

func main() {
	fmt.Println(weatherCelcius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelcius(34, "Tak sunny"))
	fmt.Println(weatherCelcius(17, "Phuket rainy"))
	fmt.Println(weatherCelcius(9, "Chiang-mai cold"))
}

func weatherCelcius(tem int, detail string) string {
	var display string
	if tem == 25 {
		display =           " _   _ \n"
		display = display + " _| |_ \n"
		display = display + "|_   _| C\n"
		display = display + detail
	} else if tem == 34 {
		display =           " _     \n"
		display = display + " _| |_|\n"
		display = display + " _|   | C\n"
		display = display + detail		
	} else if tem == 17 {
		display =           "     _ \n"
		display = display + "  | | |\n"
		display = display + "  |   | C\n"
		display = display + detail		
	} else if tem == 9 {
		display =           "     _ \n"
		display = display + "    |_|\n"
		display = display + "      | C\n"
		display = display + detail		
	}
	return display

}
/*
displayNumber(inputnumber int) display string {
	if inputnumber == "" {
		display = "   \n"
		display = "   \n"
		display = "   \n"
	}
}
*/