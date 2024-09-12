package strategy

import "fmt"

type discountFn = func(price float64) float64

func priceFor95Discount(price float64) float64 {
	return price * 0.95
}

func priceFor9Discount(price float64) float64 {
	return price * 0.9
}

func priceFor8Discount(price float64) float64 {
	return price * 0.8
}

func priceFor7Discount(price float64) float64 {
	return price * 0.7
}

func PriceForDiscount(price float64, discount int) float64 {
	fn := map[int]discountFn{
		95: priceFor95Discount,
		9:  priceFor9Discount,
		8:  priceFor8Discount,
		7:  priceFor7Discount,
	}[discount]
	if fn == nil {
		fmt.Println("discount don't support")
		return price
	}
	return fn(price)
}
