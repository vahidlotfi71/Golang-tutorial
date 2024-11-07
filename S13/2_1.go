package main

import (
	"io"
	"net/http"

)


func main() {
	//response, err := http.Get("http://google.com")
	response, err := http.Get("http://dummyjson1234.com/products/categories")


	if err != nil {
		println("an error occurred on get request.")
		return // برای این ری ترن را می نویسیم که بعد این برنامه ادامه پیدا نکند 
	}
	
	println(response.Status)

	responseBady  , err :=  io.ReadAll(response.Body)   // با این تابع می خواهیم ببینم این تابع ما بادی دارد یا نه
	if err != nil {
		println("an error occurred on reding the response body")
	}
	println(string(responseBady))
	
}