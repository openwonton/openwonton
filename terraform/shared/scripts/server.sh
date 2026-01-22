#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


set -e

CONFIGDIR=/ops/shared/config

CONSULCONFIGDIR=/etc/consul.d
VAULTCONFIGDIR=/etc/vault.d
NOMADCONFIGDIR=/etc/nomad.d
CONSULTEMPLATECONFIGDIR=/etc/consul-template.d
HOME_USER=ubuntu
HOME_PATH=/home/ubuntu

# Wait for network
sleep 15

DOCKER_BRIDGE_IP_ADDRESS=(`ifconfig docker0 2>/dev/null|awk '/inet addr:/ {print $2}'|sed 's/addr://'`)
CLOUD=$1
SERVER_COUNT=$2
RETRY_JOIN=$3
NOMAD_BINARY=$4
IP_ADDRESS_OVERRIDE=$5

if [ "$CLOUD" = "hcloud" ] && [ ! -d "$HOME_PATH" ]; then
  HOME_USER=root
  HOME_PATH=/root
fi

# Get IP from metadata service
if [ -n "$IP_ADDRESS_OVERRIDE" ]; then
  IP_ADDRESS=$IP_ADDRESS_OVERRIDE
elif [ "$CLOUD" = "gce" ]; then
  IP_ADDRESS=$(curl -H "Metadata-Flavor: Google" http://metadata/computeMetadata/v1/instance/network-interfaces/0/ip)
elif [ "$CLOUD" = "hcloud" ]; then
  IP_ADDRESS=$(ip -4 -o addr show scope global | awk 'NR==1{print $4}' | cut -d/ -f1)
else
  IP_ADDRESS=$(curl -s http://instance-data/latest/meta-data/local-ipv4)
fi
# IP_ADDRESS="$(/sbin/ifconfig eth0 | grep 'inet addr:' | cut -d: -f2 | awk '{ print $1}')"

# Consul
sed -i "s/IP_ADDRESS/$IP_ADDRESS/g" $CONFIGDIR/consul.json
sed -i "s/SERVER_COUNT/$SERVER_COUNT/g" $CONFIGDIR/consul.json
sed -i "s/RETRY_JOIN/$RETRY_JOIN/g" $CONFIGDIR/consul.json
sudo cp $CONFIGDIR/consul.json $CONSULCONFIGDIR
sudo cp $CONFIGDIR/consul_$CLOUD.service /etc/systemd/system/consul.service

sudo systemctl enable consul.service
sudo systemctl start consul.service
sleep 10
export CONSUL_HTTP_ADDR=$IP_ADDRESS:8500
export CONSUL_RPC_ADDR=$IP_ADDRESS:8400

# Vault
sed -i "s/IP_ADDRESS/$IP_ADDRESS/g" $CONFIGDIR/vault.hcl
sudo cp $CONFIGDIR/vault.hcl $VAULTCONFIGDIR
sudo cp $CONFIGDIR/vault.service /etc/systemd/system/vault.service

sudo systemctl enable vault.service
sudo systemctl start vault.service

# Nomad

## Replace existing Nomad binary if remote file exists
if [ -n "$NOMAD_BINARY" ]; then
  if [[ `wget -S --spider $NOMAD_BINARY  2>&1 | grep 'HTTP/1.1 200 OK'` ]]; then
    curl -L $NOMAD_BINARY > nomad.zip
    sudo unzip -o nomad.zip -d /usr/local/bin
    sudo chmod 0755 /usr/local/bin/nomad
    sudo chown root:root /usr/local/bin/nomad
  fi
fi

sed -i "s/IP_ADDRESS/$IP_ADDRESS/g" $CONFIGDIR/nomad.hcl
sed -i "s/SERVER_COUNT/$SERVER_COUNT/g" $CONFIGDIR/nomad.hcl
sudo cp $CONFIGDIR/nomad.hcl $NOMADCONFIGDIR
sudo cp $CONFIGDIR/nomad.service /etc/systemd/system/nomad.service

sudo systemctl enable nomad.service
sudo systemctl start nomad.service
sleep 10
export NOMAD_ADDR=http://$IP_ADDRESS:4646

# Consul Template
sudo cp $CONFIGDIR/consul-template.hcl $CONSULTEMPLATECONFIGDIR/consul-template.hcl
sudo cp $CONFIGDIR/consul-template.service /etc/systemd/system/consul-template.service

# Add hostname to /etc/hosts

echo "127.0.0.1 $(hostname)" | sudo tee --append /etc/hosts

# Add Docker bridge network IP to /etc/resolv.conf (at the top)

echo "nameserver $DOCKER_BRIDGE_IP_ADDRESS" | sudo tee /etc/resolv.conf.new
cat /etc/resolv.conf | sudo tee --append /etc/resolv.conf.new
sudo mv /etc/resolv.conf.new /etc/resolv.conf

# Move examples directory to $HOME
sudo mv /ops/examples $HOME_PATH
sudo chown -R $HOME_USER:$HOME_USER $HOME_PATH/examples
sudo chmod -R 775 $HOME_PATH/examples

# Set env vars for tool CLIs
echo "export CONSUL_RPC_ADDR=$IP_ADDRESS:8400" | sudo tee --append $HOME_PATH/.bashrc
echo "export CONSUL_HTTP_ADDR=$IP_ADDRESS:8500" | sudo tee --append $HOME_PATH/.bashrc
echo "export VAULT_ADDR=http://$IP_ADDRESS:8200" | sudo tee --append $HOME_PATH/.bashrc
echo "export NOMAD_ADDR=http://$IP_ADDRESS:4646" | sudo tee --append $HOME_PATH/.bashrc
JAVA_HOME=$(readlink -f /usr/bin/java | sed "s:bin/java::")
echo "export JAVA_HOME=$JAVA_HOME"  | sudo tee --append $HOME_PATH/.bashrc
