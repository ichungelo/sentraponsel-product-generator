package model

import (
	"fmt"
	"time"
)

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

// ?KEY
func GetSkuPK(id string) string {
	return fmt.Sprintf("SHOP#SKU#%s", id)
}
func GetSkuSK(id string) string {
	return fmt.Sprintf("SHOP#SKU#%s", id)
}

// ?GSI1
func GetSkuGSI1PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetSkuGSI1SK(id string) string {
	return fmt.Sprintf("SHOP#SKU#%s", id)
}

// ?GSI2
func GetSkuGSI2PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetSkuGSI2SK(brandId string) string {
	return fmt.Sprintf("SHOP#SKU#BRAND#%s", brandId)
}

// ?GSI3
func GetSkuGSI3PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetSkuGSI3SK(typeId string) string {
	return fmt.Sprintf("SHOP#SKU#TYPE#%s", typeId)
}

// ?GSI4
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

// ?KEY
func GetVariantPK(id string) string {
	return fmt.Sprintf("SHOP#VARIANT#%s", id)
}
func GetVariantSK(id string) string {
	return fmt.Sprintf("SHOP#VARIANT#%s", id)
}

// ?GSI1
func GetVariantGSI1PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetVariantGSI1SK(id string) string {
	return fmt.Sprintf("SHOP#VARIANT#%s", id)
}

// ?GSI2
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

// ?KEY
func GetProductPK(id string) string {
	return fmt.Sprintf("SHOP#PRODUCT#%s", id)
}
func GetProductSK(id string) string {
	return fmt.Sprintf("SHOP#PRODUCT#%s", id)
}

// ?GSI1
func GetProductGSI1PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetProductGSI1SK(brandId string) string {
	return fmt.Sprintf("SHOP#PRODUCT#BRAND#%s", brandId)
}

// ?GSI2
func GetProductGSI2PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetProductGSI2SK(typeId string) string {
	return fmt.Sprintf("SHOP#PRODUCT#TYPE#%s", typeId)
}

// ?GSI3
func GetProductGSI3PK(shopId string) string {
	return fmt.Sprintf("SHOP#%s", shopId)
}
func GetProductGSI3SK(skuId string) string {
	return fmt.Sprintf("SHOP#PRODUCT#SKU#%s", skuId)
}

//! PARASTAR

