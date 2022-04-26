// Copyright 2021-present Anon. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package commands

import (
	"github.com/anonistas/notya/pkg"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// initCommand is a setup command of notya.
var initCommand = &cobra.Command{
	Use:     "init",
	Aliases: []string{"setup"},
	Short:   "Initialize application related files/folders",
	Run:     runInitCommand,
}

// initSetupCommand adds initCommand to main application command.
func initSetupCommand() {
	appCommand.AddCommand(initCommand)
}

// runInitCommand runs appropriate functionalities to setup notya and make it ready-to-use.
func runInitCommand(cmd *cobra.Command, args []string) {
	loading.Start()
	err := service.Init()
	loading.Stop()

	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	pkg.Alert(pkg.SuccessL, `Application initialized successfully`)
	pkg.Print(" > [notya -h/help] for help", color.FgBlue)
}
