package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        *string
	Price       decimal.Decimal
	StockNumber int
	StockCode   string
	Description string
	CategoryId  [16]byte
	Category    ProductCategory
}
