package cmds

import (
	"bytes"
	"github.com/spf13/cobra"
	"testing"
)

func executeCommand(thisCmd *cobra.Command, args ...string) (string, error) {

	buf := new(bytes.Buffer)
	thisCmd.SetOut(buf)
	thisCmd.SetErr(buf)
	thisCmd.SetArgs(args)

	_, err := thisCmd.ExecuteC()
	out := buf.String()
	return out, err
}

func Test_Serve(t *testing.T) {
	args := []string{"cron"}

	out, err := executeCommand(RootCmd, args...)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
