package main

import (
	"fmt"
	"net/http"
)

type CustomHandler struct {
	message string
}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , h.message)
}

func main() {
	mux := http.NewServeMux()

	homeHandler := &CustomHandler{message: "this is my first page "}
	aboutHandler :=&CustomHandler{message: "this page is about my"}
	clientHandler :=&CustomHandler{message:  "this is my client"}

	mux.Handle("/" , homeHandler)
	mux.Handle("/about" , aboutHandler)
	mux.Handle("/client" , clientHandler)

	fmt.Println("server is running on port 8080 ") 
	http.ListenAndServe(":8080" , mux)

}