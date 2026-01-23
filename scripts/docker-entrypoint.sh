#!/usr/bin/env ash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


case "$1" in
  "agent" )
    if [[ -z "${NOMAD_SKIP_DOCKER_IMAGE_WARN}" && -z "${WONTON_SKIP_DOCKER_IMAGE_WARN}" ]]
    then
      echo "====================================================================================="
      echo "!! Running OpenWonton clients inside Docker containers is not supported.           !!"
      echo "!! Refer to https://openwonton.io/s/nomad-in-docker for more information.    !!"
      echo "!! Set WONTON_SKIP_DOCKER_IMAGE_WARN (or NOMAD_SKIP_DOCKER_IMAGE_WARN) to skip.     !!"
      echo "====================================================================================="
      echo ""
      sleep 2
    fi
esac

exec wonton "$@"
