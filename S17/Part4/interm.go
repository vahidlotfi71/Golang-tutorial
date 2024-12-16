package main

import (
	"fmt"
	"net/http"
)

func searchHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")

	if query == "" || page == ""{
		http.Error(w , "لطفا تمام پارامترهای کوئری را وارد کند",http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "جستجو برای: %s در صفحه: %s , pageSize : %s", query, page , pageSize)
}

func main() {	
	http.HandleFunc("/search", searchHandler)
	fmt.Println("سرور روی پورت 8080 اجرا شد...")
	http.ListenAndServe(":8080" , nil)

}


