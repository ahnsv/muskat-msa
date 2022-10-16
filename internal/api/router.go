package api

import (
	v1 "github.com/ahnsv/muskat-msa/internal/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/placeOrder", v1.PlaceOrder)
	}
	return r
}
