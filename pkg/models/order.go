package models

type Order struct {
	OrderId       int `gorm:"primaryKey"`
	ProductName   string
	OrderType     string
	OrderPrice    int
	OrderQuantity int
}
