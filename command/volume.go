// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type VolumeCommand struct {
	Meta
}

func (c *VolumeCommand) Help() string {
	helpText := `
Usage: wonton volume <subcommand> [options]

  volume groups commands that interact with volumes.

  Register a new volume or update an existing volume:

      $ wonton volume register <input>

  Examine the status of a volume:

      $ wonton volume status <id>

  Deregister an unused volume:

      $ wonton volume deregister <id>

  Detach an unused volume:

      $ wonton volume detach <vol id> <node id>

  Create an external volume and register it:

      $ wonton volume create <input>

  Delete an external volume and deregister it:

      $ wonton volume delete <external id>

  Please see the individual subcommand help for detailed usage information.
`
	return strings.TrimSpace(helpText)
}

func (c *VolumeCommand) Name() string {
	return "volume"
}

func (c *VolumeCommand) Synopsis() string {
	return "Interact with volumes"
}

func (c *VolumeCommand) Run(args []string) int {
	return cli.RunResultHelp
}
