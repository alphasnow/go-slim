package cmds

import (
	"errors"
	"github.com/jinzhu/inflection"
	"github.com/spf13/cobra"
)

var pluralCmd = &cobra.Command{
	Use:     "plural",
	Short:   "Get the plural of the word",
	Example: "cli plural text word",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("no word")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, word := range args {
			cmd.Printf("%s -> plural = %s \n", word, inflection.Plural(word))
		}
	},
}

func init() {
	RootCmd.AddCommand(pluralCmd)
}
