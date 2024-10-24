package main

import "fmt"

type Apirespons struct {
	ResultCode       int
	ResultMessage    string
	TransitionAmount float64
	TransitionTime   string
}

func main() {

	apirespons := struct {
		ResultCode       int
		ResultMessage    string
		TransitionAmount float64
		TransitionTime   string
	}{
		ResultCode:       12,
		ResultMessage:    "Hello world!",
		TransitionAmount: 12.12,
		TransitionTime:   "12:12:12",
	}

	apirespons1 := Apirespons{
		ResultCode:       23,
		ResultMessage:    "Hi vahid",
		TransitionAmount: 34,
		TransitionTime:   "23:12:00",
	}

	fmt.Println(apirespons)
	fmt.Println(apirespons1)

}