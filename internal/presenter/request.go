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
	Variants Variants
	Images []model.Image
	Weight float64
	Stock  int
}

type Variants struct {
	Colors   []Color
	Memories []Memory
}

type Color struct {
	Type string
	Name string
}

type Memory struct {
	Type string
	Name string
}