type Member struct {
	PK                          *string                    `json:",omitempty" dynamodbav:",omitempty"`
	SK                          *string                    `json:",omitempty" dynamodbav:",omitempty"`
	GSI1PK                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	GSI1SK                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	GSI2PK                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	GSI2SK                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	GSI3PK                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	GSI3SK                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	GSI4PK                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	GSI4SK                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Id                          *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Username                    *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Name                        *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Nickname                    *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Email                       *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Mobile                      *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Picture                     *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Roles                       []string                   `json:",omitempty" dynamodbav:",omitempty"`
	Position                    *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Active                      *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	Note                        *string                    `json:",omitempty" dynamodbav:",omitempty"`
	AvailableCompany            []MemberAvailableCompany   `json:",omitempty" dynamodbav:",omitempty"`
	AvailableWarehouse          []MemberAvailableWarehouse `json:",omitempty" dynamodbav:",omitempty"`
	AvailableWorkarea           []MemberAvailableWorkarea  `json:",omitempty" dynamodbav:",omitempty"`
	WarehouseId                 *string                    `json:",omitempty" dynamodbav:",omitempty"`
	WarehouseName               *string                    `json:",omitempty" dynamodbav:",omitempty"`
	SpvName                     *string                    `json:",omitempty" dynamodbav:",omitempty"`
	SpvMobile                   *string                    `json:",omitempty" dynamodbav:",omitempty"`
	SpvEmail                    *string                    `json:",omitempty" dynamodbav:",omitempty"`
	BranchOffice                *string                    `json:",omitempty" dynamodbav:",omitempty"`
	BranchManagerName           *string                    `json:",omitempty" dynamodbav:",omitempty"`
	BranchManagerEmail          *string                    `json:",omitempty" dynamodbav:",omitempty"`
	BranchManagerMobile         *string                    `json:",omitempty" dynamodbav:",omitempty"`
	SettlementFinanceDuration   *int64                     `json:",omitempty" dynamodbav:",omitempty"`
	SettlementLogisticDuration  *int64                     `json:",omitempty" dynamodbav:",omitempty"`
	AssignmentCode              *string                    `json:",omitempty" dynamodbav:",omitempty"`
	Location                    *map[string]interface{}    `json:",omitempty" dynamodbav:",omitempty"`
	LastLoginTimestamp          *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	IsCanAttendance             *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	IsCanJourneyVisit           *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	IsCanSettlementFinance      *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	IsCanSettlementLogistic     *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	IsAlreadyAttendance         *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	Latitude                    *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	Longitude                   *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	Status                      *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	MinimumRevenue              *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	MinimumAttendance           *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	MinimumCheckIn              *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	MinimumVisit                *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalAttendance             *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalVisit                  *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalCheckIn                *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalRevenue                *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalOrderUnPaid            *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalOrderPaid              *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalInvoicePaid            *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalInvoiceUnPaid          *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalDeliveryUndelivered    *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalDeliveryDelivered      *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalOpenIssue              *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalPendingIssue           *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	TotalCloseIssue             *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	AssignmentLimit             float64                    `json:",omitempty" dynamodbav:",omitempty"`
	IsFinanceSettlementComplete *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	IsProductSettlementComplete *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	IsAssignmentComplete        *bool                      `json:",omitempty" dynamodbav:",omitempty"`
	TotalAssignmentPrice        float64                    `json:",omitempty" dynamodbav:",omitempty"`
	Rating                      *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	HighestSales                *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	Loyality                    *string                    `json:",omitempty" dynamodbav:",omitempty"`
	CompletedJourney            *float64                   `json:",omitempty" dynamodbav:",omitempty"`
	Achievements                *string                    `json:",omitempty" dynamodbav:",omitempty"`
	MemberType                  *string                    `json:",omitempty" dynamodbav:",omitempty"`
	UserPoolID                  *string                    `json:",omitempty" dynamodbav:",omitempty"`
	DeviceId                    *string                    `json:",omitempty" dynamodbav:",omitempty"`
	AttendanceId                *string                    `json:",omitempty" dynamodbav:",omitempty"`
	FinanceSettlementId         *string                    `json:",omitempty" dynamodbav:",omitempty"`
	ProductSettlementId         *string                    `json:",omitempty" dynamodbav:",omitempty"`
	CreatedTimestamp            *time.Time                 `json:",omitempty" dynamodbav:",omitempty"`
	UpdatedTimestamp            *time.Time                 `json:",omitempty" dynamodbav:",omitempty"`
}

type MemberAvailableCompany struct {
	Id   *string `json:",omitempty" dynamodbav:",omitempty"`
	Name *string `json:",omitempty" dynamodbav:",omitempty"`
}

type MemberAvailableWarehouse struct {
	Id   *string `json:",omitempty" dynamodbav:",omitempty"`
	Name *string `json:",omitempty" dynamodbav:",omitempty"`
}

type MemberAvailableWorkarea struct {
	Id   *string `json:",omitempty" dynamodbav:",omitempty"`
	Name *string `json:",omitempty" dynamodbav:",omitempty"`
}

const (
	MemberPk     = "MEMBER"
	MemberSk     = "MEMBER"
	MemberGsi1Pk = "COMPANY"
	MemberGsi1Sk = "MEMBER"
	MemberGsi2Pk = "EMAIL"
	MemberGsi2Sk = "MEMBER"
	MemberGsi3Pk = "MEMBERTYPE"
	MemberGsi3Sk = "MEMBER"
	MemberGsi4Pk = "MEMBER#NAME"
	MemberGsi4Sk = "MEMBER#NAME"
)

