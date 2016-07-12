package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/cloudfoundry-incubator/furious-seven/commands"
	"github.com/jessevdk/go-flags"
)

func main() {
	if os.Getenv("V3") == "0" {
		parser := flags.NewParser(new(cfcli), flags.HelpFlag)
		_, err := parser.Parse()
		if err != nil {
			if flagsErr, ok := err.(*flags.Error); !ok || flagsErr.Type != flags.ErrUnknownCommand {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	} else {
		parser := flags.NewParser(new(cfcliv3), flags.HelpFlag)
		_, err := parser.Parse()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	}

	cmd := exec.Command("cf", os.Args...)
	_ = cmd.Run()
	os.Exit(0)
}

type cfcli struct {
	AppShow commands.AppCmd `command:"app" description:"show the app details"`
}

type cfcliv3 struct {
	AppShow commands.AppCmdv3 `command:"app" description:"show the app details"`
}
