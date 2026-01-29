// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

const { visit } = require('unist-util-visit');

const ALIAS_REGEX = /\s*\(\(#([^)]+)\)\)\s*$/;

function parseAliases(raw) {
  return raw
    .split(',')
    .map((entry) => entry.trim().replace(/^#/, ''))
    .filter(Boolean);
}

function extractAliases(node) {
  if (!node.children || !node.children.length) {
    return null;
  }

  for (let i = node.children.length - 1; i >= 0; i -= 1) {
    const child = node.children[i];
    if (child.type !== 'text') {
      continue;
    }

    const match = child.value.match(ALIAS_REGEX);
    if (!match) {
      continue;
    }

    const aliases = parseAliases(match[1]);
    child.value = child.value.replace(ALIAS_REGEX, '').trimEnd();
    if (child.value === '') {
      node.children.splice(i, 1);
    }

    return aliases.length ? aliases : null;
  }

  return null;
}

function buildAnchorSpan(alias) {
  return {
    type: 'mdxJsxTextElement',
    name: 'span',
    attributes: [
      {
        type: 'mdxJsxAttribute',
        name: 'id',
        value: alias,
      },
    ],
    children: [],
  };
}

module.exports = function anchorAliases() {
  return function transformer(tree) {
    visit(tree, ['heading', 'paragraph'], (node) => {
      const aliases = extractAliases(node);
      if (!aliases || !aliases.length) {
        return;
      }

      if (node.type === 'heading') {
        node.data = node.data || {};
        node.data.hProperties = node.data.hProperties || {};
        if (!node.data.hProperties.id) {
          node.data.hProperties.id = aliases[0];
        }

        const extraAliases = aliases.slice(1);
        if (extraAliases.length) {
          node.children.unshift(...extraAliases.map(buildAnchorSpan));
        }
        return;
      }

      node.children.unshift(...aliases.map(buildAnchorSpan));
    });
  };
};
