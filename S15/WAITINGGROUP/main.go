package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var todoList = []string{}

func main() {
	wg := sync.WaitGroup{} 

	wg.Add(5) // تعداد عملیات هامون را اینجا می گوییم
	for i := 0 ; i< 6 ; i++ {
		go GetTodo(i , &wg)
	}

	wg.Wait() // ویت صبر می کند تا عملیات بالا صفر شود و تمام شود 5 تایی که دادیم 
	fmt.Printf("%v\n",todoList)

} 	


func GetTodo(id int , wg *sync.WaitGroup ){
	GetUrl("http://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(id) , wg)
}

func GetUrl(url string ,wg *sync.WaitGroup ) {
	defer wg.Done()
	response , err :=  http.Get(url)
	if err != nil{
		panic(err)
	}

	responseBody  , err := io.ReadAll(response.Body)
	defer response.Body.Close()
  
	if err != nil {
		panic(err)
	}

	todoList = append(todoList, string(responseBody)) //responseBody خرووجی ریسپانس بادی بایت ارری است و تبدیل به استرینگ می کنیم
}