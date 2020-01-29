package cmd

import (
	"fmt"

	"github.com/smiyaguchi/headwater/config"
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
	quote      string
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate test date",
	Long:  "Generate test date for big data",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := schema.ReadFile(schemaFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		c := config.New(count, loss, header, quote)

		g := generator.New(mode)
		if err = g.Generate(s, c); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	genCmd.PersistentFlags().StringVarP(&schemaFile, "schema", "s", "schema.json", "input schema file")
	genCmd.PersistentFlags().IntVarP(&count, "count", "c", 1000, "generate count")
	genCmd.PersistentFlags().BoolVarP(&loss, "loss", "l", false, "include null values in data")
	genCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "NORMAL", "generate mode")
	genCmd.PersistentFlags().BoolVarP(&header, "header", "", true, "header flag")
	genCmd.PersistentFlags().StringVarP(&quote, "quote", "q", "NONE", "csv quote flag. Can use 'ALL', 'NonNumeric', 'None'.")

	rootCmd.AddCommand(genCmd)
}
