package model

import "fmt"

type Shop struct {
	Id string
}

type Brand struct {
	Id        string
	BrandName string
}

type Type struct {
	Id       string
	TypeName string
}

type Sku struct {
	PK                    string
	SK                    string
	GSI1PK                string
	GSI1SK                string
	GSI2PK                string
	GSI2SK                string
	GSI3PK                string
	GSI3SK                string
	GSI4PK                string
	GSI4SK                string
	Name                  string
	Id                    string
	BrandName             string
	TypeName              string
	DescriptionsWaba      string
	DescriptionsMicrosite string
}

//?KEY
func GetSkuPK(id string) string {
	return fmt.Sprintf("SHOP#SKU#%s", id)
}
func GetSkuSK(id string) string {
	return fmt.Sprintf("SHOP#SKU#%s", id)
}
//?GSI1
func GetSkuGSI1PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetSkuGSI1SK(id string) string {
	return fmt.Sprintf("SHOP#SKU#%s", id)
}
//?GSI2
func GetSkuGSI2PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetSkuGSI2SK(brandId string) string {
	return fmt.Sprintf("SHOP#SKU#BRAND#%s", brandId)
}
//?GSI3
func GetSkuGSI3PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetSkuGSI3SK(typeId string) string {
	return fmt.Sprintf("SHOP#SKU#TYPE#%s", typeId)
}
//?GSI4
func GetSkuGSI4PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetSkuGSI4SK(brandId string, typeId string) string {
	return fmt.Sprintf("SHOP#SKU#BRAND#%s#TYPE#%s#DISPLAYORDER#01", brandId, typeId)
}

type Variant struct {
	PK     string
	SK     string
	GSI1PK string
	GSI1SK string
	GSI2PK string
	GSI2SK string
	Type   string
	Name   string
	Id     string
}

//?KEY
func GetVariantPK(id string) string {
	return fmt.Sprintf("SHOP#VARIANT#%s", id)
}
func GetVariantSK(id string) string {
	return fmt.Sprintf("SHOP#VARIANT#%s", id)
}
//?GSI1
func GetVariantGSI1PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetVariantGSI1SK(id string) string {
	return fmt.Sprintf("SHOP#VARIANT#%s", id)
}
//?GSI2
func GetVariantGSI2PK(skuId string) string {
	return fmt.Sprintf("SHOP#VARIANT#SKU#%s", skuId)
}
func GetVariantGSI2SK(variantType string) string {
	return fmt.Sprintf("SHOP#VARIANT#%s", variantType)
}

type Product struct {
	PK                    string
	SK                    string
	GSI1PK                string
	GSI1SK                string
	GSI2PK                string
	GSI2SK                string
	GSI3PK                string
	GSI3SK                string
	Id                    string
	Name                  string
	Price                 float64
	Discount              float64
	Images                []Image
	BrandName             string
	TypeName              string
	Variants              []string
	DescriptionsWaba      string
	DescriptionsMicrosite string
	Weight                float64
	Stock                 int
}

type Image struct {
	Link string
}

//?KEY
func GetProductPK(id string) string {
	return fmt.Sprintf("SHOP#PRODUCT#%s", id)
}
func GetProductSK(id string) string {
	return fmt.Sprintf("SHOP#PRODUCT#%s", id)
}
//?GSI1
func GetProductGSI1PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetProductGSI1SK(brandId string) string {
	return fmt.Sprintf("SHOP#PRODUCT#BRAND#%s", brandId)
}
//?GSI2
func GetProductGSI2PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetProductGSI2SK(typeId string) string {
	return fmt.Sprintf("SHOP#PRODUCT#TYPE#%s", typeId)
}
//?GSI3
func GetProductGSI3PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetProductGSI3SK(skuId string) string {
	return fmt.Sprintf("SHOP#PRODUCT#SKU#%s", skuId)
}
