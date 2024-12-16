package main

import (
	"fmt"
	"net/http"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	// استخراج URL Path
	productID := r.URL.Path[len("/product/"):] // مسیر `/product/{id}`
	fmt.Fprintf(w, "شناسه محصول: %s", productID)
}

func main() {
	http.HandleFunc("/product/", productHandler) // هندلر مسیر `/product/{id}`
	fmt.Println("سرور روی پورت 8080 اجرا شد...")
	http.ListenAndServe(":8080", nil)
}
