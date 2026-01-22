// Copyright (c) 2025 OpenWonton Authors.
// SPDX-License-Identifier: MPL-2.0

package state

import "testing"

func TestSliceIterator(t *testing.T) {
	it := NewSliceIterator()

	if got := it.Next(); got != nil {
		t.Fatalf("expected nil on empty iterator, got %#v", got)
	}

	it.Add("first")
	it.Add("second")

	if got := it.Next(); got != "first" {
		t.Fatalf("expected first element, got %#v", got)
	}
	if got := it.Next(); got != "second" {
		t.Fatalf("expected second element, got %#v", got)
	}
	if got := it.Next(); got != nil {
		t.Fatalf("expected nil after exhaustion, got %#v", got)
	}

	if it.WatchCh() != nil {
		t.Fatalf("expected nil watch channel")
	}
}
