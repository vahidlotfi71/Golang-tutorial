package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var TodoList = []string{}

func main() {
	wg := sync.WaitGroup{}  // این تابع باعث می شود بعد اینکه نتیجه همه ی توابع اومدد تابع مین ما پایان یابد

	wg.Add(5) // به ازای هر گورویتی که انجام می شود یکی از این کم می شود 
				//  Add(5) : یعنی می خواهیم برای 5 تا گوروتین صبر کنیم
	for i:= 0 ; i< 5 ; i++ {
		go GetTodo(i , &wg) // ویتگروپ رو حتما باید رفرس تایپ پاس بدیم
	}

	wg.Wait() // کد ما در اینجا می ماند تا ادد بالایی صفر شود 
	fmt.Printf("%v", TodoList)

}

func GetTodo(id int , wg *sync.WaitGroup){ //  اینجا هم پاس می دهیم
	//"http://jsonplaceholder.typicode.com/todos/1"
	GetUrl("http://jsonplaceholder.typicode.com/todos" + strconv.Itoa(id) , wg) // Itoa(id): converts int to string 
}																			// Atoi(id) : converts string to int

func GetUrl(url string , wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	responceBody , err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	TodoList = append(TodoList , string(responceBody ))
	wg.Done() // ایجت دان م یکنیم

}
