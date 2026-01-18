// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consul

import (
	"context"

	"github.com/openwonton/openwonton/client/serviceregistration"
	"github.com/openwonton/openwonton/nomad/structs"
)

func NoopRestarter() serviceregistration.WorkloadRestarter {
	return noopRestarter{}
}

type noopRestarter struct{}

func (noopRestarter) Restart(ctx context.Context, event *structs.TaskEvent, failure bool) error {
	return nil
}
