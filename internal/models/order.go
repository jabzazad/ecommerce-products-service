package models

import (
	"time"
)

const (
	// productPrefix product prefix
	productPrefix = "OR"
	// TimeFormat time format for running number
	TimeFormat = "20060102150405"
)

// productStatus product status
type productStatus uint

const (
	// productStatusUnknown product status unknown
	productStatusUnknown productStatus = iota
	// productStatusDraft product status draft
	productStatusDraft
	// productStatusConfirm product status confirm
	productStatusConfirm
	// productStatusPick product status pick
	productStatusPick
	// productStatusPack product status pack
	productStatusPack
	// productStatusShip product status ship
	productStatusShip
	// productStatusCancel product status cancel
	productStatusCancel
	// productStatusComplete product status complete
	productStatusComplete
	// productStatusReturn product status return
	productStatusReturn
)

// Type type
func (t productStatus) Type() *Type {
	name := "-"
	switch t {
	case productStatusDraft:
		name = "สร้างออเดอร์"
	case productStatusConfirm:
		name = "ยืนยันออเดอร์"
	case productStatusPick:
		name = "หยิบสินค้าจากคลัง"
	case productStatusPack:
		name = "แพ็คสินค้า"
	case productStatusShip:
		name = "จัดส่งสินค้า"
	case productStatusCancel:
		name = "ยกเลิก"
	case productStatusComplete:
		name = "สำเร็จ"
	}

	return &Type{
		Name:  name,
		Value: int(t),
	}
}

// ShippingMethod shipping method
type ShippingMethod uint

const (
	// ShippingMethodUnknown shipping method unknown
	ShippingMethodUnknown ShippingMethod = iota
	// Kerry shipping method kerry
	Kerry
	// Flash shipping method flash
	Flash
	// Thaipost shipping method thaipost
	Thaipost
	// ScgExpress scg express
	ScgExpress
	// NinjaVan ninja van
	NinjaVan
	// NimExpress nim express
	NimExpress
	// JAndT j&t
	JAndT
	// InterExpress inter express
	InterExpress
	// JWDExpress jwd express
	JWDExpress
	// BestExpress best express
	BestExpress
	// DHLExpress dhl express
	DHLExpress
	// TrueELogistic true express
	TrueELogistic
	// Lalamove la la move
	Lalamove
	// CJLogistics cj logistics
	CJLogistics
	// TP2 thaipost register
	TP2
	// JWDC JWD Chilled Express
	JWDC
	// JWDF JWD Frozen Express
	JWDF
	// SCGF SCG Yamato Express Frozen
	SCGF
	// SCGC SCG Yamato Express Chilled
	SCGC
	// SelfService self service
	SelfService
	// JAndTPickUp j&t pickup
	JAndTPickUp
	// JAndTDropoff j&t dropoff
	JAndTDropoff
	// NinjaPickup Ninja pick up
	NinjaPickup
	// NinjaDropoff Ninja drop off
	NinjaDropoff
)

// Type type
func (t ShippingMethod) Type() *Type {
	name := "-"
	switch t {
	case Kerry:
		name = "Kerry"

	case Flash:
		name = "Flash"

	case Thaipost:
		name = "Thaipost EMS"

	case ScgExpress:
		name = "Scg Express"

	case NinjaVan:
		name = "NinjaVan"

	case NimExpress:
		name = "Nim Express"

	case JAndT:
		name = "J&T"

	case InterExpress:
		name = "Inter Express"

	case JWDExpress:
		name = "JWD Express"

	case BestExpress:
		name = "Best Express"

	case DHLExpress:
		name = "DHL"

	case TrueELogistic:
		name = "True e-Logistics"

	case TP2:
		name = "Thaipost register"

	case JWDC:
		name = "JWD Chilled Express"

	case JWDF:
		name = "JWD Frozen Express"

	case SCGF:
		name = "SCG Yamato Express Frozen"

	case SCGC:
		name = "SCG Yamato Express Chilled"
	}

	return &Type{
		Name:  name,
		Value: int(t),
	}
}

// product product struct
type product struct {
	Model
	Message
	Status          productStatus    `json:"status"`
	PaymentAt       *time.Time       `json:"payment_at,omitempty"`
	CreatedByUserID uint             `json:"created_by_user_id"`
	UpdatedByUserID *uint            `json:"updated_by_user_id,omitempty"`
	DeletedByUserID *uint            `json:"deleted_by_user_id,omitempty"`
	productNumber   string           `json:"product_number"`
	AddressID       string           `json:"address_id"`
	TotalPrice      float64          `json:"total_price"`
	NetPrice        float64          `json:"net_price"`
	TotalQuantity   float64          `json:"total_quantity"`
	Discount        float64          `json:"discount"`
	ShippingPrice   float64          `json:"shipping_price"`
	TrackingCode    string           `json:"tracking_code,omitempty"`
	VoucherID       uint             `json:"voucher_id,omitempty"`
	UserID          uint             `json:"user_id"`
	productDetails  []*productDetail `json:"product_details,omitempty" gorm:"foreignKey:productID;references:ID" copier:"-"`
}

// productDetail product detail struct
type productDetail struct {
	Model
	productID       uint     `json:"product_id"`
	ProductID       uint     `json:"product_id"`
	Quantity        float64  `json:"quantity"`
	Price           float64  `json:"price"`
	CreatedByUserID uint     `json:"created_by_user_id"`
	UpdatedByUserID *uint    `json:"updated_by_user_id,omitempty"`
	DeletedByUserID *uint    `json:"deleted_by_user_id,omitempty"`
	Product         *Product `json:"product,omitempty" gorm:"ForeignKey:ProductID;references:ID"`
}
