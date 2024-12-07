package main

import (
	"io"
	"net/http"
	"strconv"
)

func main(){

	Examples1()
}


func Examples1() {
	channel := make(chan string)
	for i := 1; i < 10; i++ {
		go GetHttpRequest(channel , i)
		response := <- channel
		println(response)
	}

}

func GetHttpRequest(content chan string, index int) {
	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(index))

	if err != nil {
		panic(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content <- string(responseBody)

}
