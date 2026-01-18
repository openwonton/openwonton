// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type OperatorRaftCommand struct {
	Meta
}

func (c *OperatorRaftCommand) Help() string {
	helpText := `
Usage: wonton operator raft <subcommand> [options]

  This command groups subcommands for interacting with OpenWonton's Raft subsystem.
  The command can be used to verify Raft peers or in rare cases to recover
  quorum by removing invalid peers.

  List Raft peers:

      $ wonton operator raft list-peers

  Remove a Raft peer:

      $ wonton operator raft remove-peer -peer-address "IP:Port"

  Display info about the raft logs in the data directory:

      $ wonton operator raft info /var/wonton/data

  Display the log entries persisted in data dir in JSON format.

      $ wonton operator raft logs /var/wonton/data

  Display the server state obtained by replaying raft log entries
  persisted in data dir in JSON format.

      $ wonton operator raft state /var/wonton/data

  Please see the individual subcommand help for detailed usage information.


`
	return strings.TrimSpace(helpText)
}

func (c *OperatorRaftCommand) Synopsis() string {
	return "Provides access to the Raft subsystem"
}

func (c *OperatorRaftCommand) Name() string { return "operator raft" }

func (c *OperatorRaftCommand) Run(args []string) int {
	return cli.RunResultHelp
}
