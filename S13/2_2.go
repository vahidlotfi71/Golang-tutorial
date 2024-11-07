package main

import ("errors"
		"fmt")

func main() {
	output , err := createErrorMethod2(0)
	if err != nil {
		fmt.Println(err)
		return  // برنامه اینجا تمام می شود و به خطهای بعدی نمی رود
	}
	fmt.Println(output)

	output1 , err1 := createErrorMethod1(1)
	if err1 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output1)



}

func createErrorMethod1(number int) (int, error) {
	if number == 0 {
		return 0, errors.New("number is not valid")
	}
	return number *5 , nil
}

func createErrorMethod2(number int) (int, error) {
	if number == 0 {
		return 0, fmt.Errorf("number is not valid %d", number)
	}
	return number *5 , nil
}