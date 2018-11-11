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
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	GitCCmd.AddCommand(cmdApply)
}

func Apply(_ *cobra.Command, args []string) {
	cmdArg.Path = args[0]
	hash := GetCommitHash()
	compressor.Decompress(os.TempDir(), hash.String(), cmdArg.Path)
}
