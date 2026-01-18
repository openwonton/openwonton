# OpenWonton
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](LICENSE)

OpenWonton is a community-driven, open source, and Nomad-compatible workload orchestrator to deploy and manage containers (docker, podman), non-containerized applications (executable, Java), and virtual machines (qemu) across on-prem and clouds at scale. It is a fork of HashiCorp Nomad.

OpenWonton is supported on Linux, Windows, and macOS.

OpenWonton provides several key features:

* **Deploy Containers and Legacy Applications**: OpenWonton's flexibility as an orchestrator enables an organization to run containers, legacy, and batch applications together on the same infrastructure. OpenWonton brings core orchestration benefits to legacy applications without needing to containerize via pluggable task drivers.

* **Simple & Reliable**: OpenWonton runs as a single binary and is entirely self contained - combining resource management and scheduling into a single system. OpenWonton does not require any external services for storage or coordination. OpenWonton automatically handles application, node, and driver failures. OpenWonton is distributed and resilient, using leader election and state replication to provide high availability in the event of failures.

* **Device Plugins & GPU Support**: OpenWonton offers built-in support for GPU workloads such as machine learning (ML) and artificial intelligence (AI). OpenWonton uses device plugins to automatically detect and utilize resources from hardware devices such as GPU, FPGAs, and TPUs.

* **Federation for Multi-Region, Multi-Cloud**: OpenWonton was designed to support infrastructure at a global scale. OpenWonton supports federation out-of-the-box and can deploy applications across multiple regions and clouds.

* **Proven Scalability**: OpenWonton is optimistically concurrent, which increases throughput and reduces latency for workloads. OpenWonton has been proven to scale to clusters of 10K+ nodes in real-world production environments.

* **Ecosystem Integrations**: OpenWonton integrates seamlessly with tools like Terraform, Consul, and Vault for provisioning, service discovery, and secrets management.

## Quick Start

You can find Terraform manifests for bringing up a development OpenWonton cluster on a public cloud in the [`terraform`](terraform/) directory.

## Documentation

Full, comprehensive documentation is under development.

## Contributing

See the [`contributing`](contributing/) directory for more developer documentation.
