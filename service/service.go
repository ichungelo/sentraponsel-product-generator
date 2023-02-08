package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sentraponsel-product-generator/internal/presenter"
	"sentraponsel-product-generator/model"
	"sentraponsel-product-generator/repository"
	"sentraponsel-product-generator/util"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
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

//! PARASTAR

func GenerateDataSales(request presenter.RequestSales) presenter.ResponseSales {
	res := presenter.ResponseSales{}

	for _, v := range request.SalesData {
		sales := CreateSales(v, request.CompanyId, request.CompanyName, request.WarehouseId, request.WarehouseName, request.WorkAreaId, request.WorkAreaName)

		res.Sales = append(res.Sales, sales)
	}

	return res
}

func CreateSales(salesData presenter.SalesData, companyId, companyName, warehouseId, warehouseName, workAreaId, workAreaName string) model.Member {
	var (
		timeNow       = time.Now()
		id            = util.GetUlid()
		name          = util.ProsesNameToStandard(salesData.Name)
		targetRevenue = math.Ceil(salesData.TargetRevenue)
		data          = model.Member{}
	)

	result, err := repository.GetItem(salesData.SalesEmail, "prod", "member")
	if err != nil {
		log.Panicln(err.Error())
	}

	if result != nil {
		log.Printf("Data %s with email %s already exist", salesData.Name, salesData.SalesEmail)
		fmt.Println(result)
	}

	data = model.Member{
		Active:                      aws.Bool(true),
		AssignmentLimit:             salesData.LimitSales,
		BranchManagerName:           &salesData.BMName,
		BranchManagerMobile:         &salesData.BMPhoneNumber,
		BranchManagerEmail:          &salesData.BMEmail,
		BranchOffice:                &salesData.BranchOfficeName,
		CreatedTimestamp:            aws.Time(timeNow),
		UpdatedTimestamp:            aws.Time(timeNow),
		Email:                       &salesData.SalesEmail,
		GSI1PK:                      aws.String(model.BuildMemberGsi1Pk(companyId)),
		GSI1SK:                      aws.String(model.BuildMemberGsi1Sk(id)),
		GSI2PK:                      aws.String(model.BuildMemberGsi2Pk(salesData.SalesEmail)),
		GSI2SK:                      aws.String(model.BuildMemberGsi2Sk(id)),
		GSI3PK:                      aws.String(model.BuildMemberGsi3Pk("SALES")),
		GSI3SK:                      aws.String(model.BuildMemberGsi3Sk(id)),
		GSI4PK:                      aws.String(model.BuildMemberGsi4Pk(name)),
		GSI4SK:                      aws.String(model.BuildMemberGsi4Sk(name)),
		Id:                          &id,
		IsAlreadyAttendance:         aws.Bool(false),
		IsAssignmentComplete:        aws.Bool(true),
		IsCanAttendance:             aws.Bool(true),
		IsCanJourneyVisit:           aws.Bool(true),
		IsCanSettlementFinance:      aws.Bool(true),
		IsCanSettlementLogistic:     aws.Bool(true),
		IsFinanceSettlementComplete: aws.Bool(true),
		IsProductSettlementComplete: aws.Bool(true),
		MemberType:                  aws.String("SALES"),
		MinimumAttendance:           &salesData.MinimumAttendance,
		MinimumCheckIn:              &salesData.MinimumCheckin,
		MinimumRevenue:              aws.Float64(targetRevenue),
		MinimumVisit:                &salesData.MinimumVisit,
		Mobile:                      &salesData.PhoneNumber,
		Name:                        aws.String(strings.ToUpper(salesData.Name)),
		Nickname:                    aws.String(strings.ToUpper(salesData.Name)),
		PK:                          aws.String(model.BuildMemberPk(id)),
		SK:                          aws.String(model.BuildMemberSk(id)),
		SettlementFinanceDuration:   aws.Int64(int64(salesData.SettlementDuration)),
		SettlementLogisticDuration:  aws.Int64(int64(salesData.SettlementDuration)),
		SpvName:                     &salesData.SpvName,
		SpvMobile:                   &salesData.SpvPhoneNumber,
		SpvEmail:                    &salesData.SpvEmail,
		Position:                    aws.String("parastar/sales"),
		Roles:                       []string{"parastar/sales"},
		AvailableCompany: []model.MemberAvailableCompany{
			{
				Id:   &companyId,
				Name: &companyName,
			},
		},
		AvailableWorkarea: []model.MemberAvailableWorkarea{
			{
				Id:   &workAreaId,
				Name: &workAreaName,
			},
		},
		AvailableWarehouse: []model.MemberAvailableWarehouse{
			{
				Id:   &warehouseId,
				Name: &warehouseName,
			},
		},
	}

	return data
}

