package main

import (
	"./cmd"
	"./cmd/utils"
)

func main() {
	utils.SetupCloseHandler()
	cmd.Execute()
}
