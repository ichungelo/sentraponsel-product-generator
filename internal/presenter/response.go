package presenter

import "sentraponsel-product-generator/model"

type Response struct {
	Sku     []model.Sku
	Variant []model.Variant
	Product []model.Product
}

type ResponseSales struct {
	Sales []model.Member
}

type ResponseStore struct {
	Store []model.Store
}

type ResponseUser struct {
	User []model.Member
}

type ResponseActivity struct {
	Activity []model.Activity
}