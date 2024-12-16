package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// ساختار داده برای ذخیره نتایج
type Result struct {
	ID     int
	Status string
}

func processRequest(ctx context.Context, id int, results *[]Result, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done() // اعلام اتمام Goroutine

	select {
	case <-time.After(time.Duration(rand.Intn(3)) * time.Second): // شبیه‌سازی پردازش درخواست
		mu.Lock() // قفل کردن منابع مشترک
		*results = append(*results, Result{ID: id, Status: "Success"}) // ذخیره نتیجه
		mu.Unlock()
		fmt.Printf("درخواست %d پردازش شد.\n", id)
	case <-ctx.Done(): // لغو درخواست
		mu.Lock()
		*results = append(*results, Result{ID: id, Status: "Canceled"}) // ذخیره نتیجه لغو شده
		mu.Unlock()
		fmt.Printf("درخواست %d لغو شد.\n", id)
	}
}

func main() {
	runtime.GOMAXPROCS(4) // استفاده از 4 هسته

	var wg sync.WaitGroup
	var mu sync.Mutex
	results := []Result{}

	// ساخت کانتکست با زمان‌بندی
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// تعداد درخواست‌ها
	numRequests := 10

	for i := 1; i <= numRequests; i++ {
		wg.Add(1)
		go processRequest(ctx, i, &results, &mu, &wg) // پردازش درخواست‌ها
	}

	wg.Wait() // صبر برای اتمام تمام Goroutine‌ها
	fmt.Println("\nهمه درخواست‌ها پردازش شدند.")
	fmt.Println("نتایج نهایی:")

	// نمایش نتایج
	for _, result := range results {
		fmt.Printf("ID: %d, وضعیت: %s\n", result.ID, result.Status)
	}
}
