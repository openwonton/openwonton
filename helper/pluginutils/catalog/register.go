// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package catalog

import (
	"github.com/openwonton/openwonton/drivers/docker"
	"github.com/openwonton/openwonton/drivers/exec"
	"github.com/openwonton/openwonton/drivers/java"
	"github.com/openwonton/openwonton/drivers/qemu"
	"github.com/openwonton/openwonton/drivers/rawexec"
)

// This file is where all builtin plugins should be registered in the catalog.
// Plugins with build restrictions should be placed in the appropriate
// register_XXX.go file.
func init() {
	RegisterDeferredConfig(rawexec.PluginID, rawexec.PluginConfig, rawexec.PluginLoader)
	Register(exec.PluginID, exec.PluginConfig)
	Register(qemu.PluginID, qemu.PluginConfig)
	Register(java.PluginID, java.PluginConfig)
	RegisterDeferredConfig(docker.PluginID, docker.PluginConfig, docker.PluginLoader)
}
