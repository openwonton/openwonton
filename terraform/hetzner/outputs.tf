output "server_public_ipv4" {
  value       = hcloud_server.server[*].ipv4_address
  description = "Public IPv4 addresses for server instances."
}

output "client_public_ipv4" {
  value       = hcloud_server.client[*].ipv4_address
  description = "Public IPv4 addresses for client instances."
}

output "server_private_ips" {
  value       = local.server_ips
  description = "Private IP addresses assigned to server instances."
}

output "client_private_ips" {
  value       = local.client_ips
  description = "Private IP addresses assigned to client instances."
}

output "ssh_command" {
  value       = "ssh -i ${local.ssh_private_key_path} root@${hcloud_server.server[0].ipv4_address}"
  description = "Convenience SSH command for the first server."
}

output "nomad_http_addr" {
  value       = "http://${hcloud_server.server[0].ipv4_address}:4646"
  description = "Nomad HTTP address for the first server."
}

output "consul_http_addr" {
  value       = "http://${hcloud_server.server[0].ipv4_address}:8500"
  description = "Consul HTTP address for the first server."
}
