package main

import (
	"github.com/tottokotkd/dockervlm/command"
	"github.com/tottokotkd/dockervlm/common"
	"os"
)

func main() {
	cd, err := os.Getwd()
	if err != nil {
		os.Exit(int(common.ExitCodeGetwdError))
	}
	env := common.Env{OutStream: os.Stdout, ErrStream: os.Stderr, CurrentDirectory: cd, Args: os.Args}
	exitCode := command.Run(env)
	os.Exit(int(exitCode))
}
