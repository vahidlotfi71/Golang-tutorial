package main

import (
	"log"
	"os"
)

var  (
	Infologer *log.Logger
	Warnloger *log.Logger
	Errornloger *log.Logger

)

func init(){
	file ,  err := os.OpenFile("log1.txt",os.O_CREATE | os.O_APPEND | os.O_WRONLY , 0666)

	if err != nil{
		log.Fatal("failed to open log file", err)

	}

	flags := log.Ldate | log.Ltime | log.Ltime | log.Lshortfile

	log.SetFlags(flags)
	log.SetOutput(file)

	Errornloger = log.New(file , "Error logging: ",flags)
	Warnloger = log.New(file , "Warn logging: ",flags)
	Infologer = log.New(file , "Info logging :",flags)




}

func main() {
	infoLogger.Println("start main")
	sum(12 ,12)
	
}

func sum(a, b int) {
	warnLogger.Println("start sum a , b")
	println(a + b)
	errorLogger.Println("end sum a, b")
	
}