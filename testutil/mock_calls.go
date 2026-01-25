// Copyright (c) 2025 OpenWonton Authors.
// SPDX-License-Identifier: MPL-2.0

package testutil

// Clean-room replacement; see CLEAN_ROOM_NOTES.md.
import (
	"sync"
	"sync/atomic"

	"github.com/mitchellh/go-testing-interface"
)

func NewCallCounter() *CallCounter {
	return &CallCounter{}
}

// CallCounter tracks named call counts for tests.
type CallCounter struct {
	counts sync.Map
}

func (c *CallCounter) Inc(name string) {
	counter := c.counterFor(name)
	counter.Add(1)
}

func (c *CallCounter) Get() map[string]int {
	out := make(map[string]int)
	c.counts.Range(func(key, value any) bool {
		name, ok := key.(string)
		if !ok {
			return true
		}
		counter, ok := value.(*atomic.Int64)
		if !ok {
			return true
		}
		out[name] = int(counter.Load())
		return true
	})
	return out
}

func (c *CallCounter) AssertCalled(t testing.T, name string) {
	t.Helper()
	counts := c.Get()
	if _, ok := counts[name]; !ok {
		t.Errorf("expected %q to be called; counts: %v", name, counts)
	}
}

func (c *CallCounter) counterFor(name string) *atomic.Int64 {
	if value, ok := c.counts.Load(name); ok {
		return value.(*atomic.Int64)
	}
	counter := &atomic.Int64{}
	actual, _ := c.counts.LoadOrStore(name, counter)
	return actual.(*atomic.Int64)
}
