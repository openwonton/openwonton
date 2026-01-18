// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testutils

import (
	"runtime"
	"strings"
	"testing"

	"github.com/docker/docker/libnetwork/resolvconf"
	"github.com/openwonton/openwonton/plugins/drivers"
	"github.com/shoenig/test/must"
)

// TestTaskDNSConfig asserts that a task is running with the given DNSConfig
func TestTaskDNSConfig(t *testing.T, driver *DriverHarness, taskID string, dns *drivers.DNSConfig) {
	t.Run("dns_config", func(t *testing.T) {
		caps, err := driver.Capabilities()
		must.NoError(t, err)

		// FS isolation is used here as a proxy for network isolation.
		// This is true for the current built-in drivers but it is not necessarily so.
		isolated := caps.FSIsolation != drivers.FSIsolationNone
		usesHostNetwork := caps.FSIsolation != drivers.FSIsolationImage

		if !isolated {
			t.Skip("dns config not supported on non isolated drivers")
		}

		// write to a file and check it presence in host
		r := execTask(t, driver, taskID, `cat /etc/resolv.conf`,
			false, "")
		must.Zero(t, r.exitCode)

		resolvConf := []byte(strings.TrimSpace(r.stdout))
		containerNameservers := resolvconf.GetNameservers(resolvConf, resolvconf.IP)

		if dns != nil {
			if len(dns.Servers) > 0 {
				must.SliceContainsAll(t, dns.Servers, containerNameservers)
			}
			if len(dns.Searches) > 0 {
				must.SliceContainsAll(t, dns.Searches, resolvconf.GetSearchDomains(resolvConf))
			}
			if len(dns.Options) > 0 {
				must.SliceContainsAll(t, dns.Options, resolvconf.GetOptions(resolvConf))
			}
		} else {
			if runtime.GOOS == "darwin" && !usesHostNetwork && isDockerDesktopResolver(containerNameservers) {
				must.Greater(t, 0, len(containerNameservers))
				return
			}

			systemPath := "/etc/resolv.conf"
			if !usesHostNetwork {
				systemPath = resolvconf.Path()
			}

			system, specificErr := resolvconf.GetSpecific(systemPath)
			must.NoError(t, specificErr)
			must.SliceContainsAll(t, resolvconf.GetNameservers(system.Content, resolvconf.IP), containerNameservers)
			must.SliceContainsAll(t, resolvconf.GetSearchDomains(system.Content), resolvconf.GetSearchDomains(resolvConf))
			must.SliceContainsAll(t, resolvconf.GetOptions(system.Content), resolvconf.GetOptions(resolvConf))
		}
	})
}

func isDockerDesktopResolver(nameservers []string) bool {
	if len(nameservers) != 1 {
		return false
	}
	return nameservers[0] == "127.0.0.11" || nameservers[0] == "192.168.65.7"
}
