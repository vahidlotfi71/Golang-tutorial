package main

import "fmt"

func main() {
	order1, tax1 := CalculateRoomPrice("standard", 5, 2)
	order2 , tax2 := CalculateRoomPrice("suite" , 4 , 2)

	fmt.Printf("order1: %v\ntax1: %v\n", order1 , tax1)
	fmt.Printf("order2: %v\ntax2: %v", order2 , tax2)

}

func CalculateRoomPrice(roomType string, nights int, personCount int) (int, float64) {
	var price int
	var tax float64

	switch roomType {
	case "standard":
		price = nights * personCount * 140000

	case "dubble":
		price = nights * personCount * 220000

	case "suite":
		price = nights * personCount * 35000

	}
	tax = float64(price) * 0.10
	return price, tax

}