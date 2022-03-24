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
	"sigs.k8s.io/yaml"
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

func detectVersion(target string) (Version, error) {
	log.Println("Detecting version...")
	v := gjson.Get(target, "openapi")

	if v.Str == "" {
		return UNKNOWN, errors.New("invalid format")
	}

	log.Println("version: " + v.Str)

	s := strings.Split(v.Str, ".")
	if s[0] == "3" && s[1] == "0" {
		return V30, nil
	} else if s[0] == "3" && s[1] == "1" {
		return V31, nil
	}
	return UNKNOWN, nil
}

func readFromFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func yamlToJSON(target string) (string, error) {
	j, err := yaml.YAMLToJSON([]byte(target))
	if err != nil {
		return "", err
	}
	return string(j), nil
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate",
	Long:  `validate`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fileName == "" {
			return errors.New("input --file")
		}

		target, err := readFromFile(fileName)
		if err != nil {
			log.Fatalf(err.Error())
		}

		target, err = yamlToJSON(target)
		if err != nil {
			log.Fatalf(err.Error())
		}

		version, err := detectVersion(target)
		if err != nil {
			log.Fatalf(err.Error())
		}

		var schema string
		switch version {
		case V30:
			schema = string(schemaV30)
		case V31:
			schema = string(schemaV31)
		default:
			log.Fatal("unknown version")
		}

		sch, err := jsonschema.CompileString("", schema)
		if err != nil {
			log.Fatalf(err.Error())
		}

		var v interface{}
		if err := json.Unmarshal([]byte(target), &v); err != nil {
			log.Fatal(err)
		}

		if err = sch.Validate(v); err != nil {
			log.Fatalf(err.Error())
		}

		log.Println("Validate OK!!")
		return nil
	},
}
