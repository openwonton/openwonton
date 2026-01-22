variable "name" {
  type        = string
  default     = "openwonton"
  description = "Prefix for Hetzner resources."
}

variable "location" {
  type        = string
  default     = "nbg1"
  description = "Hetzner location for the servers."
}

variable "network_zone" {
  type        = string
  default     = "eu-central"
  description = "Hetzner network zone for the private network."
}

variable "network_cidr" {
  type        = string
  default     = "10.10.0.0/24"
  description = "CIDR for the private network."
}

variable "server_count" {
  type        = number
  default     = 3
  description = "Number of server instances to deploy."
}

variable "client_count" {
  type        = number
  default     = 2
  description = "Number of client instances to deploy."
}

variable "server_type" {
  type        = string
  default     = "cpx22"
  description = "Server type for Nomad servers."
}

variable "client_type" {
  type        = string
  default     = "cpx22"
  description = "Server type for Nomad clients."
}

variable "image" {
  type        = string
  default     = "ubuntu-22.04"
  description = "Base image for the servers."
}

variable "server_ip_offset" {
  type        = number
  default     = 10
  description = "Starting host offset for server private IPs."
}

variable "client_ip_offset" {
  type        = number
  default     = 50
  description = "Starting host offset for client private IPs."
}

variable "ssh_public_key_path" {
  type        = string
  default     = "~/.ssh/hetzner_id_ed25519.pub"
  description = "Path to the SSH public key uploaded to Hetzner."
}

variable "ssh_private_key_path" {
  type        = string
  default     = "~/.ssh/hetzner_id_ed25519"
  description = "Path to the SSH private key used for provisioning."
}

variable "ssh_allow_cidrs" {
  type        = list(string)
  default     = ["0.0.0.0/0"]
  description = "CIDR blocks allowed to SSH into the servers."
}

variable "ui_allow_cidrs" {
  type        = list(string)
  default     = ["0.0.0.0/0"]
  description = "CIDR blocks allowed to access Nomad/Consul/Vault HTTP ports."
}

variable "nomad_binary_url" {
  type        = string
  default     = ""
  description = "Optional URL to replace the Nomad binary during provisioning."
}

variable "wonton_binary_path" {
  type        = string
  default     = ""
  description = "Optional local path to the wonton binary; defaults to ../../pkg/linux_amd64/wonton."
}

variable "ssh_timeout" {
  type        = string
  default     = "10m"
  description = "SSH timeout used by provisioners."
}
