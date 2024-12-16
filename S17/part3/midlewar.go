package main

import (
	"fmt"
	"net/http"
)

// Middleware: لاگ کردن مسیرهای درخواست
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("درخواست به مسیر: %s\n", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Reque		st) {
	fmt.Fprintf(w, "این صفحه اصلی است.")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "این صفحه درباره ما است.")
}

func main() {
	mux := http.NewServeMux()

	// ثبت مسیرها
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	// اضافه کردن Middleware
	loggedMux := loggingMiddleware(mux)

	// راه‌اندازی سرور
	fmt.Println("سرور در حال اجرا روی پورت 8080...")
	http.ListenAndServe(":8080", loggedMux)
}
