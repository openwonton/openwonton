# OpenWonton Project Overview

This document provides a summary of the OpenWonton project for agent reference.

## Core Project Information

*   **Project Name:** OpenWonton
*   **Genesis:** A community fork of HashiCorp Nomad.
*   **Primary Mandate:** To rename the fork, establishing a distinct identity aligned with other "Open" projects like OpenTofu and OpenBao.

## Key Goals & Requirements

*   **CLI Rename:** The primary command-line interface (CLI) will be renamed from `nomad` to `wonton`.
*   **Branding:** All project assets, including documentation, binaries, and release artifacts, will be updated to reflect the "OpenWonton" name.
*   **Compatibility:** A high degree of compatibility with the original Nomad is a priority. To support existing users and scripts, a `nomad` shim or symlink that invokes `wonton` is planned. This is intended to provide a low-friction migration path.
*   **Go Module Path:** A decision needs to be made on whether to keep the original Go module path (`github.com/hashicorp/nomad`) to minimize downstream breakage or to change it to a new path under an `openwonton` organization for a cleaner identity.
*   **Documentation:** A key deliverable is updated documentation, including a migration guide for users transitioning from `nomad` to `wonton`.

## Implementation High-Level Plan

The project plan is divided into several phases:
1.  **Phase 0: Decision & Prep:** Finalize key decisions like the Go module strategy and compatibility approach.
2.  **Phase 1: Brand + Docs:** Update `README.md`, documentation, and create a migration guide.
3.  **Phase 2: CLI Rename:** Implement the `wonton` CLI and the `nomad` compatibility shim.
4.  **Phase 3: CI + Releases:** Update CI/CD pipelines to produce `wonton`-named artifacts.
5.  **Phase 4: Validation:** Thoroughly test the changes, including integration and migration paths.
6.  **Phase 5: Release:** Publish the first official OpenWonton release.

## Agent's Role

The agent's primary role is to assist in executing this renaming plan. Key tasks will likely involve:
*   Finding and replacing references to "Nomad" and `nomad` in the codebase.
*   Updating documentation and examples.
*   Modifying build scripts and CI/CD pipeline configurations.
*   Assisting with the implementation of the `wonton` CLI and the `nomad` shim.

This summary is based on the `plan.md` file. It should be used as a reference for all future work on this project.

## Repository Structure

This section provides a high-level overview of the key directories in the OpenWonton repository.

*   `main.go`: The entry point for the `wonton` (formerly `nomad`) binary.
*   `command/`: Contains the implementation of the CLI subcommands (e.g., `agent`, `job`, `status`). This is where the CLI logic resides.
*   `nomad/`: This directory contains the core logic of OpenWonton, including the agent, server, and scheduling algorithms.
    *   `nomad/structs`: Defines the core data structures used throughout the project.
*   `client/`: Contains the client for interacting with the OpenWonton API, used by the CLI and other tools.
*   `api/`: Defines the data structures for the HTTP API.
*   `scheduler/`: The scheduling logic that determines how to place tasks onto client nodes.
*   `drivers/`: Implementations of task drivers for running different types of workloads (e.g., Docker, exec, Java, QEMU). Each driver is in its own subdirectory.
*   `ui/`: The source code for the web user interface. It's likely a single-page application (SPA).
*   `website/`: The source code for the project's public-facing website (e.g., openwonton.io).
*   `jobspec/` & `jobspec2/`: Code related to parsing and handling job specification files.
*   `acl/`: Access Control List (ACL) system implementation for security and authorization.
*   `e2e/`: End-to-end tests for the project.
*   `scripts/`: Various helper scripts for development, building, and releasing.
*   `terraform/`: Terraform configurations for deploying OpenWonton.

## Tests

Some tests can have a huge output, which can be a problem for token usage or could even freeze the computer. We need to be smart about it:
**Always** redirect output to a file, and check how many lines of code with `wc -l`. **If the number of lines is unreasonably big, do no not read it fully!! Find a smarter way to get the information you need.**
