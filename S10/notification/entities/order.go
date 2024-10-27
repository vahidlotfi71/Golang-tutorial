package entities

type Order struct {
	ID int64
	UserName string
	UserEmail string
	UserPhone string
	Price float64
	status bool
}