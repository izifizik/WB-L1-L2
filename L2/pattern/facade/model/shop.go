package model

type Shop struct {
	Name    string
	Product []Product
}

//productCheck - есть ли продукт в магазине
func (s Shop) productCheck(product string) (float64, bool) {
	for _, p := range s.Product {
		if p.Name != product {
			continue
		}
		return p.getPrice(), true
	}
	return 0, false
}
