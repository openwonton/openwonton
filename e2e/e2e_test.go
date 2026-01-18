// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// This package exists to wrap our e2e provisioning and test framework so that it
// can be run via 'go test ./e2e'. See './framework/framework.go'
package e2e

import (
	"os"
	"testing"

	"github.com/openwonton/openwonton/e2e/framework"

	_ "github.com/openwonton/openwonton/e2e/affinities"
	_ "github.com/openwonton/openwonton/e2e/clientstate"
	_ "github.com/openwonton/openwonton/e2e/connect"
	_ "github.com/openwonton/openwonton/e2e/consul"
	_ "github.com/openwonton/openwonton/e2e/consultemplate"
	_ "github.com/openwonton/openwonton/e2e/csi"
	_ "github.com/openwonton/openwonton/e2e/deployment"
	_ "github.com/openwonton/openwonton/e2e/eval_priority"
	_ "github.com/openwonton/openwonton/e2e/events"
	_ "github.com/openwonton/openwonton/e2e/isolation"
	_ "github.com/openwonton/openwonton/e2e/lifecycle"
	_ "github.com/openwonton/openwonton/e2e/metrics"
	_ "github.com/openwonton/openwonton/e2e/networking"
	_ "github.com/openwonton/openwonton/e2e/nomadexec"
	_ "github.com/openwonton/openwonton/e2e/oversubscription"
	_ "github.com/openwonton/openwonton/e2e/parameterized"
	_ "github.com/openwonton/openwonton/e2e/periodic"
	_ "github.com/openwonton/openwonton/e2e/podman"
	_ "github.com/openwonton/openwonton/e2e/quotas"
	_ "github.com/openwonton/openwonton/e2e/remotetasks"
	_ "github.com/openwonton/openwonton/e2e/scaling"
	_ "github.com/openwonton/openwonton/e2e/scalingpolicies"
	_ "github.com/openwonton/openwonton/e2e/scheduler_sysbatch"
	_ "github.com/openwonton/openwonton/e2e/scheduler_system"
	_ "github.com/openwonton/openwonton/e2e/spread"
	_ "github.com/openwonton/openwonton/e2e/taskevents"
	_ "github.com/openwonton/openwonton/e2e/vaultsecrets"

	// these are no longer on the old framework but by importing them
	// we get a quick check that they compile on every commit
	_ "github.com/openwonton/openwonton/e2e/disconnectedclients"
	_ "github.com/openwonton/openwonton/e2e/namespaces"
	_ "github.com/openwonton/openwonton/e2e/nodedrain"
	_ "github.com/openwonton/openwonton/e2e/rescheduling"
	_ "github.com/openwonton/openwonton/e2e/volumes"
)

func TestE2E(t *testing.T) {
	if os.Getenv("NOMAD_E2E") == "" {
		t.Skip("Skipping e2e tests, NOMAD_E2E not set")
	} else {
		framework.Run(t)
	}
}
