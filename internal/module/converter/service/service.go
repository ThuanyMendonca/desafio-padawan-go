package service

import (
	"time"

	"github.com/ThuanyMendonca/desafio-padawan-go/gormdb"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/common/coinconverter"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter/domain"
)

type ICurrentyConverter interface {
	RegisterCurrentyConverter(currencyRequest *domain.CurrencyConverter) (*ServiceResponse, error)
}

type CurrentyConverterService struct {
	ConverterRepository domain.IConverterRepository
}

func NewCurrentiesConverter(services CurrentyConverterService) ICurrentyConverter {
	return services
}

func (c CurrentyConverterService) RegisterCurrentyConverter(currencyRequest *domain.CurrencyConverter) (*ServiceResponse, error) {
	currencyConverter, err := domain.RegisterCurrencyConververt(currencyRequest)
	if err != nil {
		return nil, err
	}

	amountConverted := coinconverter.ConvertCurrentyAmount(currencyConverter.Amount, currencyConverter.Rate)
	currentySymbol := coinconverter.CurrencySymbol(string(currencyConverter.To))

	response := &ServiceResponse{
		SimboloMoeda:    currentySymbol,
		ValorConvertido: amountConverted,
	}

	if err := c.ConverterRepository.Add(&gormdb.Conversion{
		Id:            currencyConverter.Id,
		From:          currencyConverter.From.String(),
		To:            currencyConverter.To.String(),
		Amount:        currencyConverter.Amount,
		CurrentyValue: currencyConverter.Rate,
		CreatedAt:     time.Now(),
	}); err != nil {
		return nil, err
	}

	return response, nil
}
