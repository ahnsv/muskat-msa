package v1

import (
	"net/http"

	"github.com/ahnsv/muskat-msa/pkg/app"
	"github.com/ahnsv/muskat-msa/pkg/e"
	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	gApp := app.Gin{C: c}
	gApp.Response(
		http.StatusOK, e.SUCCESS, map[string]string{
			"hello": "world",
		},
	)
}
