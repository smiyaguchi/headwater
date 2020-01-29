package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of headwater",
	Long:  "Print the version number of headwater",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.2.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
