package cmd

import (
	_ "embed"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"
)

var (
	dryRun bool

	//go:embed schemas/v3.1.json
	schemaV31 []byte
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
		if dryRun {
			log.Println("dry run")
		}

		if fileName == "" {
			return errors.New("input --file")
		}

		sch, err := jsonschema.CompileString("", string(schemaV31))
		if err != nil {
			log.Fatalf("%#v", err)
		}

		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		var v interface{}
		if err := json.Unmarshal(data, &v); err != nil {
			log.Fatal(err)
		}

		if err = sch.Validate(v); err != nil {
			log.Fatalf("%#v", err)
		}

		log.Println("Validate OK!!")
		return nil
	},
}
