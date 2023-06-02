package cmds

import (
	"github.com/spf13/cobra"
	"go-slim/internal/build"
)

var QueueCmd = &cobra.Command{
	Use:     "queue",
	Short:   "run queue task",
	Example: "cli queue",
	Run: func(cmd *cobra.Command, args []string) {
		cli := build.BuildCli()
		cli.Queue.StartBlocking()
	},
}

func init() {
	RootCmd.AddCommand(QueueCmd)
}
