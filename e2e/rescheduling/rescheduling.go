// Copyright (c) 2025 OpenWonton Authors.
// SPDX-License-Identifier: MPL-2.0

package rescheduling

import (
	"fmt"

	"github.com/openwonton/openwonton/e2e/e2eutil"
	"github.com/openwonton/openwonton/e2e/framework"
	"github.com/openwonton/openwonton/helper/uuid"
	"github.com/stretchr/testify/require"
)

type ReschedulingTest struct {
	framework.TC
	jobIDs []string
}

func init() {
	framework.AddSuites(&framework.TestSuite{
		Component:   "Rescheduling",
		CanRunLocal: true,
		Cases: []framework.TestCase{
			new(ReschedulingTest),
		},
	})
}

func (tc *ReschedulingTest) BeforeAll(f *framework.F) {
	e2eutil.WaitForLeader(f.T(), tc.Nomad())
}

func (tc *ReschedulingTest) AfterEach(f *framework.F) {
	for _, jobID := range tc.jobIDs {
		e2eutil.WaitForJobStopped(f.T(), tc.Nomad(), jobID)
	}
	tc.jobIDs = nil
}

func (tc *ReschedulingTest) TestReschedulingRegister(f *framework.F) {
	t := f.T()

	jobID := fmt.Sprintf("reschedule-%s", uuid.Generate()[0:8])
	tc.jobIDs = append(tc.jobIDs, jobID)

	require.NoError(t, e2eutil.Register(jobID, "rescheduling/input/rescheduling_default.nomad"))

	err := e2eutil.WaitForAllocStatusComparison(
		func() ([]string, error) { return e2eutil.AllocStatuses(jobID, "") },
		func(statuses []string) bool { return len(statuses) > 0 },
		nil,
	)
	require.NoError(t, err)
}
