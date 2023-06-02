package cmds

import (
	"github.com/spf13/cobra"
	"go-slim/internal/build"
)

var CronCmd = &cobra.Command{
	Use:     "cron",
	Short:   "run cron task",
	Example: "cli cron",
	Run: func(cmd *cobra.Command, args []string) {
		// https://github.com/go-co-op/gocron
		cli := build.BuildCli()
		cli.Cron.NewScheduler().StartBlocking()
	},
}

func init() {
	RootCmd.AddCommand(CronCmd)
}
