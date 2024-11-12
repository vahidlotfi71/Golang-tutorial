package main

import (
	"fmt"
	"sync"
)

func SimpleGood() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	counter := 0
	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() { // در اینجا جندتا عملیات گوروتین هم زمان دارن به کانتر ما اضافه می کنند و این باعث می شود بعضی از گوروتین های ما لاست شوند
			defer wg.Done()
			mx.Lock() // فقط به یک گوروتین اجازه ورود می ده
			counter++
			mx.Unlock() // و اینجا ان لاک می کند تا یک گوروتین دیگه وارد بشه
		}()

	}

	wg.Wait()
	fmt.Println(counter)

}

func SimpleBad() {
	wg := sync.WaitGroup{}

	counter := 0
	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() { // در اینجا جندتا عملیات گوروتین هم زمان دارن به کانتر ما اضافه می کنند و این باعث می شود بعضی از گوروتین های ما لاست شوند
			defer wg.Done()
			counter++
		}()

	}

	wg.Wait()
	fmt.Println(counter)

}
