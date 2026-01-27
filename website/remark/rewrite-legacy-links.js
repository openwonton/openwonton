// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

const { visit } = require('unist-util-visit');

const UPSTREAM_HOST = 'https://developer.hashicorp.com';

const PRODUCT_PREFIXES = [
  { from: '/nomad/consul/', to: `${UPSTREAM_HOST}/consul/` },
  { from: '/consul/', to: `${UPSTREAM_HOST}/consul/` },
  { from: '/nomad/vault/', to: `${UPSTREAM_HOST}/vault/` },
  { from: '/vault/', to: `${UPSTREAM_HOST}/vault/` },
  { from: '/nomad/waypoint/', to: `${UPSTREAM_HOST}/waypoint/` },
  { from: '/waypoint/', to: `${UPSTREAM_HOST}/waypoint/` },
  { from: '/nomad/vagrant/', to: `${UPSTREAM_HOST}/vagrant/` },
  { from: '/vagrant/', to: `${UPSTREAM_HOST}/vagrant/` },
  { from: '/nomad/terraform/', to: `${UPSTREAM_HOST}/terraform/` },
  { from: '/terraform/', to: `${UPSTREAM_HOST}/terraform/` },
];

const NOMAD_UPSTREAM_PREFIXES = [
  '/nomad/tutorials/',
  '/nomad/docs/enterprise/',
  '/nomad/docs/v',
  '/nomad/downloads',
];

function rewriteLegacyUrl(url) {
  if (!url || !url.startsWith('/')) {
    return url;
  }

  for (const { from, to } of PRODUCT_PREFIXES) {
    if (url.startsWith(from)) {
      return `${to}${url.slice(from.length)}`;
    }
  }

  for (const prefix of NOMAD_UPSTREAM_PREFIXES) {
    if (url.startsWith(prefix)) {
      return `${UPSTREAM_HOST}${url}`;
    }
  }

  return url;
}

module.exports = function rewriteLegacyLinks() {
  return (tree) => {
    visit(tree, ['link', 'definition'], (node) => {
      node.url = rewriteLegacyUrl(node.url);
    });
  };
};
