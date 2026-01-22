# Provision an OpenWonton cluster on Hetzner Cloud

This setup provisions a small OpenWonton (Nomad + Consul + Vault) sandbox on
Hetzner Cloud using Terraform and the shared provisioning scripts.

## Prerequisites

- Terraform 1.3+
- Hetzner Cloud API token in `HCLOUD_TOKEN` (loaded from `.env`)
- An SSH keypair available locally (defaults to `~/.ssh/hetzner_id_ed25519`)
- A local Linux `wonton` binary (defaults to `../../pkg/linux_amd64/wonton`)

## Configure

From the repository root, load your `.env` so Terraform can read `HCLOUD_TOKEN`:

```bash
set -a
source .env
set +a
```

Optional: create a `terraform.tfvars` in this directory to override defaults:

```hcl
ssh_public_key_path  = "~/.ssh/hetzner_id_ed25519.pub"
ssh_private_key_path = "~/.ssh/hetzner_id_ed25519"
location             = "nbg1"
network_zone         = "eu-central"
server_count         = 3
client_count         = 2
```

Build a Linux `wonton` binary (required when running from macOS):

```bash
make pkg/linux_amd64/wonton
```

If you are on macOS without cross-compilers, use Docker:

```bash
docker run --rm --platform=linux/amd64 -v "$PWD":/workspace -w /workspace golang:1.21 \
  bash -c "apt-get update && apt-get install -y git && make pkg/linux_amd64/wonton"
```

## Provision

```bash
cd terraform/hetzner
terraform init
terraform apply
```

## Validate

Terraform outputs a ready-to-use SSH command. Once connected, run:

```bash
consul members
wonton server members
wonton node status
```

## Cleanup

```bash
terraform destroy
```
