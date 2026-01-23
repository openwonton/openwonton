// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type ACLPolicyCommand struct {
	Meta
}

func (f *ACLPolicyCommand) Help() string {
	helpText := `
Usage: wonton acl policy <subcommand> [options] [args]

  This command groups subcommands for interacting with ACL policies. OpenWonton
  provides a Nomad-compatible ACL system to control access to data and APIs.
  ACL policies allow a set of capabilities or actions to be granted or
  allowlisted. For a full guide see:
  https://openwonton.io/docs/concepts/acl

  Create an ACL policy:

      $ wonton acl policy apply <name> <policy-file>

  List ACL policies:

      $ wonton acl policy list

  Inspect an ACL policy:

      $ wonton acl policy info <policy>

  Please see the individual subcommand help for detailed usage information.
`
	return strings.TrimSpace(helpText)
}

func (f *ACLPolicyCommand) Synopsis() string {
	return "Interact with ACL policies"
}

func (f *ACLPolicyCommand) Name() string { return "acl policy" }

func (f *ACLPolicyCommand) Run(args []string) int {
	return cli.RunResultHelp
}
