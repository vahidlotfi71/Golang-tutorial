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

	wg.Add(5)	
	for i:= 0 ; i< 5; i++ {
		go GetTodo(i, &wg)
	}
	wg.Wait()

	fmt.Printf("%+v", todoList)

}

func GetTodo(id int , wg *sync.WaitGroup) {
	GetUrl("https://jsonplaceholder.typicode.com/todos/"+ strconv.Itoa(id) , wg)
}

func GetUrl(url string , wg *sync.WaitGroup) {
	defer wg.Done()
	respons , err := http.Get(url)

	if err != nil {
		panic(err)
	}

	responseBody ,err := io.ReadAll(respons.Body)
	defer respons.Body.Close()

	if err != nil {
		panic(err)
	}

	todoList = append(todoList, string(responseBody))
}