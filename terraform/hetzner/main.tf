provider "hcloud" {}

locals {
  ssh_public_key_path  = pathexpand(var.ssh_public_key_path)
  ssh_private_key_path = pathexpand(var.ssh_private_key_path)

  server_ips = [
    for i in range(var.server_count) : cidrhost(var.network_cidr, var.server_ip_offset + i)
  ]
  client_ips = [
    for i in range(var.client_count) : cidrhost(var.network_cidr, var.client_ip_offset + i)
  ]

  retry_join = join("\",\"", local.server_ips)

  wonton_binary_path = var.wonton_binary_path != "" ? pathexpand(var.wonton_binary_path) : abspath("${path.root}/../../pkg/linux_amd64/wonton")
}

resource "hcloud_ssh_key" "main" {
  name       = "${var.name}-ssh"
  public_key = file(local.ssh_public_key_path)
}

resource "hcloud_network" "cluster" {
  name     = var.name
  ip_range = var.network_cidr
}

resource "hcloud_network_subnet" "cluster" {
  network_id   = hcloud_network.cluster.id
  type         = "cloud"
  network_zone = var.network_zone
  ip_range     = var.network_cidr
}

resource "hcloud_firewall" "cluster" {
  name = "${var.name}-firewall"

  rule {
    direction  = "in"
    protocol   = "tcp"
    port       = "22"
    source_ips = var.ssh_allow_cidrs
  }

  rule {
    direction  = "in"
    protocol   = "tcp"
    port       = "4646"
    source_ips = var.ui_allow_cidrs
  }

  rule {
    direction  = "in"
    protocol   = "tcp"
    port       = "8500"
    source_ips = var.ui_allow_cidrs
  }

  rule {
    direction  = "in"
    protocol   = "tcp"
    port       = "8200"
    source_ips = var.ui_allow_cidrs
  }

  rule {
    direction  = "in"
    protocol   = "icmp"
    source_ips = [var.network_cidr]
  }

  rule {
    direction  = "in"
    protocol   = "tcp"
    port       = "1-65535"
    source_ips = [var.network_cidr]
  }

  rule {
    direction  = "in"
    protocol   = "udp"
    port       = "1-65535"
    source_ips = [var.network_cidr]
  }
}

resource "hcloud_server" "server" {
  count       = var.server_count
  name        = "${var.name}-server-${count.index}"
  image       = var.image
  server_type = var.server_type
  location    = var.location

  ssh_keys     = [hcloud_ssh_key.main.id]
  firewall_ids = [hcloud_firewall.cluster.id]

  public_net {
    ipv4_enabled = true
    ipv6_enabled = false
  }

  network {
    network_id = hcloud_network.cluster.id
    ip         = local.server_ips[count.index]
  }

  labels = {
    role = "server"
  }

  depends_on = [hcloud_network_subnet.cluster]

  connection {
    host        = self.ipv4_address
    type        = "ssh"
    user        = "root"
    private_key = file(local.ssh_private_key_path)
    timeout     = var.ssh_timeout
  }

  provisioner "remote-exec" {
    inline = [
      "mkdir -p /ops",
      "chmod 777 /ops",
    ]
  }

  provisioner "file" {
    source      = "${path.module}/../shared"
    destination = "/ops"
  }

  provisioner "file" {
    source      = "${path.module}/../examples"
    destination = "/ops"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /ops/shared/scripts/setup.sh /ops/shared/scripts/server.sh",
      "/ops/shared/scripts/setup.sh",
      "bash /ops/shared/scripts/server.sh \"hcloud\" \"${var.server_count}\" '${local.retry_join}' \"${var.nomad_binary_url}\" \"${local.server_ips[count.index]}\"",
    ]
  }

  provisioner "file" {
    source      = local.wonton_binary_path
    destination = "/tmp/wonton"
  }

  provisioner "remote-exec" {
    inline = [
      "install -m 0755 /tmp/wonton /usr/local/bin/wonton",
      "ln -sf /usr/local/bin/wonton /usr/local/bin/nomad",
      "systemctl restart nomad",
    ]
  }
}

resource "hcloud_server" "client" {
  count       = var.client_count
  name        = "${var.name}-client-${count.index}"
  image       = var.image
  server_type = var.client_type
  location    = var.location

  ssh_keys     = [hcloud_ssh_key.main.id]
  firewall_ids = [hcloud_firewall.cluster.id]

  public_net {
    ipv4_enabled = true
    ipv6_enabled = false
  }

  network {
    network_id = hcloud_network.cluster.id
    ip         = local.client_ips[count.index]
  }

  labels = {
    role = "client"
  }

  depends_on = [hcloud_network_subnet.cluster]

  connection {
    host        = self.ipv4_address
    type        = "ssh"
    user        = "root"
    private_key = file(local.ssh_private_key_path)
    timeout     = var.ssh_timeout
  }

  provisioner "remote-exec" {
    inline = [
      "mkdir -p /ops",
      "chmod 777 /ops",
    ]
  }

  provisioner "file" {
    source      = "${path.module}/../shared"
    destination = "/ops"
  }

  provisioner "file" {
    source      = "${path.module}/../examples"
    destination = "/ops"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /ops/shared/scripts/setup.sh /ops/shared/scripts/client.sh",
      "/ops/shared/scripts/setup.sh",
      "bash /ops/shared/scripts/client.sh \"hcloud\" '${local.retry_join}' \"${var.nomad_binary_url}\" \"${local.client_ips[count.index]}\"",
    ]
  }

  provisioner "file" {
    source      = local.wonton_binary_path
    destination = "/tmp/wonton"
  }

  provisioner "remote-exec" {
    inline = [
      "install -m 0755 /tmp/wonton /usr/local/bin/wonton",
      "ln -sf /usr/local/bin/wonton /usr/local/bin/nomad",
      "systemctl restart nomad",
    ]
  }
}