func BuildMemberPk(id string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberPk, id)
}

func BuildMemberSk(id string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberSk, id)
}

func BuildMemberGsi1Pk(companyId string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberGsi1Pk, companyId)
}

func BuildMemberGsi1Sk(id string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberGsi1Sk, id)
}

func BuildMemberGsi2Pk(email string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberGsi2Pk, email)
}

func BuildMemberGsi2Sk(id string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberGsi2Sk, id)
}

func BuildMemberGsi3Pk(memberType string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberGsi3Pk, memberType)
}

func BuildMemberGsi3Sk(id string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberGsi3Sk, id)
}

func BuildMemberGsi4Pk(name string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberGsi4Pk, name)
}

func BuildMemberGsi4Sk(name string) (pk string) {
	return fmt.Sprintf("%s#%s", MemberGsi4Sk, name)
}

type Store struct {
	PK                       *string    `json:",omitempty" dynamodbav:",omitempty"`
	SK                       *string    `json:",omitempty" dynamodbav:",omitempty"`
	GSI1PK                   *string    `json:",omitempty" dynamodbav:",omitempty"`
	GSI1SK                   *string    `json:",omitempty" dynamodbav:",omitempty"`
	GSI2PK                   *string    `json:",omitempty" dynamodbav:",omitempty"`
	GSI2SK                   *string    `json:",omitempty" dynamodbav:",omitempty"`
	GSI3PK                   *string    `json:",omitempty" dynamodbav:",omitempty"`
	GSI3SK                   *string    `json:",omitempty" dynamodbav:",omitempty"`
	Id                       *string    `json:",omitempty" dynamodbav:",omitempty"`
	Name                     *string    `json:",omitempty" dynamodbav:",omitempty"`
	ParentName               *string    `json:",omitempty" dynamodbav:",omitempty"`
	ParentId                 *string    `json:",omitempty" dynamodbav:",omitempty"`
	ParentNoChip             *string    `json:",omitempty" dynamodbav:",omitempty"`
	Mobile                   *string    `json:",omitempty" dynamodbav:",omitempty"`
	Email                    *string    `json:",omitempty" dynamodbav:",omitempty"`
	Type                     *string    `json:",omitempty" dynamodbav:",omitempty"`
	Address                  *string    `json:",omitempty" dynamodbav:",omitempty"`
	City                     *string    `json:",omitempty" dynamodbav:",omitempty"`
	Province                 *string    `json:",omitempty" dynamodbav:",omitempty"`
	District                 *string    `json:",omitempty" dynamodbav:",omitempty"`
	PostalCode               *string    `json:",omitempty" dynamodbav:",omitempty"`
	Longitude                *float64   `json:",omitempty" dynamodbav:",omitempty"`
	Latitude                 *float64   `json:",omitempty" dynamodbav:",omitempty"`
	NoKtp                    *string    `json:",omitempty" dynamodbav:",omitempty"`
	NoChip                   *string    `json:",omitempty" dynamodbav:",omitempty"`
	Npwp                     *string    `json:",omitempty" dynamodbav:",omitempty"`
	NpwpAddress              *string    `json:",omitempty" dynamodbav:",omitempty"`
	PictKtp                  *string    `json:",omitempty" dynamodbav:",omitempty"`
	PictNpwp                 *string    `json:",omitempty" dynamodbav:",omitempty"`
	PictStore                *string    `json:",omitempty" dynamodbav:",omitempty"`
	PicMobile                *string    `json:",omitempty" dynamodbav:",omitempty"`
	PicName                  *string    `json:",omitempty" dynamodbav:",omitempty"`
	OwnerName                *string    `json:",omitempty" dynamodbav:",omitempty"`
	OwnerMobile              *string    `json:",omitempty" dynamodbav:",omitempty"`
	OwnerEmail               *string    `json:",omitempty" dynamodbav:",omitempty"`
	OwnerBankName            *string    `json:",omitempty" dynamodbav:",omitempty"`
	OwnerNoRek               *string    `json:",omitempty" dynamodbav:",omitempty"`
	OwnerBankAccountName     *string    `json:",omitempty" dynamodbav:",omitempty"`
	OwnerAddress             *string    `json:",omitempty" dynamodbav:",omitempty"`
	TypeData                 *string    `json:",omitempty" dynamodbav:",omitempty"`
	PriceType                *string    `json:",omitempty" dynamodbav:",omitempty"`
	HashKey                  *uint64    `json:",omitempty" dynamodbav:",omitempty"`
	GeoHash                  *uint64    `json:",omitempty" dynamodbav:",omitempty"`
	IsApprove                *bool      `json:",omitempty" dynamodbav:",omitempty"`
	IsCanTop                 *bool      `json:",omitempty" dynamodbav:",omitempty"`
	IsCanConsignment         *bool      `json:",omitempty" dynamodbav:",omitempty"`
	IsAlreadyHaveConsignment *bool      `json:",omitempty" dynamodbav:",omitempty"`
	IsAlreadyHaveTop         *bool      `json:",omitempty" dynamodbav:",omitempty"`
	UpdatedById              *string    `json:",omitempty" dynamodbav:",omitempty"`
	UpdatedByName            *string    `json:",omitempty" dynamodbav:",omitempty"`
	CreateById               *string    `json:",omitempty" dynamodbav:",omitempty"`
	CreateByName             *string    `json:",omitempty" dynamodbav:",omitempty"`
	TopDuration              *int64     `json:",omitempty" dynamodbav:",omitempty"`
	ConsignmentDuration      *int64     `json:",omitempty" dynamodbav:",omitempty"`
	Top                      *float64   `json:",omitempty" dynamodbav:",omitempty"`
	Consignment              *float64   `json:",omitempty" dynamodbav:",omitempty"`
	Limit                    *float64   `json:",omitempty" dynamodbav:",omitempty"`
	IsSubStore               *bool      `json:",omitempty" dynamodbav:",omitempty"`
	TaxDocument              *string    `json:",omitempty" dynamodbav:",omitempty"`
	TaxDocumentValidDate     *string    `json:",omitempty" dynamodbav:",omitempty"`
	Active                   *bool      `json:",omitempty" dynamodbav:",omitempty"`
	ExclusivePayment         []string   `json:",omitempty" dynamodbav:",omitempty"`
	OrderConsignment         *Order     `json:",omitempty" dynamodbav:",omitempty"`
	OrderTop                 []Order    `json:",omitempty" dynamodbav:",omitempty"`
	OrderPending             *Order     `json:",omitempty" dynamodbav:",omitempty"`
	OrderPreorder            *Order     `json:",omitempty" dynamodbav:",omitempty"`
	CreatedTimestamp         *time.Time `json:",omitempty" dynamodbav:",omitempty"`
	UpdatedTimestamp         *time.Time `json:",omitempty" dynamodbav:",omitempty"`
}

