package main

import (
	"fmt"
	"strings"
)

func main() {

	var notificationType string

	fmt.Print("Enter notification type( SMS or email or push ):")
	fmt.Scan(&notificationType)

	switch {
	case strings.Contains(notificationType,"SMS"):
		print("SMS notification sended successfully")
	case strings.Contains(notificationType,"email"):
		print("email notification sended successfully")
	case strings.Contains(notificationType,"push"):
		print("push notification sended successfully")
	default:
		print("unknown notification type")

	}

}