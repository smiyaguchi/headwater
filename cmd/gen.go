package cmd

import (
	"fmt"

	"github.com/smiyaguchi/headwater/generator"
	"github.com/smiyaguchi/headwater/schema"
	"github.com/spf13/cobra"
)

var (
	schemaFile string
	count      int
	loss       bool
	mode       string
	header     bool
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate test date",
	Long:  "Generate test date for big data",
	Run: func(cmd *cobra.Command, args []string) {
		schemaFile, _ = cmd.PersistentFlags().GetString("schema")
		count, _ = cmd.PersistentFlags().GetInt("count")
		loss, _ = cmd.PersistentFlags().GetBool("loss")
		mode, _ = cmd.PersistentFlags().GetString("mode")
		header, _ = cmd.PersistentFlags().GetBool("header")

		var s schema.Schema
		s, err := schema.ReadFile(schemaFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		g := generator.New(mode)
		g.Generate(s, count, loss, header)
	},
}

func init() {
	genCmd.PersistentFlags().StringP("schema", "s", "schema.json", "input schema file")
	genCmd.PersistentFlags().IntP("count", "c", 1000, "generate count")
	genCmd.PersistentFlags().BoolP("loss", "l", false, "include null values in data")
	genCmd.PersistentFlags().StringP("mode", "m", "NORMAL", "generate mode")
	genCmd.PersistentFlags().BoolP("header", "", true, "header flag")

	rootCmd.AddCommand(genCmd)
}
