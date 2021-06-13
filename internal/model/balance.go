package model

type Wallet struct {
	ID      uint    `gorm: "primaryKey"`
	Balance float32 `gorm: "not null,default:'0'"`
}
