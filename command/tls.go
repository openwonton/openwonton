// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"os"
	"strings"

	"github.com/mitchellh/cli"
)

type TLSCommand struct {
	Meta
}

func fileDoesNotExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return true
	}
	return false
}

func (c *TLSCommand) Help() string {
	helpText := `
Usage: wonton tls <subcommand> <subcommand> [options]

This command groups subcommands for creating certificates for OpenWonton TLS
configuration. The TLS command allows operators to generate self signed
certificates to use when securing your OpenWonton cluster.

Some simple examples for creating certificates can be found here.
More detailed examples are available in the subcommands or the documentation.

Create a CA

    $ wonton tls ca create

Create a server certificate

    $ wonton tls cert create -server

Create a client certificate

    $ wonton tls cert create -client

`
	return strings.TrimSpace(helpText)
}

func (c *TLSCommand) Synopsis() string {
	return "Generate self-signed TLS certificates for OpenWonton"
}

func (c *TLSCommand) Name() string { return "tls" }

func (c *TLSCommand) Run(_ []string) int {
	return cli.RunResultHelp
}
