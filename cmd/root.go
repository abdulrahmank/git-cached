package cmd

import (
	"fmt"
	"github.com/abdulrahmank/git-cached/compressor"
	"github.com/abdulrahmank/git-cached/parser"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
	"os"
	"gopkg.in/src-d/go-git.v4"
)

var rootCmd = &cobra.Command{
	Use:   "git-cached",
	Short: "Caches all the files ignored using .gitignore",
	Run:   RunGitCached,
}

type CommandArgs struct {
	Path string
}

var cmdArg CommandArgs

func init() {
	rootCmd.PersistentFlags().StringVar(&cmdArg.Path, "path", "",
		"Path of root git directory where .gitignore file exists")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RunGitCached(_ *cobra.Command, _ []string) {
	ignoredFilesRegex, err := parser.Parse("/Users/kabdul/go/src/github.com/abdulrahmank/git-cached/")
	if err != nil {
		log.Fatal("Unable to parse git ignore file")
		return
	}
	bytes := getCommitHash()
	if err != nil {
		log.Fatalf("Unable to get hash %v", err)
		return
	}
	previousHash := bytes.String()
	log.Println(
		fmt.Sprintf("Caching git ignored files for hash %s with bytes length %d",
			previousHash, len(bytes)))
	err = compressor.Compress(ignoredFilesRegex, previousHash,
		os.TempDir(), cmdArg.Path)
	if err != nil {
		log.Fatal("Unable to compress")
		return
	}
	log.Println(fmt.Sprintf("Wrote git ignored files to dir: %v", os.TempDir()))
}

func getCommitHash() plumbing.Hash {
	r, e := git.PlainOpen("/Users/kabdul/go/src/github.com/abdulrahmank/git-cached/")
	if e != nil {
		log.Fatal(e.Error())
		return plumbing.Hash{}
	}
	ref, _ := r.Head()
	return ref.Hash()
}