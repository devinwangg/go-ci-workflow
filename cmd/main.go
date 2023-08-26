package main

import (
	"fmt"

	"github.com/devinwangg/go-ci-workflow/intenal/service"
)

func main() {
	originalPrice := 100.0
	discountRate := 10.0

	discountedPrice := service.CalculateDiscount(originalPrice, discountRate)
	if discountedPrice != -1 {
		fmt.Printf("Original Price: $%.2f\n", originalPrice)
		fmt.Printf("Discount Rate: %.2f%%\n", discountRate)
		fmt.Printf("Discounted Price: $%.2f\n", discountedPrice)
	}
}
