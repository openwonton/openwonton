// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"fmt"
	"testing"
	"time"

	"github.com/mitchellh/cli"
	"github.com/openwonton/openwonton/api"
	"github.com/openwonton/openwonton/ci"
	"github.com/openwonton/openwonton/testutil"
	"github.com/stretchr/testify/require"
)

func TestServiceListCommand_Run(t *testing.T) {
	ci.Parallel(t)

	srv, client, url := testServer(t, true, nil)
	defer srv.Shutdown()

	// Wait until our test node is ready.
	testutil.WaitForResult(func() (bool, error) {
		nodes, _, err := client.Nodes().List(nil)
		if err != nil {
			return false, err
		}
		if len(nodes) == 0 {
			return false, fmt.Errorf("missing node")
		}
		if _, ok := nodes[0].Drivers["mock_driver"]; !ok {
			return false, fmt.Errorf("mock_driver not ready")
		}
		return true, nil
	}, func(err error) {
		require.NoError(t, err)
	})

	ui := cli.NewMockUi()
	cmd := &ServiceListCommand{
		Meta: Meta{
			Ui:          ui,
			flagAddress: url,
		},
	}

	// Run the command with some random arguments to ensure we are performing
	// this check.
	require.Equal(t, 1, cmd.Run([]string{"-address=" + url, "pretty-please"}))
	require.Contains(t, ui.ErrorWriter.String(), "This command takes no arguments")
	ui.ErrorWriter.Reset()

	// Create a test job with a Nomad service.
	testJob := testJob("service-discovery-nomad-list")
	testJob.TaskGroups[0].Tasks[0].Services = []*api.Service{
		{Name: "service-discovery-nomad-list", Provider: "nomad", Tags: []string{"foo", "bar"}}}

	// Register that job.
	regResp, _, err := client.Jobs().Register(testJob, nil)
	require.NoError(t, err)
	registerCode := waitForSuccess(ui, client, fullId, t, regResp.EvalID)
	require.Equal(t, 0, registerCode)

	// Reset the output writer, otherwise we will have additional information here.
	ui.OutputWriter.Reset()

	// Job register doesn't assure the service registration has completed.
	// Wait for the service to appear before running the command assertions.
	testutil.WaitForResultUntil(testutil.Timeout(10*time.Second), func() (bool, error) {
		list, _, err := client.Services().List(nil)
		if err != nil {
			return false, err
		}
		for _, reg := range list {
			for _, svc := range reg.Services {
				if svc.ServiceName == "service-discovery-nomad-list" {
					return true, nil
				}
			}
		}
		return false, fmt.Errorf("service not registered")
	}, func(err error) {
		require.NoError(t, err)
	})

	// Perform a standard lookup.
	require.Equal(t, 0, cmd.Run([]string{"-address=" + url}))

	// Test each header and data entry.
	s := ui.OutputWriter.String()
	require.Contains(t, s, "Service Name")
	require.Contains(t, s, "Tags")
	require.Contains(t, s, "service-discovery-nomad-list")
	require.Contains(t, s, "[bar,foo]")

	// Perform a wildcard namespace lookup.
	code := cmd.Run([]string{"-address=" + url, "-namespace", "*"})
	require.Equal(t, 0, code)

	// Test each header and data entry.
	s = ui.OutputWriter.String()
	require.Contains(t, s, "Service Name")
	require.Contains(t, s, "Namespace")
	require.Contains(t, s, "Tags")
	require.Contains(t, s, "service-discovery-nomad-list")
	require.Contains(t, s, "default")
	require.Contains(t, s, "[bar,foo]")

	ui.OutputWriter.Reset()
	ui.ErrorWriter.Reset()
}
