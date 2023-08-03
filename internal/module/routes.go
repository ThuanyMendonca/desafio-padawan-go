package module

import (
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module/converter"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	converter.Router(e.Group("/exchange"))
}
