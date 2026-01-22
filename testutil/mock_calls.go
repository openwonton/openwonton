// Copyright (c) 2025 OpenWonton Authors.
// SPDX-License-Identifier: MPL-2.0

package testutil

import (
	"sync"

	"github.com/mitchellh/go-testing-interface"
)

// CallCounter tracks named call counts for tests.
type CallCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewCallCounter() *CallCounter {
	return &CallCounter{
		counts: make(map[string]int),
	}
}

func (c *CallCounter) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[name] = c.counts[name] + 1
}

func (c *CallCounter) Get() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	clone := make(map[string]int, len(c.counts))
	for key, value := range c.counts {
		clone[key] = value
	}
	return clone
}

func (c *CallCounter) AssertCalled(t testing.T, name string) {
	t.Helper()
	counts := c.Get()
	if _, ok := counts[name]; !ok {
		t.Errorf("expected %q to be called; counts: %v", name, counts)
	}
}
