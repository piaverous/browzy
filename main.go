package main

import (
	"github.com/piaverous/browzy/cmd"
	"github.com/piaverous/browzy/cmd/utils"
)

func main() {
	utils.SetupCloseHandler()
	cmd.Execute()
}
