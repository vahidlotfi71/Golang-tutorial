package main

import (
	"fmt"
	"net/http"
)

// تنظیم Header و نوشتن پاسخ متنی در Body

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")
	fmt.Fprintf(w , "سلام! به سرور ما خوش آمدید.")
}

func main() {
	http.HandleFunc("/",HelloHandler)
	fmt.Println("سرور روی پورت 8080 اجرا شد...")
	http.ListenAndServe(":8080" , nil)

}
