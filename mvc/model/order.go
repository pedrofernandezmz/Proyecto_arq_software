package model

type Order struct {
	OrderId    int     `gorm:"primaryKey;AUTO_INCREMENT"`
	UserId     int     `gorm:"type:int"`
	Date       string  `gorm:"type:date;not null;"`
	Total      float32 `gorm:"type:decimal;not null;"`
}

type Orders []Order
