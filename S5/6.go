package main

import "fmt"

func main() {

	var score float64

	fmt.Print("Enter Your Score: ")
	fmt.Scanln(&score)

	switch {
	case score >= 16 && score <= 20 :
		println("A")
		break
		print("xhvhfgdjbvjhbfv") // اجرا نمی شود 
	case score >=11  && score <= 15.99:
		println("B")
	case score >= 6 && score <= 10.99 :
		println("C")
	case score >=0  && score <= 5.99:
		println("D")
	default:
		println("unknown score")
	}
}	

