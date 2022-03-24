package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "simple-openapi-validator",
		Short: "simple-openapi-validator",
		Long:  `simple-openapi-validator`,
	}

	fileName string
)

func Execute() error {
	return rootCmd.Execute()
}
