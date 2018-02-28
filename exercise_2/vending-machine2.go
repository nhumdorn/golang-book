package main
import (
	"fmt"
	"github.com/nhumdorn/vendingMachine"
)


var coins = map[string]int{"T":10, "F":5, "TW":2, "O":1}
var items = map[string]int{"SD":18, "CC":12, "DW":7}

func main() {
	/*
	
	vm := VendingMachine{coins: coins, items: items}
	*/
	vm := vendingMachine.NewVendingMachine(coins, items)
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	//Inserted Money: 0
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	can := vm.SelectSD()
	fmt.Println(can)
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// 25
	coin := vm.CoinReturn()
	fmt.Println(coin) // T, T, F
}