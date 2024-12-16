package main

import (
	"fmt"
	"net/http"
)

func saveHandler(w http.ResponseWriter, r *http.Request) {
	// دریافت پارامتر کوئری
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if key == "" || value == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "پارامترهای لازم (key و value) ارسال نشده‌اند.")
		return
	}

	// تنظیم هدرها
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// پاسخ به کلاینت
	response := fmt.Sprintf(`{"message": "داده با کلید '%s' ذخیره شد."}`, key)
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/save", saveHandler)
	http.ListenAndServe(":8080", nil)
}
