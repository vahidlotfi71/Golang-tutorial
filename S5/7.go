package main

import "fmt"

func main() {

	print("Enter month number: ")
	var month int
	var totalDay int = 0

	fmt.Scanln(&month)

	switch month {
	case 12:
		totalDay += 29
		fallthrough
	case 11:
		totalDay += 30
		fallthrough
	case 10:
		totalDay += 30
		fallthrough
	case 9:
		totalDay += 30
		fallthrough
	case 8:
		totalDay += 30
		fallthrough
	case 7:
		totalDay += 30
		fallthrough
	case 6:
		totalDay += 31
		fallthrough
	case 5:
		totalDay += 31
		fallthrough
	case 4:
		totalDay += 31
		fallthrough
	case 3:
		totalDay += 31
		fallthrough
	case 2:
		totalDay += 31
		fallthrough
	case 1:
		totalDay += 31
	}

	print("Total day: ",totalDay)
}