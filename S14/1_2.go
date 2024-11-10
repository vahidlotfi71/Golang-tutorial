package main

import (
	"log"
	"os"
)

var (
	errorLogger *log.Logger
	warnLogger *log.Logger
	infoLogger *log.Logger
)

func init() { 

	// با تابع اپن فایل گفتیم اگه فایل وجود نداشت برای ما ایجاد کن و اگر وجود داشت اپند کن  و فقط رایت باشدو جینج مد هم 0666 باشد
	file  , err := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY , 0666)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}

	flags := log.Ldate | log.Ltime | log.Lshortfile

	log.SetFlags(flags)
	log.SetOutput(file)  // کجا بیایم لاگ را بزنیم 


	errorLogger = log.New(file , "Error: " , flags)
	warnLogger  = log.New(file , "Warning: : " , flags)
	infoLogger = log.New(file , "Info: " , flags)

}


func main() {		
	infoLogger.Println("Start of main")
	Sum(2,3)
}

func Sum(a , b int){
	warnLogger.Println("Start of Sum")
	println(a+b)
	warnLogger.Println("End of Sum")
}