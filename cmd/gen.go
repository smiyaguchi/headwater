package cmd

import (
	"fmt"

	"github.com/smiyaguchi/headwater/generator"
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

		var s schema.Schema
		s, err := schema.ReadFile(schemaFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		generator.Generate(s)
	},
	
}

func init() {
	genCmd.PersistentFlags().StringP("schema", "s", "schema.json", "input schema file")
	genCmd.PersistentFlags().StringP("output", "o", "testdata.csv", "output test data file")

	rootCmd.AddCommand(genCmd)
}
