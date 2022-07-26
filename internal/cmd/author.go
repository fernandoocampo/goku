package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authorCmd = cobra.Command{
	Use:   "author",
	Short: "Print author name for copyright attribution",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("https://github.com/fernandoocampo")
	},
}

func init() {
	rootCmd.AddCommand(&authorCmd)
}
