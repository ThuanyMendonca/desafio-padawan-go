package server

import (
	"fmt"
	"net/http"

	"github.com/ThuanyMendonca/desafio-padawan-go/config/env"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	env.Load()

	gin.SetMode(env.GinMode)
	r := gin.New()

	r.Use(gin.LoggerWithWriter(gin.DefaultErrorWriter, "/"))
	r.Use(gin.CustomRecovery(PanicFilter))

	return r
}

func PanicFilter(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		fmt.Println(err)
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
}
