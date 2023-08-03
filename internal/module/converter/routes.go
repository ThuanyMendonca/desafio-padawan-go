package converter

import (
	"github.com/ThuanyMendonca/desafio-padawan-go/config/dependency"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter/controller"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter/service"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	service := service.NewCurrentiesConverter(service.CurrentyConverterService{
		ConverterRepository: dependency.ConverterRepository,
	})

	c := controller.NewConverterController(service)

	r.GET("/:amount/:from/:to/:rate", c.GetCurrentyConverter)
}
