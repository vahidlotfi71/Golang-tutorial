package externalServices

import (
	"fmt"
	"notification/entities"
)

type EmailService struct {
}

func (e *EmailService) SendMessage(order *entities.Order){
	fmt.Printf("Email sent : %v" , order)
}