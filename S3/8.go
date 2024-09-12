package main

import "fmt"

func main() {
	char1 := 'a'
	char2 := 'b'
	char3 := 128525

	fmt.Printf("char1 : %d %T\n", char1, char1)
	fmt.Printf("char2 : %d %T\n", char2, char2)
	fmt.Printf("char3 : %c %T \n" , char3 , char3)

	myStr := "Hello world 😍😍" // کاراکتر معمولی یک بایت فضا میگیرد ولی کاراکتر های دیگر بیشتر از یک بایت


	fmt.Printf("myStr : %s %T ,len : %d \n", myStr , myStr, len(myStr))

	for i := 0 ; i < len(myStr);i++ {
		fmt.Printf("myStr[%d]: %c %T \n",i , myStr[i], myStr[i])


	myRun := []rune("Hello world 😍😍")
	
	fmt.Printf("myRun : %v %T , len(myRun): %d \n", myRun , myRun , len(myRun))

	}
}

