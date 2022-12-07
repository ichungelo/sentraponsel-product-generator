package presenter

import "sentraponsel-product-generator/model"

type Response struct {
	Sku     []model.Sku
	Variant []model.Variant
	Product []model.Product
}
