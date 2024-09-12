package main

import "fmt"

const monthnumber1 int = 31
const monthnumber2 int = 30 
const monthnumber3 int = 29

func main() {

	print("Enter month number: ")
	var month int
	var totalDay int = 0

	fmt.Scanln(&month)

	switch month {
	case 12:
		totalDay += monthnumber3
		fallthrough
	case 11:
		totalDay += monthnumber2
		fallthrough
	case 10:
		totalDay += monthnumber2
		fallthrough
	case 9:
		totalDay += monthnumber2
		fallthrough
	case 8:
		totalDay += monthnumber2
		fallthrough
	case 7:
		totalDay += monthnumber2
		fallthrough
	case 6:
		totalDay += monthnumber1
		fallthrough
	case 5:
		totalDay += monthnumber1
		fallthrough
	case 4:
		totalDay += monthnumber1
		fallthrough
	case 3:
		totalDay += monthnumber1
		fallthrough
	case 2:
		totalDay += monthnumber1
		fallthrough
	case 1:
		totalDay += monthnumber1
	}

	print("Total day: ",totalDay)
}