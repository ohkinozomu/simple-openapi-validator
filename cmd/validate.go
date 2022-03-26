package cmd

import (
	"errors"
	"log"

	"github.com/ohkinozomu/simple-openapi-validator/pkg/validator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringVarP(&fileName, "file", "f", "", "file name")
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate",
	Long:  `validate`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fileName == "" {
			return errors.New("input --file")
		}

		err := validator.Validate(fileName)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}
