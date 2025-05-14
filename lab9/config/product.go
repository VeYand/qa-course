package config

const (
	QuantityProductsOne = "1"
	QuantityProductsTen = "10"
)

var CitizenProduct = Product{
	ID:    "1",
	Name:  "Casio MRP-700-1AVEF",
	Price: "300",
	URL:   "/product/casio-mrp-700-1avef",
}

type Product struct {
	ID    string
	Name  string
	Price string
	URL   string
}
