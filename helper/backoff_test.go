// Copyright (c) 2025 OpenWonton Authors.
// SPDX-License-Identifier: MPL-2.0

package helper

import (
	"testing"
	"time"
)

func TestBackoffBasic(t *testing.T) {
	tests := []struct {
		name    string
		base    time.Duration
		limit   time.Duration
		attempt uint64
		expect  time.Duration
	}{
		{
			name:    "zero base returns zero",
			base:    0,
			limit:   5 * time.Second,
			attempt: 2,
			expect:  0,
		},
		{
			name:    "negative base returns zero",
			base:    -1 * time.Second,
			limit:   5 * time.Second,
			attempt: 1,
			expect:  0,
		},
		{
			name:    "attempt zero returns base",
			base:    250 * time.Millisecond,
			limit:   10 * time.Second,
			attempt: 0,
			expect:  250 * time.Millisecond,
		},
		{
			name:    "attempt one doubles",
			base:    250 * time.Millisecond,
			limit:   10 * time.Second,
			attempt: 1,
			expect:  500 * time.Millisecond,
		},
		{
			name:    "clamps to limit",
			base:    2 * time.Second,
			limit:   5 * time.Second,
			attempt: 2,
			expect:  5 * time.Second,
		},
		{
			name:    "large attempt clamps",
			base:    1 * time.Second,
			limit:   3 * time.Second,
			attempt: 63,
			expect:  3 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Backoff(tt.base, tt.limit, tt.attempt)
			if got != tt.expect {
				t.Fatalf("Backoff(%v, %v, %d)=%v; want %v", tt.base, tt.limit, tt.attempt, got, tt.expect)
			}
		})
	}
}
