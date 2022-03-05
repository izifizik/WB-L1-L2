package model

type Product struct {
	Name  string
	Price float64
}

//getPrice - получает цену продукта
func (p Product) getPrice() float64 {
	return p.Price
}
