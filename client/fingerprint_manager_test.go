// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"runtime"
	"testing"

	"github.com/openwonton/openwonton/ci"
	"github.com/openwonton/openwonton/client/config"
	"github.com/stretchr/testify/require"
)

func requireCPUFrequency(t *testing.T, attributes map[string]string) {
	t.Helper()
	if attributes["cpu.frequency"] != "" {
		return
	}
	if runtime.GOOS == "darwin" {
		if attributes["cpu.frequency.power"] != "" || attributes["cpu.frequency.efficiency"] != "" {
			return
		}
	}
	t.Fatalf("expected cpu frequency attribute")
}

func requireNoCPUFrequency(t *testing.T, attributes map[string]string) {
	t.Helper()
	require.NotContains(t, attributes, "cpu.frequency")
	require.NotContains(t, attributes, "cpu.frequency.power")
	require.NotContains(t, attributes, "cpu.frequency.efficiency")
}

func TestFingerprintManager_Run_ResourcesFingerprint(t *testing.T) {
	ci.Parallel(t)
	require := require.New(t)

	testClient, cleanup := TestClient(t, nil)
	defer cleanup()

	fm := NewFingerprintManager(
		testClient.config.PluginSingletonLoader,
		testClient.GetConfig,
		testClient.config.Node,
		testClient.shutdownCh,
		testClient.updateNodeFromFingerprint,
		testClient.logger,
	)

	err := fm.Run()
	require.Nil(err)

	node := testClient.config.Node

	require.NotEqual(0, node.Resources.CPU)
	require.NotEqual(0, node.Resources.MemoryMB)
	require.NotZero(node.Resources.DiskMB)
}

func TestFimgerprintManager_Run_InWhitelist(t *testing.T) {
	ci.Parallel(t)
	require := require.New(t)

	testClient, cleanup := TestClient(t, func(c *config.Config) {
		c.Options = map[string]string{
			"test.shutdown_periodic_after":    "true",
			"test.shutdown_periodic_duration": "2",
		}
	})
	defer cleanup()

	fm := NewFingerprintManager(
		testClient.config.PluginSingletonLoader,
		testClient.GetConfig,
		testClient.config.Node,
		testClient.shutdownCh,
		testClient.updateNodeFromFingerprint,
		testClient.logger,
	)

	err := fm.Run()
	require.Nil(err)

	node := testClient.config.Node

	requireCPUFrequency(t, node.Attributes)
}

func TestFingerprintManager_Run_InDenylist(t *testing.T) {
	ci.Parallel(t)
	require := require.New(t)

	testClient, cleanup := TestClient(t, func(c *config.Config) {
		c.Options = map[string]string{
			"fingerprint.allowlist": "  arch,memory,foo,bar	",
			"fingerprint.denylist":  "  cpu	",
		}
	})
	defer cleanup()

	fm := NewFingerprintManager(
		testClient.config.PluginSingletonLoader,
		testClient.GetConfig,
		testClient.config.Node,
		testClient.shutdownCh,
		testClient.updateNodeFromFingerprint,
		testClient.logger,
	)

	err := fm.Run()
	require.Nil(err)

	node := testClient.config.Node

	requireNoCPUFrequency(t, node.Attributes)
	require.NotEqual(node.Attributes["memory.totalbytes"], "")
}

func TestFingerprintManager_Run_Combination(t *testing.T) {
	ci.Parallel(t)
	require := require.New(t)

	testClient, cleanup := TestClient(t, func(c *config.Config) {
		c.Options = map[string]string{
			"fingerprint.allowlist": "  arch,cpu,memory,foo,bar	",
			"fingerprint.denylist":  "  memory,host	",
		}
	})
	defer cleanup()

	fm := NewFingerprintManager(
		testClient.config.PluginSingletonLoader,
		testClient.GetConfig,
		testClient.config.Node,
		testClient.shutdownCh,
		testClient.updateNodeFromFingerprint,
		testClient.logger,
	)

	err := fm.Run()
	require.Nil(err)

	node := testClient.config.Node

	requireCPUFrequency(t, node.Attributes)
	require.NotEqual(node.Attributes["cpu.arch"], "")
	require.NotContains(node.Attributes, "memory.totalbytes")
	require.NotContains(node.Attributes, "os.name")
}

func TestFingerprintManager_Run_CombinationLegacyNames(t *testing.T) {
	ci.Parallel(t)
	require := require.New(t)

	testClient, cleanup := TestClient(t, func(c *config.Config) {
		c.Options = map[string]string{
			"fingerprint.whitelist": "  arch,cpu,memory,foo,bar	",
			"fingerprint.blacklist": "  memory,host	",
		}
	})
	defer cleanup()

	fm := NewFingerprintManager(
		testClient.config.PluginSingletonLoader,
		testClient.GetConfig,
		testClient.config.Node,
		testClient.shutdownCh,
		testClient.updateNodeFromFingerprint,
		testClient.logger,
	)

	err := fm.Run()
	require.Nil(err)

	node := testClient.config.Node

	requireCPUFrequency(t, node.Attributes)
	require.NotEqual(node.Attributes["cpu.arch"], "")
	require.NotContains(node.Attributes, "memory.totalbytes")
	require.NotContains(node.Attributes, "os.name")
}
