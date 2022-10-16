package cmd

import (
	"fmt"

	"github.com/ahnsv/muskat-msa/cmd/api"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   Name,
	Short: fmt.Sprintf("'%s' is a commerce solution powered by msa", Name),
}

func init() {
	RootCmd.AddCommand(
		api.Cmd,
	)
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
