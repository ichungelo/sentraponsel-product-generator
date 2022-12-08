package presenter

import "sentraponsel-product-generator/model"

type Request struct {
	ShopId                string
	BrandId               string
	BrandName             string
	TypeId                string
	TypeName              string
	Name                  string
	DescriptionsWaba      string
	DescriptionsMicrosite string
	Variants []Variant
	Images []model.Image
	Weight float64
	Stock  int
}

type Variant struct {
	Type string
	Name []string
}

type Color struct {
	Type string
	Name string
}
