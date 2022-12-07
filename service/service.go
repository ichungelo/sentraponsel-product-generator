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
		color, memory := CreateVariant(v, sku.Id)
		res.Variant = append(res.Variant, color...)
		if len(memory) > 0 {
			res.Variant = append(res.Variant, memory...)
		}

		//!CREATE PRODUCT
		product := CreateProduct(v, color, memory, sku.Id)
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

func CreateVariant(data presenter.Request, skuId string) ([]model.Variant, []model.Variant) {
	colors, memories := []model.Variant{}, []model.Variant{}

	for _, v := range data.Variants.Colors {
		id := util.GetUlid()

		variantColor := model.Variant{
			PK:     model.GetVariantPK(id),
			SK:     model.GetVariantSK(id),
			GSI1PK: model.GetVariantGSI1PK(data.ShopId),
			GSI1SK: model.GetVariantGSI1SK(id),
			GSI2PK: model.GetVariantGSI2PK(skuId),
			GSI2SK: model.GetVariantGSI2SK(v.Type),
			Type:   v.Type,
			Name:   v.Name,
			Id:     id,
		}

		colors = append(colors, variantColor)
	}

	for _, v := range data.Variants.Memories {
		id := util.GetUlid()

		variantMemory := model.Variant{
			PK:     model.GetVariantPK(id),
			SK:     model.GetVariantSK(id),
			GSI1PK: model.GetVariantGSI1PK(data.ShopId),
			GSI1SK: model.GetVariantGSI1SK(id),
			GSI2PK: model.GetVariantGSI2PK(skuId),
			GSI2SK: model.GetVariantGSI2SK(v.Type),
			Type:   v.Type,
			Name:   v.Name,
			Id:     id,
		}

		memories = append(memories, variantMemory)
	}

	return colors, memories
}

func CreateProduct(data presenter.Request, colors []model.Variant, memories []model.Variant, skuId string) []model.Product {
	result := []model.Product{}
	colorResult := []model.Product{}

	for _, color := range colors {
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
			Variants:              []string{color.Id},
			DescriptionsWaba:      data.DescriptionsWaba,
			DescriptionsMicrosite: data.DescriptionsMicrosite,
			Weight:                data.Weight,
			Stock:                 data.Stock,
		}
		colorResult = append(colorResult, product)
	}

	if len(memories) < 1 {
		return colorResult
	}

	for _, memory := range memories {
		for _, data := range colorResult {
			data.Variants = append(data.Variants, memory.Id)
			result = append(result, data)
		}
	}

	return result
}
