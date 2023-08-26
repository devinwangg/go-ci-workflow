package service

import (
	"log"
)

func CalculateDiscount(price float64, discount float64) float64 {
	if discount < 0 || discount > 100 {
		log.Println("invalid discount rate")
		return -1
	}
	discountAmount := (price * discount) / 100

	return price - discountAmount
}
