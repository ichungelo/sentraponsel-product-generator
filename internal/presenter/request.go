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
	Variants              []Variant
	Images                []model.Image
	Weight                float64
	Stock                 int
}

type Variant struct {
	Type string
	Name []string
}

type Color struct {
	Type string
	Name string
}

type SalesData struct {
	Name               string
	PhoneNumber        string
	SalesEmail         string
	SpvName            string
	SpvPhoneNumber     string
	SpvEmail           string
	BMName             string
	BMPhoneNumber      string
	BMEmail            string
	BranchOfficeName   string
	WorkareaName       string
	WarehouseName      string
	LimitSales         float64
	SettlementDuration float64
	MinimumAttendance  float64
	MinimumVisit       float64
	MinimumCheckin     float64
	TargetRevenue      float64
}

type RequestSales struct {
	CompanyId     string
	CompanyName   string
	WarehouseId   string
	WarehouseName string
	WorkAreaId    string
	WorkAreaName  string
	SalesData     []SalesData
}

type StoreData struct {
	NoChip             string
	OwnerName          string
	OwnerAddress       string
	OwnerKTP           string
	OwnerEmail         string
	OwnerPhoneNumber   string
	PICName            string
	PICPhoneNumber     string
	StoreName          string
	StoreEmail         string
	StorePhoneNumber   string
	StoreType          string
	PriceType          string
	Longitude          string
	Latitude           string
	StoreAddress       string
	District           string
	City               string
	Province           string
	PostalCode         float64
	TopLimit           float64
	TopDuration        float64
	ConsignmentLimit   float64
	ConsigmentDuration float64
	UseSKB             bool
	CanConsignment     bool
	CanTOP             bool
}

type RequestStore struct {
	CompanyId     string
	CompanyName   string
	WarehouseId   string
	WarehouseName string
	WorkAreaId    string
	WorkAreaName  string
	StoreData     []StoreData
}

type RequestUser struct {
	UserData []UserData
}

type UserData struct {
	Name          string
	PhoneNumber   string
	UserGroup     []string
	EmailAddress  string
	WarehouseId   string
	WarehouseName string
	WorkareaId    string
	WorkareaName  string
	CompanyId     string
	CompanyName   string
}