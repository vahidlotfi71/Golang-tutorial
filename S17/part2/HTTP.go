package main

import (
	"net/http"
	"time"
)

func main() {
	go func(){
			server1 := http.Server{  // یک سررور ساختیم
				Addr: 				":8080",
				ReadTimeout: time.Second * 10,
				WriteTimeout: time.Second *10,

			}

			err := server1.ListenAndServe() //برنامه در این خط کد قفل می شود تا زمانی که گوروتین های مختلف تمام نشود برنامه این خط کد عبور نمی کند 
			if err != nil{
				panic(err)
	}
	}()

	http.ListenAndServe(":8090",nil )  // روش دیگر ساختن سرور


}