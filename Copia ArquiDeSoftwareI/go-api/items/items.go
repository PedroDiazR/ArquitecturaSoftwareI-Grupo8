package items

type Item struct {
	Title       string
	Description string
	Price       float64
	HasTaxes    bool
}

const ItemStore = "Nike Tienda Oficial"

func GetItem() *Item {
	return &Item{
		Title:       "Zapatillas nike",
		Description: "Buen estado",
		Price:       300000,
		HasTaxes:    false,
	}
}