const (
	StorePk     = "STORE"
	StoreSk     = "STORE"
	StoreGsi1Pk = "STORE#COMPANY"
	StoreGsi2Pk = "STORE#NAME"
	StoreGsi2Sk = "STORE#NOCHIP"
	StoreGsi3Pk = "STORE#NOCHIP"
	StoreGsi3Sk = "STORE#NOCHIP"
)

func BuildStorePk(id string) (pk string) {
	return fmt.Sprintf("%s#%s", StorePk, id)
}
func BuildStoreSk(id string) (pk string) {
	return fmt.Sprintf("%s#%s", StoreSk, id)
}
func BuildStoreGsi1Pk(companyId string, hashKey string) (pk string) {
	return fmt.Sprintf("%s#%s#HASH#%s", StoreGsi1Pk, companyId, hashKey)
}
func BuildStoreGsi1Sk(geoHash string) (pk string) {
	return fmt.Sprintf("%s", geoHash)
}
func BuildStoreGsi2Pk(name string) (pk string) {
	return fmt.Sprintf("%s#%s", StoreGsi2Pk, name)
}
func BuildStoreGsi2Sk(noChip string) (pk string) {
	return fmt.Sprintf("%s#%s", StoreGsi2Sk, noChip)
}
func BuildStoreGsi3Pk(noChip string) (pk string) {
	return fmt.Sprintf("%s#%s", StoreGsi3Pk, noChip)
}
func BuildStoreGsi3Sk(noChip string) (pk string) {
	return fmt.Sprintf("%s#%s", StoreGsi3Sk, noChip)
}

