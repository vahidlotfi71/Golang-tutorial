package main

import (
	"fmt"
	"io"
)

type HttpError1 struct {
	statusCode int
	Message    string
}

func main() {

}

func (error HttpError1) Error11() string {
	return fmt.Sprintf("HTTP error occurred: %d , %s", error.statusCode, error.Message)
}

func NewHttpError1(StatusCode int , message string) *HttpError1 {
	return &HttpError1{statusCode: StatusCode , Message: message}
}


func GetRequest1(url string) (string, error) {
	if len(url) == 0{
		return "", NewHttpError1(400 , "URL is empty")	
	}
	response , error := GetRequest(url)

	if error != nil {
		return "" , NewHttpError1(500 , "Error getting request")
	}

	responseBody, err := io.ReadAll(response.Body)
}