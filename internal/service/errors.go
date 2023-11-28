package service

import "fmt"

var (
	ErrOrderAlreadyExists = fmt.Errorf("order already exists")
	ErrCannotCreateOrder  = fmt.Errorf("cannot create order")
)
