// Copyright 2022-present Anon. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package commands

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/anonistas/notya/assets"
	"github.com/anonistas/notya/lib/services"
	"github.com/anonistas/notya/pkg"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Overwrites [Y] service's data with [X] service (in case of [X] service being current running service)",
	Run:   runMigrateCommand,
}

func initMigrateCommand() {
	appCommand.AddCommand(migrateCommand)
}

func runMigrateCommand(cmd *cobra.Command, args []string) {
	determineService()
	loading.Start()

	availableServices := []string{}
	// Generate a list of availabe services
	// by not including current service.
	for _, s := range services.Services {
		if service.Type() == s {
			continue
		}

		availableServices = append(availableServices, s)
	}

	loading.Stop()

	// Ask for servie selection.
	var selected string
	survey.AskOne(
		assets.ChooseRemotePrompt(availableServices),
		&selected,
	)
	selectedService := serviceFromType(selected, true)

	loading.Start()
	migratedNodes, errs := service.Migrate(selectedService)
	loading.Stop()

	if len(migratedNodes) == 0 && len(errs) == 0 {
		pkg.Print("Everything up-to-date", color.FgHiGreen)
		return
	}

	pkg.PrintErrors("migrate", errs)
	pkg.Alert(pkg.SuccessL, fmt.Sprintf("Migrated %v nodes", len(migratedNodes)))
}
