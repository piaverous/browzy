package cmd

import (
	"github.com/piaverous/browzy/cmd/health"
	"github.com/piaverous/browzy/cmd/utils"
	"github.com/piaverous/browzy/cmd/utils/prompt"

	"github.com/spf13/cobra"
)

var checkEncrypted bool

func displayHealth(urls []string) {
	masterMind := health.CreateMasterMind(urls)
	masterMind.Start()
}

func Health(cmd *cobra.Command, args []string) {
	var secret []byte
	if checkEncrypted {
		pwd, err := prompt.PromptPassword()
		utils.Check(err)
		secret = []byte(pwd + pwd)[0:16]

		Filepath = Filepath + ".enc"
	}
	_, urls, err := utils.GetItemsAndUrls(Filepath, secret)
	if err == nil {
		displayHealth(urls)
	}
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "See the status of your bookmarks",
	Run:   Health,
}

func init() {
	healthCmd.PersistentFlags().BoolVarP(&checkEncrypted, "encrypt", "e", false, "checks encrypted bookmarks")
	rootCmd.AddCommand(healthCmd)
}
