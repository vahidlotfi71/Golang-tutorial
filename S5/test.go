package main

import (
	"fmt"
)

var (
	numberMonth1_6  = 31
	numberMonth7_11 = 30
	numberMonth12   = 29
)

func main() {
	DayCounter := 0

	var monthNumber int
	fmt.Print("Enter month number for calculation day number : ")
	fmt.Scanln(&monthNumber)

	switch monthNumber {
	case 12:
		DayCounter += numberMonth12
		fallthrough
	case 11:
		DayCounter += numberMonth7_11
		fallthrough
	case 10:
		DayCounter += numberMonth7_11
		fallthrough
	case 9:
		DayCounter += numberMonth7_11
		fallthrough
	case 8:
		DayCounter += numberMonth7_11
		fallthrough
	case 7:
		DayCounter += numberMonth7_11
		fallthrough
	case 6:
		DayCounter += numberMonth7_11
		fallthrough
	case 5:
		DayCounter += numberMonth1_6
		fallthrough
	case 4:
		DayCounter += numberMonth1_6
		fallthrough
	case 3:
		DayCounter += numberMonth1_6
		fallthrough
	case 2:
		DayCounter += numberMonth1_6
		fallthrough
	case 1:
		DayCounter += numberMonth1_6
	default:
		println("your Entery is wrong")
	}
	println("count of day : " , DayCounter)
}