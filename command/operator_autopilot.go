// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type OperatorAutopilotCommand struct {
	Meta
}

func (c *OperatorAutopilotCommand) Name() string { return "operator autopilot" }

func (c *OperatorAutopilotCommand) Run(args []string) int {
	return cli.RunResultHelp
}

func (c *OperatorAutopilotCommand) Synopsis() string {
	return "Provides tools for modifying Autopilot configuration"
}

func (c *OperatorAutopilotCommand) Help() string {
	helpText := `
Usage: wonton operator autopilot <subcommand> [options]

  This command groups subcommands for interacting with the OpenWonton Autopilot
  subsystem. Autopilot provides automatic, operator-friendly management of
  OpenWonton servers. The command can be used to view or modify the current
  Autopilot configuration. For a full guide see:
  https://openwonton.io/docs/configuration/autopilot

  Get the current Autopilot configuration:

      $ wonton operator autopilot get-config

  Set a new Autopilot configuration, enabling automatic dead server cleanup:

      $ wonton operator autopilot set-config -cleanup-dead-servers=true

  Please see the individual subcommand help for detailed usage information.
  `
	return strings.TrimSpace(helpText)
}
