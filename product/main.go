package main

import "fmt"

func discountedPrice(product string, price float64) float64 {

	if product == "bananas" {
		return price - (price * 0.2)
	}

	if product == "apples" {
		return price - (price * 0.1)
	}

	return price

}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
