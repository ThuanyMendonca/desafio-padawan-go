package server

import (
	"fmt"
	"net/http"

	"github.com/ThuanyMendonca/desafio-padawan-go/config/dependency"
	"github.com/ThuanyMendonca/desafio-padawan-go/config/env"
	"github.com/ThuanyMendonca/desafio-padawan-go/internal/module"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	env.Load()

	if err := dependency.Load(); err != nil {
		fmt.Println(err)
	}

	gin.SetMode(env.GinMode)
	r := gin.New()

	r.Use(gin.LoggerWithWriter(gin.DefaultErrorWriter, "/"))
	r.Use(gin.CustomRecovery(PanicFilter))

	module.Router(r)

	return r
}

func PanicFilter(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		fmt.Println(err)
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
}
