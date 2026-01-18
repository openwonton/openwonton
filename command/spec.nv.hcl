# A variable path can be specified in the specification file
# and will be used when writing the variable without specifying a
# path in the command or when writing JSON directly to the `/var/`
# HTTP API endpoint
# path = "path/to/variable"

# The Namespace to write the variable can be included in the specification. This
# value can be overridden by specifying the "-namespace" flag on the "put"
# command.
# namespace = "default"

# The items map is the only strictly required part of a variable
# specification, since path and namespace can be set via other means.
# It contains the sensitive material to encrypt and store as a Nomad
# variable. The entire items map is encrypted and decrypted as a
# single unit.

# REMINDER: While keys in the items map can contain dots, using them
# in templates is easier when they do not. As a best practice, avoid
# dotted keys when possible.
items {
  key1 = "value 1"
  key2 = "value 2"
}
