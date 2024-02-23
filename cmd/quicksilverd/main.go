package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/quicksilver-zone/quicksilver/app"
	cmdcfg "github.com/quicksilver-zone/quicksilver/cmd/config"
	webui "github.com/quicksilver-zone/quicksilver/web-ui"
)

func main() {
	cmdcfg.SetupConfig()
	cmdcfg.RegisterDenoms()
	webui.UI() // Start the web server.  TODO: ensure that this is easy for operators to turn off, and that it handles cors or whatever.

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	app.DefaultNodeHome = filepath.Join(userHomeDir, ".quicksilverd")

	rootCmd, _ := NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "QUICKSILVERD", app.DefaultNodeHome); err != nil {
		var exitError *server.ErrorCode
		if errors.As(err, &exitError) {
			os.Exit(exitError.Code)
		}

		os.Exit(1)
	}

	os.Exit(0)
}
