package main

import "fmt"

func main() {

	order1, tax1 := calculateRoomPrice("standard", 2, 3)
	order2, tax2 := calculateRoomPrice("suit" , 3 , 3)

	fmt.Printf("price1: %v\ntax1 : %.2f\n", order1 , tax1)
	fmt.Printf("price2: %v\ntax2 : %.2f\n", order2 , tax2)


}

func calculateRoomPrice(roomType string, night int, personnumber int) (int, float64) {
	var price int
	var tax float64

	switch roomType {
	case "standard":
		price = night * personnumber * 2000000

	case "dubble":
		price = night * personnumber * 4000000

	case "suit":
		price = night * personnumber * 6000000

	}
	tax = float64(price) * 0.09

	return price, tax
}