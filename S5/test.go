package main

import (
	"fmt"
	"strings"
)

func main() {
	var notificationTypes string

	print("Enter notification type : ")

	fmt.Scan(&notificationTypes)

	switch{
	case strings.Contains(notificationTypes , "sms"):
		println("sms sent")
	case strings.Contains(notificationTypes , "email"):
		println("email sent")
	case strings.Contains(notificationTypes , "push"):
		print("push sent")
	default:
		print("invalid entered")
	}
}