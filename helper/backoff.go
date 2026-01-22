// Copyright (c) 2025 OpenWonton Authors.
// SPDX-License-Identifier: MPL-2.0

package helper

import "time"

// Backoff returns an exponential backoff duration capped by backoffLimit.
func Backoff(backoffBase time.Duration, backoffLimit time.Duration, attempt uint64) time.Duration {
	if backoffBase <= 0 {
		return 0
	}

	const maxInt64 = int64(^uint64(0) >> 1)
	if attempt > 62 {
		return backoffLimit
	}

	base := int64(backoffBase)
	if base <= 0 || base > (maxInt64>>attempt) {
		return backoffLimit
	}

	wait := time.Duration(base << attempt)
	if wait > backoffLimit {
		return backoffLimit
	}

	return wait
}
