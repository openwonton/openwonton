// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fingerprint

import (
	"runtime"
	"testing"

	"github.com/openwonton/openwonton/ci"
	"github.com/openwonton/openwonton/client/config"
	"github.com/openwonton/openwonton/helper/testlog"
	"github.com/openwonton/openwonton/nomad/structs"
)

func TestHostFingerprint(t *testing.T) {
	ci.Parallel(t)

	f := NewHostFingerprint(testlog.HCLogger(t))
	node := &structs.Node{
		Attributes: make(map[string]string),
	}

	request := &FingerprintRequest{Config: &config.Config{}, Node: node}
	var response FingerprintResponse
	err := f.Fingerprint(request, &response)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if !response.Detected {
		t.Fatalf("expected response to be applicable")
	}

	if len(response.Attributes) == 0 {
		t.Fatalf("should generate a diff of node attributes")
	}

	commonAttributes := []string{"os.name", "os.version", "unique.hostname", "kernel.name"}
	nonWindowsAttributes := append(commonAttributes, "kernel.version")
	windowsAttributes := append(commonAttributes, "os.build")

	expectedAttributes := nonWindowsAttributes
	if runtime.GOOS == "windows" {
		expectedAttributes = windowsAttributes
	}

	// Host info
	for _, key := range expectedAttributes {
		assertNodeAttributeContains(t, response.Attributes, key)
	}
}
