package cmd

import (
	"log"
	"os/user"

	"./utils"

	"github.com/spf13/cobra"
)

var Filepath string

var (
	rootCmd = &cobra.Command{
		Use:   "browzy",
		Short: "A generator for Kubernetes based Applications",
		Long: `Browzy is a CLI for managing Web bookmarks directly 
in your terminal.`,
	}
)

// Execute executes the root command.
func Execute() error {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	var defaultFilePath string = usr.HomeDir + "/.browzy"

	rootCmd.PersistentFlags().StringVarP(&Filepath, "file", "f", defaultFilePath, "path to bookmarks file")
	utils.InitBookmarksFile(Filepath)

	return rootCmd.Execute()
}
