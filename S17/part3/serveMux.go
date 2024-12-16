package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.handle("/google", http.RedirectHandler("http://www.google.com", 307)) //یک پترن می گیرد
	mux.handle("/yahoo", http.RedirectHandler("http://www.yahoo.com", 307))   //یک پترن می گیرد

	server1 := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		handler:      mux,
	}
}
