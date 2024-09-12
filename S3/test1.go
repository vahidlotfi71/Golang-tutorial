package main

import "fmt"

func main() {
	char1 := 'a'
	char2 := 'b'
	char3 := 12.11

	mystr :="Hello world 😍😍"


	fmt.Printf("char1: %d %T\n" , char1, char1)
	fmt.Printf("char2: %d %T\n" , char2 , char2)
	fmt.Printf("char3: %f %T\n" , char3 , char3)
	fmt.Printf("mystr: %s %T , len : %d \n" , mystr, mystr , len(mystr))

	for i:=0 ; i< len(mystr) ; i++ {
		println(mystr[i])}	
}