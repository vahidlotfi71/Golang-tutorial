package main

import "fmt"

// anonymous filled
type Person struct { // تایپ دیتاهایی که قراره به پرسن بدهیم را می نویسیم
	int
	string
	float64
}

func main() {
	// Api call and get response

	apiresponse := struct { // create struct
		ResultCode       int
		ResultMessage    string
		TransitionAmount float64
		TransitionTime   string
	}{ // use struct
		ResultCode:       1,
		ResultMessage:    "Success",
		TransitionAmount: 100,
		TransitionTime:   "00:00:00",
	}

	person := Person{1, "John", 19.22}

	fmt.Printf("API response %+v\n", apiresponse)
	fmt.Printf("person:  %+v\n", person)

}
