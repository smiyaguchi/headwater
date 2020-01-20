package cmd

import (
	"fmt"

	"github.com/smiyaguchi/headwater/generator"
	"github.com/smiyaguchi/headwater/schema"
	"github.com/spf13/cobra"
)

var (
	schemaFile string
	count int
	loss bool
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate test date",
	Long:  "Generate test date for big data",
	Run: func(cmd *cobra.Command, args []string) {
		schemaFile, _ = cmd.PersistentFlags().GetString("schema")
		count, _ = cmd.PersistentFlags().GetInt("count")
		loss, _ = cmd.PersistentFlags().GetBool("loss")

		var s schema.Schema
		s, err := schema.ReadFile(schemaFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		generator.Generate(s, count, loss)
	},
}

func init() {
	genCmd.PersistentFlags().StringP("schema", "s", "schema.json", "input schema file")
	genCmd.PersistentFlags().StringP("output", "o", "testdata.csv", "output test data file")
	genCmd.PersistentFlags().IntP("count", "c", 1000, "generate count")
	genCmd.PersistentFlags().BoolP("loss", "l", false, "include null values in data")

	rootCmd.AddCommand(genCmd)
}
