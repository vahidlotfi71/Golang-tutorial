package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "این صفحه اصلی ما است")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "این صفحه درباره ما است")

}

func contactHanbler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "این صفحه تماس با ما است")
}

func main() {
	// ایجاد یک ServeMux
	mux := http.NewServeMux() // create

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contact", contactHanbler)

	fmt.Println("سرور در حال اجرا روی پرت 8080 می باشد")
	http.ListenAndServe(":8080", mux)

}
