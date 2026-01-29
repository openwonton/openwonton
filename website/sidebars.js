// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

const fs = require('fs');
const path = require('path');

const contentRoot = path.join(__dirname, 'content');

function normalizeRoute(route) {
  return route.replace(/^\/+/, '').replace(/\/$/, '');
}

function resolveDocId(baseDir, route) {
  const normalized = normalizeRoute(route);
  const relative = path.posix.join(baseDir, normalized);
  const candidates = [
    `${relative}.mdx`,
    `${relative}.md`,
    path.posix.join(relative, 'index.mdx'),
    path.posix.join(relative, 'index.md'),
  ];

  for (const candidate of candidates) {
    const fullPath = path.join(contentRoot, candidate);
    if (fs.existsSync(fullPath)) {
      return candidate.replace(/\.(mdx|md)$/, '');
    }
  }

  // Fall back to the original path if the file is missing.
  // Docusaurus will surface missing doc IDs at build time.
  return relative;
}

function mapNavItems(items, baseDir) {
  return items.map((item) => {
    if (item.routes) {
      return {
        type: 'category',
        label: item.title,
        items: mapNavItems(item.routes, baseDir),
      };
    }

    if (item.path) {
      return {
        type: 'doc',
        id: resolveDocId(baseDir, item.path),
        label: item.title,
      };
    }

    if (item.href) {
      return {
        type: 'link',
        label: item.title,
        href: item.href,
      };
    }

    return null;
  }).filter(Boolean);
}

function buildSidebar(navFile, baseDir) {
  const navPath = path.join(__dirname, 'data', navFile);
  const navItems = JSON.parse(fs.readFileSync(navPath, 'utf8'));
  return mapNavItems(navItems, baseDir);
}

module.exports = {
  intro: buildSidebar('intro-nav-data.json', 'intro'),
  docs: buildSidebar('docs-nav-data.json', 'docs'),
  tools: buildSidebar('tools-nav-data.json', 'tools'),
  plugins: buildSidebar('plugins-nav-data.json', 'plugins'),
  apiDocs: buildSidebar('api-docs-nav-data.json', 'api-docs'),
};