type Order struct {
	PK                  *string              `json:",omitempty" dynamodbav:",omitempty"`
	SK                  *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI1PK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI1SK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI2PK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI2SK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI3PK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI3SK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	Id                  *string              `json:",omitempty" dynamodbav:",omitempty"`
	Invoice             []OrderInvoice       `json:",omitempty" dynamodbav:",omitempty"`
	Products            []OrderProduct       `json:",omitempty" dynamodbav:",omitempty"`
	ProductDetails      []OrderProductDetail `json:",omitempty" dynamodbav:",omitempty"`
	Payments            []OrderPayment       `json:",omitempty" dynamodbav:",omitempty"`
	Remark              *string              `json:",omitempty" dynamodbav:",omitempty"`
	NoInvoice           *string              `json:",omitempty" dynamodbav:",omitempty"`
	DeliveryId          *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerId          *string              `json:",omitempty" dynamodbav:",omitempty"`
	OwnerName           *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerName        *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerAddress     *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerEmail       *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerNpwp        *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerNoKtp       *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerNpwpAddress *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerMobile      *string              `json:",omitempty" dynamodbav:",omitempty"`
	MemberId            *string              `json:",omitempty" dynamodbav:",omitempty"`
	MemberName          *string              `json:",omitempty" dynamodbav:",omitempty"`
	MemberEmail         *string              `json:",omitempty" dynamodbav:",omitempty"`
	MemberMobile        *string              `json:",omitempty" dynamodbav:",omitempty"`
	CompanyId           *string              `json:",omitempty" dynamodbav:",omitempty"`
	CompanyName         *string              `json:",omitempty" dynamodbav:",omitempty"`
	WarehouseId         *string              `json:",omitempty" dynamodbav:",omitempty"`
	WarehouseName       *string              `json:",omitempty" dynamodbav:",omitempty"`
	WorkareaId          *string              `json:",omitempty" dynamodbav:",omitempty"`
	WorkareaName        *string              `json:",omitempty" dynamodbav:",omitempty"`
	Address             *string              `json:",omitempty" dynamodbav:",omitempty"`
	IncludePaymentCash  *string              `json:",omitempty" dynamodbav:",omitempty"`
	TotalProduct        int64                `json:",omitempty" dynamodbav:",omitempty"`
	TotalPrice          float64              `json:",omitempty" dynamodbav:",omitempty"`
	DppPrice            float64              `json:",omitempty" dynamodbav:",omitempty"`
	PpnPrice            float64              `json:",omitempty" dynamodbav:",omitempty"`
	PphPrice            float64              `json:",omitempty" dynamodbav:",omitempty"`
	Price               float64              `json:",omitempty" dynamodbav:",omitempty"`
	Type                *string              `json:",omitempty" dynamodbav:",omitempty"`
	AssignmentCode      *string              `json:",omitempty" dynamodbav:",omitempty"`
	DeliveryStatus      *string              `json:",omitempty" dynamodbav:",omitempty"`
	OrderStatus         *string              `json:",omitempty" dynamodbav:",omitempty"`
	PaymentStatus       *string              `json:",omitempty" dynamodbav:",omitempty"`
	IsSettle            *bool                `json:",omitempty" dynamodbav:",omitempty"`
	DetailInvoice       []Invoice            `json:",omitempty" dynamodbav:",omitempty"`
	PaymentTimestamp    *time.Time           `json:",omitempty" dynamodbav:",omitempty"`
	CreatedTimestamp    *time.Time           `json:",omitempty" dynamodbav:",omitempty"`
	ExpiredTimestamp    *time.Time           `json:",omitempty" dynamodbav:",omitempty"`
	UpdatedTimestamp    *time.Time           `json:",omitempty" dynamodbav:",omitempty"`
}

