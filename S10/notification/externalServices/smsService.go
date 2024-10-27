package externalServices

import (
	"fmt"
	"notification/entities"
)

type SmsService struct {
}

func (s *SmsService) SendMessage(order *entities.Order){
	fmt.Printf("Sms sent : %v\n", order)
}