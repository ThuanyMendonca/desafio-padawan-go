package domain

import "github.com/ThuanyMendonca/desafio-padawan-go/gormdb"

type IConverterRepository interface {
	Add(currencyConverter *gormdb.Conversion) error
}
