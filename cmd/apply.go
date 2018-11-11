package cmd

import (
	"github.com/abdulrahmank/gitc/compressor"
	"github.com/spf13/cobra"
	"os"
)

var cmdApply = &cobra.Command{
	Use:   "apply",
	Run:   Apply,
	Short: "Cache all the file anf directories ignored in git",
}

func init() {
	GitCCmd.AddCommand(cmdApply)
}

func Apply(_ *cobra.Command, args []string) {
	if len(args) > 0 {
		cmdArg.Path = args[0]
	} else {
		cmdArg.Path = "."
	}

	hash := GetCommitHash()
	compressor.Decompress(os.TempDir(), hash.String(), cmdArg.Path)
}
