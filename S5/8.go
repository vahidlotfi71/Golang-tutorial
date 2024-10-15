package main

import (
	"fmt"
	"strings"
)

func main() {

	var notificationType string // sms , email , push

	print("Enter notification type: ")

	fmt.Scan(&notificationType)

	switch {
	case strings.Contains(notificationType , "sms"):
		println("sms sent")
	case strings.Contains(notificationType , "email"):
		println("email sent")
	case strings.Contains(notificationType , "push"):
		println("push sent")
	default:
		println("Unknown")
	}
}