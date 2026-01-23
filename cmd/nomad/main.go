// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	wontonPath, err := resolveWonton()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if os.Getenv("WONTON_NOMAD_SHIM_SILENT") == "" && os.Getenv("NOMAD_SHIM_SILENT") == "" {
		fmt.Fprintln(os.Stderr, "WARNING: `nomad` is a compatibility shim for OpenWonton. Prefer `wonton` directly.")
		fmt.Fprintln(os.Stderr, "Set WONTON_NOMAD_SHIM_SILENT=1 to disable this warning.")
		fmt.Fprintln(os.Stderr, "")
	}

	cmd := exec.Command(wontonPath, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func resolveWonton() (string, error) {
	if exe, err := os.Executable(); err == nil {
		dir := filepath.Dir(exe)
		if path, ok := firstExisting(candidatePaths(dir)); ok {
			return path, nil
		}
	}

	path, err := exec.LookPath("wonton")
	if err != nil {
		return "", fmt.Errorf("wonton executable not found; install OpenWonton or ensure `wonton` is on PATH")
	}

	return path, nil
}

func candidatePaths(dir string) []string {
	if runtime.GOOS == "windows" {
		return []string{
			filepath.Join(dir, "wonton.exe"),
			filepath.Join(dir, "wonton"),
		}
	}

	return []string{filepath.Join(dir, "wonton")}
}

func firstExisting(paths []string) (string, bool) {
	for _, path := range paths {
		info, err := os.Stat(path)
		if err == nil && !info.IsDir() {
			return path, true
		}
	}

	return "", false
}