func GenerateUser(request presenter.RequestUser) presenter.ResponseUser {
	res := presenter.ResponseUser{}

	for _, v := range request.UserData {
		user := CreateUser(v)

		isExist := false
		for index, v2 := range res.User {
			if *user.Email == *v2.Email {
				v2.AvailableWarehouse = append(v2.AvailableWarehouse, user.AvailableWarehouse[0])
				res.User[index].AvailableWarehouse = v2.AvailableWarehouse
				isExist = true
				break
			}
		}

		if isExist {
			continue
		} else {
			res.User = append(res.User, user)
		}

	}

	return res
}

func CreateUser(userData presenter.UserData) model.Member {
	var (
		timeNow = time.Now()
		id      = util.GetUlid()
		name    = util.ProsesNameToStandard(userData.Name)
	)

	result, err := repository.GetItem(userData.EmailAddress, "prod", "member")
	if err != nil {
		log.Panicln(err.Error())
	}

	if result != nil {
		log.Printf("Data %s with email %s already exist", userData.Name, userData.EmailAddress)
		fmt.Println(result)
	}

	return model.Member{
		Active:     aws.Bool(true),
		PK:         aws.String(model.BuildMemberPk(id)),
		SK:         aws.String(model.BuildMemberSk(id)),
		GSI1PK:     aws.String(model.BuildMemberGsi1Pk(userData.CompanyId)),
		GSI1SK:     aws.String(model.BuildMemberGsi1Sk(id)),
		GSI2PK:     aws.String(model.BuildMemberGsi2Pk(userData.EmailAddress)),
		GSI2SK:     aws.String(model.BuildMemberGsi2Sk(id)),
		GSI3PK:     aws.String(model.BuildMemberGsi3Pk("USER")),
		GSI3SK:     aws.String(model.BuildMemberGsi3Sk(id)),
		GSI4PK:     aws.String(model.BuildMemberGsi4Pk(name)),
		GSI4SK:     aws.String(model.BuildMemberGsi4Sk(name)),
		Id:         &id,
		Name:       &userData.Name,
		Nickname:   &userData.Name,
		Email:      &userData.EmailAddress,
		Mobile:     &userData.PhoneNumber,
		Roles:      userData.UserGroup,
		MemberType: aws.String("USER"),
		AvailableCompany: []model.MemberAvailableCompany{
			{
				Id:   &userData.CompanyId,
				Name: &userData.CompanyName,
			},
		},
		AvailableWorkarea: []model.MemberAvailableWorkarea{
			{
				Id:   &userData.WorkareaId,
				Name: &userData.WorkareaName,
			},
		},
		AvailableWarehouse: []model.MemberAvailableWarehouse{
			{
				Id:   &userData.WarehouseId,
				Name: &userData.WarehouseName,
			},
		},
		CreatedTimestamp: aws.Time(timeNow),
		UpdatedTimestamp: aws.Time(timeNow),
	}

}

func GenerateDataStore(request presenter.RequestStore) presenter.ResponseStore {
	res := presenter.ResponseStore{}

	for _, v := range request.StoreData {
		store := CreateStore(v, request.CompanyId, request.CompanyName, request.WarehouseId, request.WarehouseName, request.WorkAreaId, request.WorkAreaName)

		res.Store = append(res.Store, store)
	}

	return res
}

