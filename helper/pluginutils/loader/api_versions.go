// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loader

import (
	"github.com/openwonton/openwonton/plugins/base"
	"github.com/openwonton/openwonton/plugins/device"
	"github.com/openwonton/openwonton/plugins/drivers"
)

var (
	// AgentSupportedApiVersions is the set of API versions supported by the
	// Nomad agent by plugin type.
	AgentSupportedApiVersions = map[string][]string{
		base.PluginTypeDevice: {device.ApiVersion010},
		base.PluginTypeDriver: {drivers.ApiVersion010},
	}
)
