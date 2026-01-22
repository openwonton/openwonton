// Copyright (c) 2025 OpenWonton Authors.
// SPDX-License-Identifier: MPL-2.0

package csimanager

import (
	"context"
	"fmt"

	"github.com/openwonton/openwonton/client/pluginmanager"
	"github.com/openwonton/openwonton/nomad/structs"
	"github.com/openwonton/openwonton/plugins/csi"
	"github.com/openwonton/openwonton/testutil"
)

// MockExpandVolumeCall captures ExpandVolume arguments for assertions.
type MockExpandVolumeCall struct {
	VolID     string
	RemoteID  string
	AllocID   string
	UsageOpts *UsageOptions
	Capacity  *csi.CapacityRange
}

// MockVolumeManager is a test double for VolumeManager.
type MockVolumeManager struct {
	CallCounter          *testutil.CallCounter
	Mounts               map[string]bool
	NextMountVolumeErr   error
	NextUnmountVolumeErr error
	NextHasMountErr      error
	NextExpandVolumeErr  error
	LastExpandVolumeCall *MockExpandVolumeCall
	ExternalIDValue      string
}

func (m *MockVolumeManager) MountVolume(ctx context.Context, vol *structs.CSIVolume, alloc *structs.Allocation, usageOpts *UsageOptions, publishContext map[string]string) (*MountInfo, error) {
	if m.CallCounter != nil {
		m.CallCounter.Inc("MountVolume")
	}
	if m.NextMountVolumeErr != nil {
		err := m.NextMountVolumeErr
		m.NextMountVolumeErr = nil
		return nil, err
	}
	source := fmt.Sprintf("test-alloc-dir/%s/%s/%s", alloc.ID, vol.ID, usageOpts.ToFS())
	if m.Mounts != nil {
		m.Mounts[source] = true
	}
	return &MountInfo{Source: source}, nil
}

func (m *MockVolumeManager) UnmountVolume(ctx context.Context, volID, remoteID, allocID string, usageOpts *UsageOptions) error {
	if m.CallCounter != nil {
		m.CallCounter.Inc("UnmountVolume")
	}
	if m.NextUnmountVolumeErr != nil {
		err := m.NextUnmountVolumeErr
		m.NextUnmountVolumeErr = nil
		return err
	}
	if m.Mounts != nil {
		source := fmt.Sprintf("test-alloc-dir/%s/%s/%s", allocID, volID, usageOpts.ToFS())
		delete(m.Mounts, source)
	}
	return nil
}

func (m *MockVolumeManager) HasMount(ctx context.Context, mountInfo *MountInfo) (bool, error) {
	if m.CallCounter != nil {
		m.CallCounter.Inc("HasMount")
	}
	if m.NextHasMountErr != nil {
		err := m.NextHasMountErr
		m.NextHasMountErr = nil
		return false, err
	}
	if m.Mounts == nil {
		return false, nil
	}
	return m.Mounts[mountInfo.Source], nil
}

func (m *MockVolumeManager) ExpandVolume(ctx context.Context, volID, remoteID, allocID string, usageOpts *UsageOptions, capacity *csi.CapacityRange) (int64, error) {
	m.LastExpandVolumeCall = &MockExpandVolumeCall{
		VolID:     volID,
		RemoteID:  remoteID,
		AllocID:   allocID,
		UsageOpts: usageOpts,
		Capacity:  capacity,
	}
	if m.NextExpandVolumeErr != nil {
		err := m.NextExpandVolumeErr
		m.NextExpandVolumeErr = nil
		return 0, err
	}
	if capacity == nil {
		return 0, nil
	}
	return capacity.RequiredBytes, nil
}

func (m *MockVolumeManager) ExternalID() string {
	if m.ExternalIDValue == "" {
		return "mock-external-id"
	}
	return m.ExternalIDValue
}

// MockCSIManager is a test double for Manager.
type MockCSIManager struct {
	VM                      *MockVolumeManager
	NextWaitForPluginErr    error
	NextManagerForPluginErr error
}

func (m *MockCSIManager) PluginManager() pluginmanager.PluginManager {
	return nil
}

func (m *MockCSIManager) WaitForPlugin(ctx context.Context, pluginType, pluginID string) error {
	if m.NextWaitForPluginErr != nil {
		err := m.NextWaitForPluginErr
		m.NextWaitForPluginErr = nil
		return err
	}
	return nil
}

func (m *MockCSIManager) ManagerForPlugin(ctx context.Context, pluginID string) (VolumeManager, error) {
	if m.NextManagerForPluginErr != nil {
		err := m.NextManagerForPluginErr
		m.NextManagerForPluginErr = nil
		return nil, err
	}
	if m.VM == nil {
		return &MockVolumeManager{}, nil
	}
	return m.VM, nil
}

func (m *MockCSIManager) Shutdown() {}
