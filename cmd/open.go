package cmd

import (
	"fmt"

	"./utils"
	"./utils/prompt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var openEncrypted bool

func PickItem(items []string, urls []string) {
	if len(items) == 0 {
		fmt.Println("No bookmarks yet !")
		fmt.Println("Add them with \"browzy new\", or check help with \"browzy\"")
	} else {
		prompt := promptui.Select{
			Label: "Select an item",
			Items: items,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		index := utils.IndexOf(result, items)
		if index == -1 {
			fmt.Printf("Something went wrong: could not find element")
		} else {
			// fmt.Printf("%s\n", urls[index])
			utils.OpenBrowser(urls[index])
		}
	}
}

func Browzy(cmd *cobra.Command, args []string) {
	var secret []byte
	if openEncrypted {
		pwd, err := prompt.PromptPassword()
		utils.Check(err)
		secret = []byte(pwd + pwd)[0:16]

		Filepath = Filepath + ".enc"
	}

	items, urls, err := utils.GetItemsAndUrls(Filepath, secret)
	if err == nil {
		PickItem(items, urls)
	}
}

var pickCmd = &cobra.Command{
	Use:   "open",
	Short: "Browse your bookmarks and open them in your browser",
	Run:   Browzy,
}

func init() {
	pickCmd.PersistentFlags().BoolVarP(&openEncrypted, "encrypt", "e", false, "open an encrypted bookmarks file")
	rootCmd.AddCommand(pickCmd)
}
