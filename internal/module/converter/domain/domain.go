package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrRequiredSymbol = errors.New("symbol is required")
	ErrInvalidSymbol  = errors.New("an error has occurred: invalid symbol")
	ErrInvalidAmount  = errors.New("amount is invalid")
)

type CurrentyEnum string

const (
	Real   CurrentyEnum = "BRL"
	Dollar CurrentyEnum = "USD"
	Euro   CurrentyEnum = "EUR"
	Btc    CurrentyEnum = "BTC"
)

func (s CurrentyEnum) String() string {
	return string(s)
}

func (CurrentyEnum) Values() []string {
	return []string{
		Real.String(),
		Dollar.String(),
		Euro.String(),
		Btc.String(),
	}
}

func (s CurrentyEnum) Validate() error {
	if s.String() == "" {
		return ErrRequiredSymbol
	}

	for _, v := range s.Values() {
		if s.String() == v {
			return nil
		}
	}
	return ErrInvalidSymbol
}

type CurrencyConverter struct {
	Id     uuid.UUID    `json:"id"`
	From   CurrentyEnum `json:"from"`
	To     CurrentyEnum `json:"to"`
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