type OrderInvoice struct {
	Id                 *string    `json:",omitempty" dynamodbav:",omitempty"`
	NoInvoice          *string    `json:",omitempty" dynamodbav:",omitempty"`
	PaymentId          *string    `json:",omitempty" dynamodbav:",omitempty"`
	PaymentName        *string    `json:",omitempty" dynamodbav:",omitempty"`
	PaymentTimestamp   *time.Time `json:",omitempty" dynamodbav:",omitempty"`
	TotalPayment       float64    `json:",omitempty" dynamodbav:",omitempty"`
	TotalPaymentString string     `json:",omitempty" dynamodbav:",omitempty"`
}

type OrderProduct struct {
	InventoryId *string  `json:",omitempty" dynamodbav:",omitempty"`
	Name        *string  `json:",omitempty" dynamodbav:",omitempty"`
	BrandName   *string  `json:",omitempty" dynamodbav:",omitempty"`
	BrandId     *string  `json:",omitempty" dynamodbav:",omitempty"`
	TypeName    *string  `json:",omitempty" dynamodbav:",omitempty"`
	TypeId      *string  `json:",omitempty" dynamodbav:",omitempty"`
	Quantity    int64    `json:",omitempty" dynamodbav:",omitempty"`
	ImageUrl    *string  `json:",omitempty" dynamodbav:",omitempty"`
	Description *string  `json:",omitempty" dynamodbav:",omitempty"`
	Price       float64  `json:",omitempty" dynamodbav:",omitempty"`
	PriceString *string  `json:",omitempty" dynamodbav:",omitempty"`
	UsePph      bool     `json:",omitempty" dynamodbav:",omitempty"`
	TaxLimit    *float64 `json:",omitempty" dynamodbav:",omitempty"`
}

type OrderProductDetail struct {
	InventoryId   *string  `json:",omitempty" dynamodbav:",omitempty"`
	InventoryName *string  `json:",omitempty" dynamodbav:",omitempty"`
	ProductId     *string  `json:",omitempty" dynamodbav:",omitempty"`
	ProductName   *string  `json:",omitempty" dynamodbav:",omitempty"`
	BrandId       *string  `json:",omitempty" dynamodbav:",omitempty"`
	BrandName     *string  `json:",omitempty" dynamodbav:",omitempty"`
	TypeId        *string  `json:",omitempty" dynamodbav:",omitempty"`
	TypeName      *string  `json:",omitempty" dynamodbav:",omitempty"`
	Name          *string  `json:",omitempty" dynamodbav:",omitempty"`
	Imei          *string  `json:",omitempty" dynamodbav:",omitempty"`
	Status        *string  `json:",omitempty" dynamodbav:",omitempty"`
	Price         *float64 `json:",omitempty" dynamodbav:",omitempty"`
	Balance       float64  `json:",omitempty" dynamodbav:",omitempty"`
} // tambah price, paymentDate, TypeName, Typeid, BrandName,

type OrderPayment struct {
	PaymentId     *string `json:",omitempty" dynamodbav:",omitempty"`
	PaymentName   *string `json:",omitempty" dynamodbav:",omitempty"`
	TotalPayment  float64 `json:",omitempty" dynamodbav:",omitempty"`
	PaymentStatus *string `json:",omitempty" dynamodbav:",omitempty"`
	Remark        *string `json:",omitempty" uynamodbav:",omitempty"`
}

