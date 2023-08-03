package repository

import (
	"github.com/ThuanyMendonca/desafio-padawan-go/gormdb"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter/domain"
	"gorm.io/gorm"
)

type ConverterReposity struct {
	db *gorm.DB
}

func NewConverterRepository(db *gorm.DB) domain.IConverterRepository {
	return &ConverterReposity{db}
}

func (b *ConverterReposity) Add(currencyConverter *gormdb.Conversion) error {
	return b.db.Create(currencyConverter).Error
}
