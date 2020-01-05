package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hw",
	Short: "Headwater is generate test data",
	Long:  "Headwater is generate test date for big data",
}

func Execute() error {
	return rootCmd.Execute()
}
