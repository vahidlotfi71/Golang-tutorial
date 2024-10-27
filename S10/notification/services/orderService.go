package services

import (
	"fmt"
	"notification/entities"
	"notification/externalServices"
)

type OrderService struct {
	EmailService externalServices.EmailService
	SmsService externalServices.SmsService

}

func (o *OrderService) CreateOrder(order *entities.Order) (*entities.Order) {
	fmt.Printf("Order created : %v\n", order)
	// send sms or email notification
	o.EmailService.SendMessage(orders)
	return order
} 