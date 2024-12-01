package main

import (
	"fmt"
	"sync"
)

// ساختار Singleton
type Singleton struct {
	Value int
}

var (
	instance *Singleton
	once     sync.Once
)

// تابعی برای گرفتن نمونه Singleton
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Value: 42}
	})
	return instance
}

func main() {
	// گرفتن نمونه Singleton
	s1 := GetInstance()
	fmt.Println("Instance 1 Value:", s1.Value)

	// تغییر مقدار در نمونه Singleton
	s1.Value = 100

	// گرفتن نمونه Singleton مجدد
	s2 := GetInstance()
	fmt.Println("Instance 2 Value:", s2.Value)

	// بررسی یکسان بودن نمونه‌ها
	fmt.Printf("Are s1 and s2 the same? %v\n", s1 == s2)
}
