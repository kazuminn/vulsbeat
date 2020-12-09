package main

import (
	"os"

	"github.com/kazuminn/vulsbeat/cmd"

	_ "github.com/kazuminn/vulsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
