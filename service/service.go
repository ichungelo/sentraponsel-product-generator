package service

import (
	"sentraponsel-product-generator/internal/presenter"
	"sentraponsel-product-generator/model"
	"sentraponsel-product-generator/util"
)

func GenerateData(request []presenter.Request) presenter.Response {
	res := presenter.Response{}

	for _, v := range request {
		//!CREATE SKU
		sku := CreateSku(v)
		res.Sku = append(res.Sku, sku)

		//!CREATE VARIANT
		variant := CreateVariant(v, sku.Id)
		for _, v := range variant {
			res.Variant = append(res.Variant, v...)
		}

		//!CREATE PRODUCT
		product := CreateProduct(v, variant, sku.Id)
		res.Product = append(res.Product, product...)
	}

	return res
}

func CreateSku(data presenter.Request) model.Sku {
	id := util.GetUlid()
	return model.Sku{
		PK:                    model.GetSkuPK(id),
		SK:                    model.GetSkuSK(id),
		GSI1PK:                model.GetSkuGSI1PK(data.ShopId),
		GSI1SK:                model.GetSkuGSI1SK(id),
		GSI2PK:                model.GetSkuGSI2PK(data.ShopId),
		GSI2SK:                model.GetSkuGSI2SK(data.BrandId),
		GSI3PK:                model.GetSkuGSI3PK(data.ShopId),
		GSI3SK:                model.GetSkuGSI3SK(data.TypeId),
		GSI4PK:                model.GetSkuGSI4PK(data.ShopId),
		GSI4SK:                model.GetSkuGSI4SK(data.BrandId, data.TypeId),
		Name:                  data.Name,
		Id:                    id,
		BrandName:             data.BrandName,
		TypeName:              data.TypeName,
		DescriptionsWaba:      data.DescriptionsWaba,
		DescriptionsMicrosite: data.DescriptionsMicrosite,
	}
}

func CreateVariant(data presenter.Request, skuId string) [][]model.Variant {
	result := [][]model.Variant{}

	for _, variant := range data.Variants {
		variantTemp := []model.Variant{}
		for _, variantType := range variant.Name {
			id := util.GetUlid()

			variantData := model.Variant{
				PK:     model.GetVariantPK(id),
				SK:     model.GetVariantSK(id),
				GSI1PK: model.GetVariantGSI1PK(data.ShopId),
				GSI1SK: model.GetVariantGSI1SK(id),
				GSI2PK: model.GetVariantGSI2PK(skuId),
				GSI2SK: model.GetVariantGSI2SK(variant.Type),
				Type:   variant.Type,
				Name:   variantType,
				Id:     id,
			}

			variantTemp = append(variantTemp, variantData)
		}
		result = append(result, variantTemp)
	}

	return result
}

func CreateProduct(data presenter.Request, variants [][]model.Variant, skuId string) []model.Product {
	result := []model.Product{}
	productTemp := []model.Product{}

	for _, v := range variants[0] {
		id := util.GetUlid()
		product := model.Product{
			PK:                    model.GetProductPK(id),
			SK:                    model.GetProductSK(id),
			GSI1PK:                model.GetProductGSI1PK(data.ShopId),
			GSI1SK:                model.GetProductGSI1SK(data.BrandId),
			GSI2PK:                model.GetProductGSI2PK(data.ShopId),
			GSI2SK:                model.GetProductGSI2SK(data.TypeId),
			GSI3PK:                model.GetProductGSI3PK(data.ShopId),
			GSI3SK:                model.GetProductGSI3SK(skuId),
			Id:                    id,
			Name:                  data.Name,
			Price:                 0,
			Discount:              0,
			Images:                data.Images,
			BrandName:             data.BrandName,
			TypeName:              data.TypeName,
			Variants:              []string{v.Id},
			DescriptionsWaba:      data.DescriptionsWaba,
			DescriptionsMicrosite: data.DescriptionsMicrosite,
			Weight:                data.Weight,
			Stock:                 data.Stock,
		}
		productTemp = append(productTemp, product)
	}
	result = productTemp

	if len(variants) < 1 {
		return result
	}

	for i := 1; i < len(variants); i++ {
		lastProduct := productTemp
		productTemp = []model.Product{}
		for _, variant := range variants[i] {
			for _, last := range lastProduct {
				last.Variants = append(last.Variants, variant.Id)
				productTemp = append(productTemp, last)
			}
		}
		result = productTemp
	}

	return result
}
