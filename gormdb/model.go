package gormdb

import (
	"time"

	"github.com/google/uuid"
)

type Conversion struct {
	Id            uuid.UUID `gorm:"column:id;primary_key;" json:"id"`
	From          string    `gorm:"column:from;type:varchar(50);not null;" json:"from"`
	To            string    `gorm:"column:to;type:varchar(50);not null;" json:"to"`
	Amount        float64   `gorm:"column:amount;type:decimal(15,2);not null;" json:"amount"`
	CurrentyValue float64   `gorm:"column:currentValue;type:decimal(15,2);not null;" json:"currentValue"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;not null" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime;" json:"updatedAt"`
}
