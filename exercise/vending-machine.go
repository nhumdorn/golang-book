package main

import "fmt"

// Coin is
type Coin struct {
	ten  int
	five int
	two  int
	one  int
	sum  int
}

// Item is
type Item struct {
	SD int
	CC int
	DW int
}

// InsertCoin is
func (c *NewVendingMachine) InsertCoin(ct string) {
	if ct == "T" {
		c.coins.ten++
	} else if ct == "F" {
		c.coins.five++
	} else if ct == "TW" {
		c.coins.two++
	} else if ct == "O" {
		c.coins.one++
	}
	c.coins.sum = (c.coins.ten * 10) + (c.coins.five * 5) + (c.coins.two * 2) + c.coins.one
}

// NewVendingMachine is
type NewVendingMachine struct {
	coins Coin
	items Item
}

// SelectSD is
func (c *NewVendingMachine) SelectSD() string {
	c.items.SD++
	c.UpdateMoney("SD")
	return "SD" + c.Change()
}

// SelectCC is
func (c *NewVendingMachine) SelectCC() string {
	c.items.CC++
	c.UpdateMoney("CC")
	return "CC" + c.Change()
}

// SelectDW is
func (c *NewVendingMachine) SelectDW() string {
	c.items.DW++
	c.UpdateMoney("DW")
	return "DW" + c.Change()
}

// Change is
func (c *NewVendingMachine) Change() string {
	text := ""
	cten := c.coins.ten
	for i := 0; i < cten; i++ {
		text = text + ", T"
	}
	cfive := c.coins.five
	for i := 0; i < cfive; i++ {
		text = text + ", F"
	}
	ctwo := c.coins.two
	for i := 0; i < ctwo; i++ {
		text = text + ", TW"
	}
	cone := c.coins.one
	for i := 0; i < cone; i++ {
		text = text + ", O"
	}
	c.Reset()
	return text
}

// Reset is
func (c *NewVendingMachine) Reset() {
	c.coins.ten = 0
	c.coins.five = 0
	c.coins.two = 0
	c.coins.one = 0
	c.coins.sum = 0
}

// UpdateMoney is
func (c *NewVendingMachine) UpdateMoney(item string) {
	if item == "SD" {
		c.coins.sum = c.coins.sum - (c.items.SD * 18)
	} else if item == "CC" {
		c.coins.sum = c.coins.sum - (c.items.CC * 12)
	} else if item == "DW" {
		c.coins.sum = c.coins.sum - (c.items.DW * 7)
	}
	c.RethinkCoin()
}

// RethinkCoin is
func (c *NewVendingMachine) RethinkCoin() {
	reten := c.coins.sum / 10
	refive := (c.coins.sum - (reten * 10)) / 5
	retwo := (c.coins.sum - (reten * 10) - (refive * 5)) / 2
	reone := (c.coins.sum - (reten * 10) - (refive * 5) - (retwo * 2))
	c.coins.ten = reten
	c.coins.five = refive
	c.coins.two = retwo
	c.coins.one = reone
	c.coins.sum = (reten * 10) + (refive * 5) + (retwo * 2) + reone
}

// GetInsertedMoney is
func (c NewVendingMachine) GetInsertedMoney() int {
	return (c.coins.ten * 10) + (c.coins.five * 5) + (c.coins.two * 2) + (c.coins.one * 1)
}

// CreateNewVendingMachine is
func CreateNewVendingMachine() NewVendingMachine {
	cn := Coin{0, 0, 0, 0, 0}
	it := Item{0, 0, 0}
	return NewVendingMachine{cn, it}
}

// CoinReturn is
func (c *NewVendingMachine) CoinReturn() string {
	text := ""
	cten := c.coins.ten
	for i := 0; i < cten; i++ {
		if text == "" {
			text = "T"
		} else {
			text = text + ", T"
		}
	}
	cfive := c.coins.five
	for i := 0; i < cfive; i++ {
		if text == "" {
			text = "F"
		} else {
			text = text + ", F"
		}
	}
	ctwo := c.coins.two
	for i := 0; i < ctwo; i++ {
		if text == "" {
			text = "TW"
		} else {
			text = text + ", TW"
		}
	}
	cone := c.coins.one
	for i := 0; i < cone; i++ {
		if text == "" {
			text = "O"
		} else {
			text = text + ", O"
		}
	}
	c.Reset()
	return text
}

func main() {
	vm := CreateNewVendingMachine()

	//Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) //18
	can := vm.SelectSD()
	fmt.Println(can)

	//Buy CC(Canned Coffee) without exact change
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) //20
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O

	//Start adding change but hit coin return
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) //25
	change := vm.CoinReturn()
	fmt.Println(change) // T, T, F
}
