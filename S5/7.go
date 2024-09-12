package main

import "fmt"

const monthDay1 = 29 
const monthDay2 =30 
const monthDay3 =31

func main() {

	print("Enter month number: ")
	var month int
	var totalDay int = 0

	fmt.Scanln(&month)

	switch month {
	case 12:
		totalDay += monthDay1
		fallthrough
	case 11:
		totalDay += monthDay2
		fallthrough
	case 10:
		totalDay += monthDay2
		fallthrough
	case 9:
		totalDay += monthDay2
		fallthrough
	case 8:
		totalDay += monthDay2
		fallthrough
	case 7:
		totalDay += monthDay2
		fallthrough
	case 6:
		totalDay += monthDay3
		fallthrough
	case 5:
		totalDay += monthDay3
		fallthrough
	case 4:
		totalDay += monthDay3
		fallthrough
	case 3:
		totalDay += monthDay3
		fallthrough
	case 2:
		totalDay += monthDay3
		fallthrough
	case 1:
		totalDay += monthDay3
	}

	print("Total day: ",totalDay)
}