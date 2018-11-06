package main

import (
	"github.com/spf13/cobra"
)

var cmdInit = &cobra.Command{
	Use:   "init",
	Run:   runGitCached,
	Short: "Initialize a new cluster",
}

func init() {
	cmdInit.Execute()
}

func runGitCached(c *cobra.Command, args []string) {

}