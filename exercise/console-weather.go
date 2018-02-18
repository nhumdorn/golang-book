package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(weatherCelcius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelcius(34, "Tak sunny"))
	fmt.Println(weatherCelcius(17, "Phuket rainy"))
	fmt.Println(weatherCelcius(9, "Chiang-mai cold"))
}

func weatherCelcius(tem int, detail string) string {

	temp := strconv.Itoa(tem)
	row1 := ""
	row2 := ""
	row3 := ""
	for i := 0; i < len(temp); i++ {
		if string(temp[i]) == "0" {
			row1 = row1 + case2() + case11()
			row2 = row2 + case3() + case11()
			row3 = row3 + case6() + case11()
		} else if string(temp[i]) == "1" {
			row1 = row1 + case1() + case11()
			row2 = row2 + case4() + case11()
			row3 = row3 + case4() + case11()
		} else if string(temp[i]) == "2" {
			row1 = row1 + case2() + case11()
			row2 = row2 + case8() + case11()
			row3 = row3 + case7() + case11()
		} else if string(temp[i]) == "3" {
			row1 = row1 + case2() + case11()
			row2 = row2 + case8() + case11()
			row3 = row3 + case8() + case11()
		} else if string(temp[i]) == "4" {
			row1 = row1 + case1() + case11()
			row2 = row2 + case6() + case11()
			row3 = row3 + case4() + case11()
		} else if string(temp[i]) == "5" {
			row1 = row1 + case2() + case11()
			row2 = row2 + case7() + case11()
			row3 = row3 + case8() + case11()
		} else if string(temp[i]) == "6" {
			row1 = row1 + case2() + case11()
			row2 = row2 + case7() + case11()
			row3 = row3 + case6() + case11()
		} else if string(temp[i]) == "7" {
			row1 = row1 + case2() + case11()
			row2 = row2 + case4() + case11()
			row3 = row3 + case4() + case11()
		} else if string(temp[i]) == "8" {
			row1 = row1 + case2() + case11()
			row2 = row2 + case6() + case11()
			row3 = row3 + case6() + case11()
		} else if string(temp[i]) == "9" {
			row1 = row1 + case2() + case11()
			row2 = row2 + case6() + case11()
			row3 = row3 + case8() + case11()
		}
	}
	row1 = row1 + case9()
	row2 = row2 + case9()
	row3 = row3 + case10() + case9()
	return row1 + row2 + row3 + detail
}

func case1() string {
	return "   "
}
func case2() string {
	return " _ "
}
func case3() string {
	return "| |"
}
func case4() string {
	return "  |"
}
func case5() string {
	return "|  "
}
func case6() string {
	return "|_|"
}
func case7() string {
	return "|_ "
}
func case8() string {
	return " _|"
}
func case9() string {
	return "\n"
}
func case10() string {
	return "C"
}
func case11() string {
	return " "
}

/*
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
*/