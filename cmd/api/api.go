package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "api",
	Short: "api server",
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run an api server",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("/place", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "order placed",
			})
		})
		r.Run()
	},
}

func init() {
	Cmd.AddCommand(runCmd)
}
