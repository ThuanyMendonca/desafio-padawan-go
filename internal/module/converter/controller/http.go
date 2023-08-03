package controller

import (
	"net/http"
	"strconv"

	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter/domain"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter/service"
	"github.com/gin-gonic/gin"
)

type IConverterController interface {
	GetCurrentyConverter(c *gin.Context)
}

type ConverterController struct {
	converterService service.ICurrentyConverter
}

func NewConverterController(converterService service.ICurrentyConverter) IConverterController {
	return &ConverterController{converterService}
}

func (cc *ConverterController) GetCurrentyConverter(c *gin.Context) {
	amount := c.Param("amount")
	rate := c.Param("rate")

	amountParsed, _ := strconv.ParseFloat(amount, 64)
	rateParsed, _ := strconv.ParseFloat(rate, 64)

	params := &domain.CurrencyConverter{
		From:   domain.CurrencyEnum(c.Param("from")),
		To:     domain.CurrencyEnum(c.Param("to")),
		Amount: amountParsed,
		Rate:   rateParsed,
	}

	currentyConverted, err := cc.converterService.RegisterCurrentyConverter(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, currentyConverted)
}
