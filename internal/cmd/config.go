package cmd

import (
	"github.com/fernandoocampo/goku/internal/filesystems"
	"github.com/fernandoocampo/goku/internal/settings"
	"github.com/spf13/cobra"
)

var configCmd = cobra.Command{
	Use:   "config",
	Short: "setup gokucli configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		runSetup()
	},
}

func runSetup() {
	setting := settings.New(&filesystems.OSFileSystem{})
	err := setting.SetUpDefault()
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(&configCmd)
}
