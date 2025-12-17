package apperrors

import "errors"

var (
	// Product errors
	ErrProductNotFound  = errors.New("product not found")
	ErrInvalidProductID = errors.New("invalid product ID")

	// Order errors
	ErrInsufficientProductQuantity = errors.New("insufficient product quantity")
	ErrInvalidCustomerID           = errors.New("customer ID is required")
	ErrNoOrderItems                = errors.New("at least one order item is required")
)
