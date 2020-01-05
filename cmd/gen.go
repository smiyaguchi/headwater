package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	schema string
	output string

	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate test date",
		Long:  "Generate test date for big data",
		Run: func(cmd *cobra.Command, args []string) {
			if schema, err := cmd.PersistentFlags().GetString("schema"); err == nil {
				fmt.Println("schema:", schema)
			}
			if output, err := cmd.PersistentFlags().GetString("output"); err == nil {
				fmt.Println("output:", output)
			}
		},
	}
)

func init() {
	genCmd.PersistentFlags().StringP("schema", "s", "schema.json", "input schema file")
	genCmd.PersistentFlags().StringP("output", "o", "testdata.csv", "output test data file")

	rootCmd.AddCommand(genCmd)
}
