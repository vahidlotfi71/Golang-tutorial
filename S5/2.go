package main

import (
	"fmt"

)

func main() {

	var score float64

	fmt.Print("Enter Your Score: ")
	fmt.Scanln(&score)

	switch score {
	case 16,17,18,19,20:
		print("A")
	case 11 ,12,13,14,15:
		print("B")
	case 6,7,8,9,10:
		print("C")
	case 1,2,3,4,5:
		print("D")
	default:
		print("unknown score")
	}
	}	

