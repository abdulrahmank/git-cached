package cmd

import (
	"fmt"
	"github.com/abdulrahmank/git-cached/compressor"
	"github.com/abdulrahmank/git-cached/parser"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
	"os"
)

var cmdCache = &cobra.Command{
	Use:   "cache",
	Run:   Cache,
	Short: "Cache all the file anf directories ignored in git",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	GitCCmd.AddCommand(cmdCache)
}

func Cache(_ *cobra.Command, args []string) {
	cmdArg.Path = args[0]
	if &cmdArg.Path == nil || len(cmdArg.Path) == 0 {
		log.Print("Please pass location as first argument")
	}
	log.Printf("Running git cached in location: %s", cmdArg.Path)
	ignoredFilesRegex, err := parser.Parse(cmdArg.Path + "/.gitignore")
	if err != nil {
		log.Fatalf("Unable to parse git ignore file: %s", err.Error())
		return
	}
	bytes := getCommitHash()
	if err != nil {
		log.Fatalf("Unable to get hash %s", err.Error())
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
	log.Println(fmt.Sprintf("Wrote git ignored files to file: %v/%s.zip",
		os.TempDir(), previousHash))
}

func getCommitHash() plumbing.Hash {
	r, e := git.PlainOpen(cmdArg.Path)
	if e != nil {
		log.Fatal(e.Error())
		return plumbing.Hash{}
	}
	ref, _ := r.Head()
	return ref.Hash()
}