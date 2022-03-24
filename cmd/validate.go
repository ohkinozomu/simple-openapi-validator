package cmd

import (
	_ "embed"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var (
	dryRun bool

	//go:embed schemas/v3.0.json
	schemaV30 []byte

	//go:embed schemas/v3.1.json
	schemaV31 []byte
)

type Version int

const (
	UNKNOWN Version = iota
	V30
	V31
)

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringVarP(&fileName, "file", "f", "", "file name")
}

func detectVersion(file string) (Version, error) {
	f, err := os.Open(file)
	if err != nil {
		return UNKNOWN, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return UNKNOWN, err
	}
	v := gjson.Get(string(b), "openapi")

	s := strings.Split(v.Str, ".")

	if s[0] == "3" && s[1] == "0" {
		return V30, nil
	} else if s[0] == "3" && s[1] == "1" {
		return V31, nil
	}
	return UNKNOWN, nil
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate",
	Long:  `validate`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fileName == "" {
			return errors.New("input --file")
		}

		version, err := detectVersion(fileName)
		if err != nil {
			log.Fatalf("%#v", err)
		}

		var schema string
		switch version {
		case V30:
			schema = string(schemaV30)
		case V31:
			schema = string(schemaV31)
		default:
			return errors.New("unknown version")
		}

		sch, err := jsonschema.CompileString("", schema)
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
