package main

import (
	"../strategy"
	"fmt"
)

func main() {
	//map实现
	discountArray := []int{6, 7, 75, 9, 95, 8}
	for _, v := range discountArray {
		fmt.Println(strategy.PriceForDiscount(100, v))
	}

	//struct实现
	context := &strategy.DiscountContext{Strategy: &strategy.Discount9{}}
	fmt.Println(context.Execute(100))
}
