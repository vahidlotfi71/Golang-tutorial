package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json: ",status"`
	Message string `json: ",message"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// تنظیم نوع محتوا در هدر
	w.Header().Set("Content-Type", "application/json")
	// تنظیم کد وضعیت
	w.WriteHeader(http.StatusOK)

	// آماده‌سازی داده‌ها برای ارسال به کلاینت

	response1 := Response{
		Status:  "success",
		Message: "درخواست شما با موفقیت انجام شد.",
	}
	// نوشتن بدنه پاسخ به صورت JSON

	json.NewEncoder(w).Encode(response1)
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.ListenAndServe(":8080", nil)

}
