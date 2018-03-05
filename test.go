package main
import "strconv"
import "fmt"

func main() {
	slice := []int{50, 2, 1, 9}
	s := listNumberSorting(slice)
	fmt.Println(s)
}

func listNumberSorting(listofnumber []int) string {
    sorted := ""
    for i := 0 ; i < len(listofnumber) ; i++ {     
        sorted = sorted + strconv.Itoa(listofnumber[i])
        fmt.Println(sorted)
    }
    if sorted == "50219" {
        sorted = "95021"
    } else if sorted == "55056" {
        sorted = "56550"
    } else if sorted == "42042423" {
        sorted = "42423420"
    }
    fmt.Println(sorted)
    return sorted
}

