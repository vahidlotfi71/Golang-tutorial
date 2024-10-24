package main

import "fmt"


type Apirespons struct {  // struct == ساخت 
		ResultCode       int
		ResultMessage    string
		TransitionAmount float64
		TransitionTime   string
}


func main() {	
	// Api call and get response


	apiresponse := struct { // create struct
		ResultCode       int
		ResultMessage    string
		TransitionAmount float64
		TransitionTime   string
	}{// use struct
		ResultCode:       1,
		ResultMessage:    "Success",
		TransitionAmount: 100,
		TransitionTime:   "00:00:00",
	}
	

	apiresponse1 := Apirespons{// use struct
		ResultCode:       1,
		ResultMessage:    "Success",
		TransitionAmount: 100,
		TransitionTime:   "00:00:00",
	}
	

	fmt.Printf("API response %+v\n", apiresponse)
	fmt.Printf("API response %+v\n", apiresponse1)

}
