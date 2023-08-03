package coinconverter

import (
	"fmt"
)

func ConvertToAwesomeFromTo(from, to string) string {
	convertedFromTo := fmt.Sprintf("%s-%s", from, to)

	return convertedFromTo
}

func CurrencySymbol(to string) string {
	switch to {
	case "BRL":
		return "R$"
	case "USD":
		return "$"
	case "EURO":
		return "€"
	case "BTC":
		return "₿"
	}

	return "currency symbol not found"
}

func ConvertCurrentyAmount(amount, value float64) float64 {
	return amount * value
}