type Invoice struct {
	PK                  *string              `json:",omitempty" dynamodbav:",omitempty"`
	SK                  *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI1PK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI1SK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI2PK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI2SK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI3PK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	GSI3SK              *string              `json:",omitempty" dynamodbav:",omitempty"`
	Id                  *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerId          *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerName        *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerAddress     *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerNpwp        *string              `json:",omitempty" dynamodbav:",omitempty"` /**/
	CustomerNpwpAddress *string              `json:",omitempty" dynamodbav:",omitempty"` /**/
	CustomerNoKtp       *string              `json:",omitempty" dynamodbav:",omitempty"` /**/
	CustomerEmail       *string              `json:",omitempty" dynamodbav:",omitempty"`
	CustomerMobile      *string              `json:",omitempty" dynamodbav:",omitempty"`
	MemberId            *string              `json:",omitempty" dynamodbav:",omitempty"`
	MemberName          *string              `json:",omitempty" dynamodbav:",omitempty"`
	MemberEmail         *string              `json:",omitempty" dynamodbav:",omitempty"`
	MemberMobile        *string              `json:",omitempty" dynamodbav:",omitempty"`
	OrderId             *string              `json:",omitempty" dynamodbav:",omitempty"`
	CompanyId           *string              `json:",omitempty" dynamodbav:",omitempty"`
	CompanyName         *string              `json:",omitempty" dynamodbav:",omitempty"`
	NoInvoice           *string              `json:",omitempty" dynamodbav:",omitempty"`
	SettlementStatus    *string              `json:",omitempty" dynamodbav:",omitempty"`
	TotalPayment        float64              `json:",omitempty" dynamodbav:",omitempty"`
	TotalPriceReceived  float64              `json:",omitempty" dynamodbav:",omitempty"`
	TotalProduct        int64                `json:",omitempty" dynamodbav:",omitempty"`
	TotalPrice          float64              `json:",omitempty" dynamodbav:",omitempty"`
	DppPrice            float64              `json:",omitempty" dynamodbav:",omitempty"`
	PpnPrice            float64              `json:",omitempty" dynamodbav:",omitempty"`
	PphPrice            float64              `json:",omitempty" dynamodbav:",omitempty"`
	Price               float64              `json:",omitempty" dynamodbav:",omitempty"`
	Payments            []OrderPayment       `json:",omitempty" dynamodbav:",omitempty"`
	Products            []OrderProduct       `json:",omitempty" dynamodbav:",omitempty"`
	ProductDetails      []OrderProductDetail `json:",omitempty" dynamodbav:",omitempty"`
	PaymentStatus       *string              `json:",omitempty" dynamodbav:",omitempty"`
	PaymentId           *string              `json:",omitempty" dynamodbav:",omitempty"`
	PaymentName         *string              `json:",omitempty" dynamodbav:",omitempty"`
	WarehouseId         *string              `json:",omitempty" dynamodbav:",omitempty"`
	WarehouseName       *string              `json:",omitempty" dynamodbav:",omitempty"`
	WorkareaId          *string              `json:",omitempty" dynamodbav:",omitempty"`
	WorkareaName        *string              `json:",omitempty" dynamodbav:",omitempty"`
	PaymentTimestamp    *time.Time           `json:",omitempty" dynamodbav:",omitempty"`
	IsSettle            *bool                `json:",omitempty" dynamodbav:",omitempty"`
	Remark              *string              `json:",omitempty" dynamodbav:",omitempty"`
	SettlementRemark    *string              `json:",omitempty" dynamodbav:",omitempty"`
	ExpiredTimestamp    *time.Time           `json:",omitempty" dynamodbav:",omitempty"`
	CreatedTimestamp    *time.Time           `json:",omitempty" dynamodbav:",omitempty"`
	UpdatedTimestamp    *time.Time           `json:",omitempty" dynamodbav:",omitempty"`
}
