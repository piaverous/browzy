package cmd

import (
	"os"

	"./utils"
	"./utils/prompt"

	"github.com/spf13/cobra"
)

var encrypt bool

func NewBookmark(cmd *cobra.Command, args []string) {
	name, err := prompt.PromptName()
	utils.Check(err)
	url, err := prompt.PromptUrl()
	utils.Check(err)

	result := name + " - " + url
	if encrypt {
		password, err := prompt.PromptPassword()
		utils.Check(err)

		result, err = utils.Encrypt(result, []byte(password + password)[0:16])
		utils.Check(err)
	}

	browzyFilePath := Filepath
	if encrypt {
		browzyFilePath = browzyFilePath + ".enc"
	}
	utils.InitBookmarksFile(browzyFilePath)
	file, err := os.OpenFile(browzyFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	utils.Check(err)

	_, err = file.WriteString(result + "\n")
	utils.Check(err)
	file.Close()
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new bookmark",
	Run:   NewBookmark,
}

func init() {
	newCmd.PersistentFlags().BoolVarP(&encrypt, "encrypt", "e", false, "encrypt a bookmark entry file")
	rootCmd.AddCommand(newCmd)
}
