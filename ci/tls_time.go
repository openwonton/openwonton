// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ci

import "os"

const (
	tlsTestTimeEnv     = "NOMAD_TLS_TEST_TIME"
	tlsTestTimeDefault = "2024-01-01T00:00:00Z"
)

func init() {
	if os.Getenv(tlsTestTimeEnv) == "" {
		_ = os.Setenv(tlsTestTimeEnv, tlsTestTimeDefault)
	}
}
