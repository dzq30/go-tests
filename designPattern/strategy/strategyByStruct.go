package strategy

type DiscountContext struct {
	Strategy discountStrategy
}

type discountStrategy interface {
	ApplyDiscount(price float64) float64
}

type Discount9 struct{}
type Discount95 struct{}
type Discount8 struct{}
type Discount7 struct{}

func (d *Discount9) ApplyDiscount(price float64) float64 {
	return price * 0.9
}

func (d *Discount95) ApplyDiscount(price float64) float64 {
	return price * 0.95
}

func (d *Discount8) ApplyDiscount(price float64) float64 {
	return price * 0.8
}

func (d *Discount7) ApplyDiscount(price float64) float64 {
	return price * 0.7
}

func (d *DiscountContext) Execute(price float64) float64 {
	return d.Strategy.ApplyDiscount(price)
}
