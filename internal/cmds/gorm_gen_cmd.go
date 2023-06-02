package cmds

import (
	"github.com/spf13/cobra"
	"go-slim/internal/build"
)

var GormGenCmd = &cobra.Command{
	Use:     "gormgen",
	Short:   "gorm gen model",
	Example: "cli gormgen",
	Run: func(cmd *cobra.Command, args []string) {
		// https://github.com/go-co-op/gocron
		cli := build.BuildCli()

		cli.GormGen.GenerateAllTable()
		cli.GormGen.Execute()
	},
}

func init() {
	RootCmd.AddCommand(GormGenCmd)
}
