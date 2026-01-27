// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

const { visit } = require('unist-util-visit');
const { toString } = require('mdast-util-to-string');

const MARKER_REGEX = /^\s*(~>|->|=>|!>)\s*/;
const DEFAULT_MARKER_MAP = {
  '~>': 'Note',
  '->': 'Note',
  '=>': 'Tip',
  '!>': 'Warning',
};

function normalizeLabel(label) {
  return label.trim().replace(/[:!]+$/, '');
}

function resolveComponent(marker, label) {
  if (label) {
    const normalized = label.toLowerCase();
    if (normalized.startsWith('warning') || normalized.startsWith('be careful')) {
      return 'Warning';
    }
    if (normalized.startsWith('important') || normalized.startsWith('caution')) {
      return 'Caution';
    }
    if (normalized.startsWith('tip')) {
      return 'Tip';
    }
    if (normalized.startsWith('note')) {
      return 'Note';
    }
  }

  return DEFAULT_MARKER_MAP[marker] || 'Note';
}

function shouldSetTitle(component, label) {
  if (!label) {
    return false;
  }

  const normalized = label.toLowerCase();
  if (component === 'Note' && normalized === 'note') {
    return false;
  }
  if (component === 'Warning' && normalized === 'warning') {
    return false;
  }
  if (component === 'Tip' && normalized === 'tip') {
    return false;
  }
  if (component === 'Caution' && normalized === 'caution') {
    return false;
  }

  return true;
}

module.exports = function paragraphAlerts() {
  return function transformer(tree) {
    visit(tree, 'paragraph', (node, index, parent) => {
      if (!parent || typeof index !== 'number' || !node.children.length) {
        return;
      }

      const first = node.children[0];
      if (!first || first.type !== 'text') {
        return;
      }

      const markerMatch = first.value.match(MARKER_REGEX);
      if (!markerMatch) {
        return;
      }

      const marker = markerMatch[1];
      first.value = first.value.replace(MARKER_REGEX, '');
      if (first.value.trim() === '') {
        node.children.shift();
      }

      let label;
      if (node.children[0] && node.children[0].type === 'strong') {
        const rawLabel = normalizeLabel(toString(node.children[0]));
        label = rawLabel;
        node.children.shift();
        if (node.children[0] && node.children[0].type === 'text') {
          node.children[0].value = node.children[0].value.replace(/^\s+/, '');
        }
      }

      const component = resolveComponent(marker, label);
      const attributes = [];

      if (shouldSetTitle(component, label)) {
        attributes.push({
          type: 'mdxJsxAttribute',
          name: 'title',
          value: label,
        });
      }

      parent.children[index] = {
        type: 'mdxJsxFlowElement',
        name: component,
        attributes,
        children: node.children,
      };
    });
  };
};
