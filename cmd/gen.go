package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/smiyaguchi/headwater/schema"
	"github.com/spf13/cobra"
)

var schemaFile string
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate test date",
	Long:  "Generate test date for big data",
	Run: func(cmd *cobra.Command, args []string) {
		schemaFile, _ = cmd.PersistentFlags().GetString("schema")

		fmt.Println(schemaFile)
		bytes, err := ioutil.ReadFile(schemaFile)
		if err != nil {
			panic(err)
		}

		var schema []schema.Schema
		if err := json.Unmarshal(bytes, &schema); err != nil {
			panic(err)
		}
		fmt.Println(schema)
	},
}

func init() {
	genCmd.PersistentFlags().StringP("schema", "s", "schema.json", "input schema file")
	genCmd.PersistentFlags().StringP("output", "o", "testdata.csv", "output test data file")

	rootCmd.AddCommand(genCmd)
}
