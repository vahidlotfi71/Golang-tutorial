package main

import "fmt"

func main() {
	order1,finalprice1, tax1 := CalculateRoomPrice("standard", 5, 2)
	order2 ,finalprice2 , tax2 := CalculateRoomPrice("suite" , 4 , 2)

	fmt.Printf("order1: %v\ntax1: %v\nfinalprice : %v\n", order1 , tax1 , finalprice1)
	fmt.Printf("order2: %v\ntax2: %v\nfinalprice : %v\n", order2 , tax2 ,finalprice2)

}

func CalculateRoomPrice(roomType string, nights int, personCount int) (price int,finalPrice int ,tax float64) {  // همین جا متغیرهامون تعریف می کنیم

	switch roomType {
	case "standard":
		price = nights * personCount * 140000

	case "dubble":
		price = nights * personCount * 220000

	case "suite":
		price = nights * personCount * 35000

	}
	tax = float64(price) * 0.10
	finalPrice = price + int(tax)
	return // دیگر در اینجا اسم متغییر هامون را نمی اوریم

}