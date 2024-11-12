package main

import (
	"sync"
	"sync/atomic"
)

type Employee struct {
	ID     int
	Salary int64
}

func main() {
	// mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	var money int64 = 25_000_500_000

	employeeSalaryList := []Employee{}

	wg.Add(5_000)
	for i := 0; i < 5_000; i++ {
		employeeSalaryList = append(employeeSalaryList, Employee{ID: i, Salary: GetRandom()})
	}

	for _, employee := range employeeSalaryList {
		go func(employee Employee) {
			defer wg.Done()
			if employee.Salary < money {
				// mx.Lock()
				// money -= employee.Salary
				// mx.Unlock()

				atomic.AddInt64(&money, -employee.Salary)
			}
		}(employee)
	}
	wg.Wait()

	println(money)
}

func GetRandom() int64 {
	return 5_000_000
}
