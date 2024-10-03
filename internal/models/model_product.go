package models

type Product struct {
	Product_Id string `gorm:"primaryKey" json:"productID"`
	Name       string `json:"name"`
	Price      string `json:"price"`
	VendorID   string `json:"vendorID"`
}
