// Copyright (c) 2025 OpenWonton Authors.
// SPDX-License-Identifier: MPL-2.0

package structs

// Clean-room replacement; see CLEAN_ROOM_NOTES.md.
import (
	"testing"

	nstructs "github.com/openwonton/openwonton/nomad/structs"
)

func TestCSIVolumeMountOptionsToCSIMountOptions(t *testing.T) {
	if got := (*CSIVolumeMountOptions)(nil).ToCSIMountOptions(); got != nil {
		t.Fatalf("expected nil for nil receiver, got %#v", got)
	}

	opts := &CSIVolumeMountOptions{
		Filesystem: "ext4",
		MountFlags: []string{"noatime", "nodiratime"},
	}
	got := opts.ToCSIMountOptions()
	if got == nil {
		t.Fatal("expected non-nil mount options")
	}
	if got.FSType != "ext4" {
		t.Fatalf("expected FSType ext4, got %q", got.FSType)
	}
	if len(got.MountFlags) != 2 || got.MountFlags[0] != "noatime" || got.MountFlags[1] != "nodiratime" {
		t.Fatalf("unexpected mount flags: %#v", got.MountFlags)
	}
}

func TestClientCSIControllerValidateVolumeRequestToCSIRequest(t *testing.T) {
	req := &ClientCSIControllerValidateVolumeRequest{
		VolumeID: "vol-1",
		VolumeCapabilities: []*nstructs.CSIVolumeCapability{
			{
				AttachmentMode: nstructs.CSIVolumeAttachmentModeFilesystem,
				AccessMode:     nstructs.CSIVolumeAccessModeSingleNodeReader,
			},
		},
		MountOptions: &nstructs.CSIMountOptions{
			FSType:     "xfs",
			MountFlags: []string{"noatime"},
		},
		Secrets: nstructs.CSISecrets{
			"token": "redacted",
		},
	}

	out, err := req.ToCSIRequest()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.ExternalID != "vol-1" {
		t.Fatalf("expected ExternalID vol-1, got %q", out.ExternalID)
	}
	if len(out.Capabilities) != 1 {
		t.Fatalf("expected one capability, got %d", len(out.Capabilities))
	}
	if out.Secrets["token"] != "redacted" {
		t.Fatalf("expected secret token to be preserved")
	}
}

func TestClientCSIControllerAttachVolumeRequestToCSIRequest(t *testing.T) {
	req := &ClientCSIControllerAttachVolumeRequest{
		VolumeID:        "vol-2",
		ClientCSINodeID: "node-1",
		AttachmentMode:  nstructs.CSIVolumeAttachmentModeFilesystem,
		AccessMode:      nstructs.CSIVolumeAccessModeSingleNodeWriter,
		ReadOnly:        false,
		Secrets: nstructs.CSISecrets{
			"token": "redacted",
		},
	}

	out, err := req.ToCSIRequest()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.ExternalID != "vol-2" {
		t.Fatalf("expected ExternalID vol-2, got %q", out.ExternalID)
	}
	if out.NodeID != "node-1" {
		t.Fatalf("expected NodeID node-1, got %q", out.NodeID)
	}
	if out.ReadOnly {
		t.Fatalf("expected ReadOnly false")
	}
	if out.Secrets["token"] != "redacted" {
		t.Fatalf("expected secret token to be preserved")
	}
}
