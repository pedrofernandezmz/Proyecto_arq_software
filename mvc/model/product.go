package model

type Product struct {
	ProductId   int     `gorm:"primaryKey;AUTO_INCREMENT"`
	CategoryId  int     `gorm:"foreignKey:CategoryId"`
	Name        string  `gorm:"type:varchar(100);not null;unique;"`
	Price       float32 `gorm:"type:decimal;not null;"`
	Picture     string  `gorm:"type:varchar(100);not null;"`
}

type Products []Product
