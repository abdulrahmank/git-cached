package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var GitCCmd = &cobra.Command{
	Use:   "gitc",
	Short: "Caches all the files ignored using .gitignore",
}

type CommandArgs struct {
	Path string
}

var cmdArg CommandArgs

func Execute() {
	if err := GitCCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}