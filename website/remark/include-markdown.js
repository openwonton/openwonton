// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

const fs = require('fs');
const path = require('path');
const { unified } = require('unified');
const remarkParse = require('remark-parse').default;
const remarkMdx = require('remark-mdx').default;
const remarkGfm = require('remark-gfm').default;
const { visit, SKIP } = require('unist-util-visit');

module.exports = function includeMarkdown(options = {}) {
  const rootDir =
    options.rootDir || path.join(process.cwd(), 'content', 'partials');

  const parser = unified().use(remarkParse).use(remarkMdx).use(remarkGfm);

  return function transformer(tree, file) {

    visit(tree, 'paragraph', (node, index, parent) => {
      if (!parent || typeof index !== 'number') {
        return;
      }

      if (!node.children || node.children.length !== 1) {
        return;
      }

      const first = node.children[0];
      if (!first || first.type !== 'text') {
        return;
      }

      const match = first.value.trim().match(/^@include\s+['"](.+?)['"]$/);
      if (!match) {
        return;
      }

      const includePath = match[1].replace(/^\/+/, '');
      const fullPath = path.join(rootDir, includePath);

      if (!fs.existsSync(fullPath)) {
        file.message(`Include not found: ${includePath}`, node);
        return;
      }

      const contents = fs.readFileSync(fullPath, 'utf8');
      const includeTree = parser.parse(contents);

      parent.children.splice(index, 1, ...includeTree.children);
      return [SKIP, index + includeTree.children.length];
    });
  };
};
