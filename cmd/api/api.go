package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahnsv/muskat-msa/internal/api"
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
		routerInit := api.InitRouter()
		endpoint := fmt.Sprintf(":%d", 8080)
		server := &http.Server{
			Handler: routerInit,
			Addr:    endpoint,
		}
		log.Printf("API Server starts on %s", endpoint)
		server.ListenAndServe()
	},
}

func init() {
	Cmd.AddCommand(runCmd)
}
