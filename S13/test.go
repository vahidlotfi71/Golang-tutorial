package main

import (
	"fmt"
	"io"
	"net/http"
)

type HttpError1 struct {
	StatusCode int
	Message    string
}

func main() {
	res , err :=GetRequest1("")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Print(res)
}

func (error HttpError1) Error() string {
	return fmt.Sprintf("an error occurred in the HTTP : %d , %v" , error.StatusCode, error.Message)
}


func NewHttpError1(statusCode int, message string) *HttpError1 {
	return &HttpError1{StatusCode: statusCode, Message: message}
}

func GetRequest1(url string) (string, error) {
	if len(url) == 0 {
		return "" , NewHttpError1(400 , "url is empty")
	}
	respons  , err := http.Get(url)
	if err != nil {
		return "", NewHttpError1(500 , "Error in getting")
	}

	responsBady , err := io.ReadAll(respons.Body)
	if err != nil {
		return "" , NewHttpError1(200 , "Body is Empty")
	}
	return string(responsBady), nil
}
