// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/openwonton/openwonton/api"
)

// Ensure ACLBindingRuleCommand satisfies the cli.Command interface.
var _ cli.Command = &ACLBindingRuleCommand{}

// ACLBindingRuleCommand implements cli.Command.
type ACLBindingRuleCommand struct {
	Meta
}

// Help satisfies the cli.Command Help function.
func (a *ACLBindingRuleCommand) Help() string {
	helpText := `
Usage: wonton acl binding-rule <subcommand> [options] [args]

  This command groups subcommands for interacting with ACL binding rules.
  OpenWonton provides a Nomad-compatible ACL system to control access to data
  and APIs. For a full guide see:
  https://openwonton.io/docs/concepts/acl

  Create an ACL binding rule:

      $ wonton acl binding-rule create \
          -auth-method=auth0 \
          -selector="nomad-engineering in list.groups" \
          -bind-type=role \
          -bind-name="custer-admin" \

  List all ACL binding rules:

      $ wonton acl binding-rule list

  Lookup a specific ACL binding rule:

      $ wonton acl binding-rule info <acl_binding_rule_id>

  Update an ACL binding rule:

      $ wonton acl binding-rule update \
          -description="wonton engineering team" \
          <acl_binding_rule_id>

  Delete an ACL binding rule:

      $ wonton acl binding-rule delete <acl_binding_rule_id>

  Please see the individual subcommand help for detailed usage information.
`
	return strings.TrimSpace(helpText)
}

// Synopsis satisfies the cli.Command Synopsis function.
func (a *ACLBindingRuleCommand) Synopsis() string { return "Interact with ACL binding rules" }

// Name returns the name of this command.
func (a *ACLBindingRuleCommand) Name() string { return "acl binding-rule" }

// Run satisfies the cli.Command Run function.
func (a *ACLBindingRuleCommand) Run(_ []string) int { return cli.RunResultHelp }

// formatACLBindingRule formats and converts the ACL binding rule API object
// into a string KV representation suitable for console output.
func formatACLBindingRule(aclBindingRule *api.ACLBindingRule) string {
	return formatKV([]string{
		fmt.Sprintf("ID|%s", aclBindingRule.ID),
		fmt.Sprintf("Description|%s", aclBindingRule.Description),
		fmt.Sprintf("Auth Method|%s", aclBindingRule.AuthMethod),
		fmt.Sprintf("Selector|%q", aclBindingRule.Selector),
		fmt.Sprintf("Bind Type|%s", aclBindingRule.BindType),
		fmt.Sprintf("Bind Name|%s", aclBindingRule.BindName),
		fmt.Sprintf("Create Time|%s", aclBindingRule.CreateTime),
		fmt.Sprintf("Modify Time|%s", aclBindingRule.ModifyTime),
		fmt.Sprintf("Create Index|%d", aclBindingRule.CreateIndex),
		fmt.Sprintf("Modify Index|%d", aclBindingRule.ModifyIndex),
	})
}