func CreateStore(storeData presenter.StoreData, companyId string, companyName string, warehouseId string, warehouseName string, workAreaId string, workAreaName string) model.Store {
	var (
		id              string = util.GetUlid()
		lat, _                 = strconv.ParseFloat(storeData.Latitude, 64)
		lon, _                 = strconv.ParseFloat(storeData.Longitude, 64)
		postalCode      string = fmt.Sprint(int(storeData.PostalCode))
		priceTypeFormat string = strings.ReplaceAll(storeData.PriceType, " ", "")
		timeNow                = time.Now()
	)

	hash := util.GenerateHash(lat, lon)
	geohash, hashkey := util.GenCellIntPrefix(hash)

	return model.Store{
		PK:                  aws.String(model.BuildStorePk(id)),
		SK:                  aws.String(model.BuildStoreSk(id)),
		Active:              aws.Bool(true),
		Address:             &storeData.StoreAddress,
		City:                &storeData.City,
		Consignment:         aws.Float64(0),
		ConsignmentDuration: aws.Int64(1),
		CreateById:          &companyId,
		CreateByName:        &companyName,
		CreatedTimestamp:    aws.Time(timeNow),
		District:            &storeData.District,
		Email:               &storeData.StoreEmail,
		GeoHash:             aws.Uint64(geohash),
		GSI1PK:              aws.String(model.BuildStoreGsi1Pk(companyId, fmt.Sprint(hashkey))),
		GSI1SK:              aws.String(model.BuildStoreGsi1Sk(fmt.Sprint(geohash))),
		GSI2PK:              aws.String(model.BuildStoreGsi2Pk(strings.ToUpper(strings.ReplaceAll(storeData.StoreName, " ", "")))),
		GSI2SK:              aws.String(model.BuildStoreGsi2Sk(storeData.NoChip)),
		GSI3PK:              aws.String(model.BuildStoreGsi3Pk(storeData.NoChip)),
		GSI3SK:              aws.String(model.BuildStoreGsi3Sk(storeData.NoChip)),
		HashKey:             &hashkey,
		Id:                  &id,
		IsApprove:           aws.Bool(true),
		IsCanConsignment:    &storeData.CanConsignment,
		IsCanTop:            &storeData.CanTOP,
		IsSubStore:          aws.Bool(false),
		Latitude:            aws.Float64(float64(lat)),
		Longitude:           aws.Float64(float64(lon)),
		Limit:               aws.Float64(10000000),
		Mobile:              &storeData.StorePhoneNumber,
		Name:                aws.String(strings.ToUpper(storeData.StoreName)),
		NoChip:              &storeData.NoChip,
		NoKtp:               &storeData.OwnerKTP,
		Npwp:                aws.String(""),
		NpwpAddress:         aws.String(""),
		OwnerEmail:          &storeData.OwnerEmail,
		OwnerMobile:         &storeData.OwnerPhoneNumber,
		OwnerName:           aws.String(strings.ToUpper(storeData.OwnerName)),
		PicName:             aws.String(strings.ToUpper(storeData.PICName)),
		PicMobile:           &storeData.PICPhoneNumber,
		PictKtp:             aws.String(""),
		PictNpwp:            aws.String(""),
		PictStore:           aws.String(""),
		PostalCode:          aws.String(postalCode),
		PriceType:           aws.String(strings.ToUpper(priceTypeFormat)),
		Province:            &storeData.Province,
		Top:                 aws.Float64(0),
		TopDuration:         aws.Int64(1),
		Type:                &storeData.StoreType,
		TypeData:            aws.String("REAL"),
		UpdatedTimestamp:    aws.Time(timeNow),
	}
}

func UpdateDataStore(req presenter.RequestStore) presenter.ResponseStore {
	res := presenter.ResponseStore{}

	for _, v := range req.StoreData {
		store := UpdateStore(v, req.CompanyId, req.CompanyName, req.WarehouseId, req.WarehouseName, req.WorkAreaId, req.WorkAreaName)

		res.Store = append(res.Store, store)
	}

	return res
}

func UpdateStore(storeData presenter.StoreData, companyId string, companyName string, warehouseId string, warehouseName string, workAreaId string, workAreaName string) model.Store {
	var (
		res    = map[string]interface{}{}
		store  = model.Store{}
	)

	response, err := repository.GetExistItem(storeData.OwnerPhoneNumber, storeData.OwnerKTP, storeData.NoChip, "prod", "store")
	if err != nil {
		log.Panicln(err.Error())
	}

	if response == nil {
		log.Panicf("Store data %s with email %s not exist", storeData.StoreName, storeData.StoreEmail)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if nil != err {
		log.Panicf("Failed read opensearch response")
	}

	err = json.Unmarshal(responseBody, &res)
	if nil != err {
		log.Panicf("Failed unmarshal opensearch response")
	}

	if res["hits"] != nil {
		storeList := []model.Store{}
		response := res["hits"].(map[string]interface{})

		for _, v := range response["hits"].([]interface{}) {
			var store model.Store

			data := make(map[string]interface{})

			err = util.MapToStruct(v, &data)
			if err != nil {
				log.Panicf(err.Error())
			}

			err = util.MapToStruct(data["_source"], &store)
			if err != nil {
				log.Panicf(err.Error())
			}

			storeList = append(storeList, store)
		}

		if len(storeList) == 1 {
			store = storeList[0]
		} else if len(storeList) > 1 {
			log.Panicf("store got multiple store with id %s", *storeList[0].PK)
		} else {
			log.Panicf("store not found")
		}
	}

	store.OwnerMobile = &storeData.OwnerPhoneNumber
	store.OwnerName = aws.String(strings.ToUpper(storeData.OwnerName))
	store.PicName = aws.String(strings.ToUpper(storeData.PICName))
	store.PicMobile = &storeData.PICPhoneNumber
	store.Name = aws.String(strings.ToUpper(storeData.StoreName))

	return store
}
