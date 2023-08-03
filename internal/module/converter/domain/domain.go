package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrRequiredCurrency = errors.New("symbol is required")
	ErrInvalidCurrency  = errors.New("an error has occurred: invalid symbol")
	ErrInvalidAmount    = errors.New("amount is invalid")
)

type CurrencyEnum string

const (
	Real   CurrencyEnum = "BRL"
	Dollar CurrencyEnum = "USD"
	Euro   CurrencyEnum = "EUR"
	Btc    CurrencyEnum = "BTC"
)

func (s CurrencyEnum) String() string {
	return string(s)
}

func (CurrencyEnum) Values() []string {
	return []string{
		Real.String(),
		Dollar.String(),
		Euro.String(),
		Btc.String(),
	}
}

func (s CurrencyEnum) Validate() error {
	if s.String() == "" {
		return ErrRequiredCurrency
	}

	for _, v := range s.Values() {
		if s.String() == v {
			return nil
		}
	}
	return ErrInvalidCurrency
}

type CurrencyConverter struct {
	Id     uuid.UUID    `json:"id"`
	From   CurrencyEnum `json:"from"`
	To     CurrencyEnum `json:"to"`
	Amount float64      `json:"amount"`
	Rate   float64      `json:"rate"`
}

func (c CurrencyConverter) ValidateCurrencyConvert() error {
	if err := c.From.Validate(); err != nil {
		return errors.New("an error has occurred to validate from currenty, details:" + err.Error())
	}

	if err := c.To.Validate(); err != nil {
		return errors.New("an error has occurred to validate to currenty, details:" + err.Error())
	}

	if c.Amount < 0 {
		return ErrInvalidAmount
	}

	return nil
}

func RegisterCurrencyConververt(data *CurrencyConverter) (*CurrencyConverter, error) {
	if err := data.ValidateCurrencyConvert(); err != nil {
		return nil, err
	}

	currency := &CurrencyConverter{
		Id:     uuid.New(),
		From:   data.From,
		To:     data.To,
		Amount: data.Amount,
		Rate:   data.Rate,
	}

	return currency, nil
}
