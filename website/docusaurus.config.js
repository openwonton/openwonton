// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

const path = require('path');

const baseUrl = process.env.DOCUSAURUS_BASE_URL || '/';

module.exports = {
  title: 'OpenWonton',
  tagline: 'Nomad-compatible workload orchestration',
  url: 'https://openwonton.org',
  baseUrl,
  trailingSlash: false,
  onBrokenLinks: 'warn',
  onBrokenMarkdownLinks: 'warn',
  favicon: '_favicon.ico',
  organizationName: 'openwonton',
  projectName: 'openwonton',
  staticDirectories: ['static', 'public'],
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },
  markdown: {
    mermaid: true,
  },
  themes: ['@docusaurus/theme-mermaid'],
  presets: [
    [
      'classic',
      {
        docs: {
          path: 'content',
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          exclude: ['partials/**', 'security.mdx'],
          remarkPlugins: [
            [
              require('./remark/include-markdown'),
              { rootDir: path.resolve(__dirname, 'content', 'partials') },
            ],
            require('./remark/paragraph-alerts'),
            require('./remark/anchor-aliases'),
            require('./remark/rewrite-legacy-links'),
          ],
        },
        blog: false,
        pages: {},
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
  plugins: [
    [
      '@docusaurus/plugin-client-redirects',
      {
        redirects: [
          {
            from: '/tools/autoscaling/concepts/checks',
            to: '/tools/autoscaling/concepts/policy-eval/checks',
          },
          {
            from: '/tools/autoscaling/concepts/node-selector-strategy',
            to: '/tools/autoscaling/concepts/policy-eval/node-selector-strategy',
          },
          {
            from: '/docs/enterprise',
            to: '/docs',
          },
          {
            from: '/docs/enterprise/:path*',
            to: '/docs',
          },
        ],
      },
    ],
    [
      'docusaurus-plugin-plausible',
      {
        domain: 'openwonton.org',
      },
    ],
  ],
  themeConfig: {
    image: 'img/og-image.png',
    mermaid: {
      theme: { light: 'base', dark: 'base' },
      options: {
        themeVariables: {
          darkMode: true,
          background: '#0a0f15',
          primaryColor: '#121a24',
          primaryTextColor: '#e4ecf4',
          primaryBorderColor: '#f6a23a',
          lineColor: '#f6a23a',
          secondaryColor: '#16212d',
          tertiaryColor: '#0a0f15',
          noteBkgColor: '#0d1722',
          noteTextColor: '#e4ecf4',
          clusterBkg: '#0b151f',
          clusterBorder: '#243344',
          edgeLabelBackground: '#0b151f',
          archEdgeColor: '#f6a23a',
          archEdgeArrowColor: '#f6a23a',
          archEdgeWidth: '2px',
          archGroupBorderColor: '#243344',
          archGroupBorderWidth: '1.5px',
          fontFamily: "Sora, 'Avenir Next', 'Segoe UI', sans-serif",
        },
      },
    },
    colorMode: {
      defaultMode: 'dark',
      disableSwitch: true,
      respectPrefersColorScheme: false,
    },
    navbar: {
      title: 'OpenWonton',
      logo: {
        alt: 'OpenWonton',
        src: 'img/logo_openwonton.png',
      },
      items: [
        { type: 'docSidebar', sidebarId: 'intro', label: 'Intro', position: 'left' },
        { type: 'docSidebar', sidebarId: 'docs', label: 'Docs', position: 'left' },
        { type: 'docSidebar', sidebarId: 'apiDocs', label: 'API', position: 'left' },
        { type: 'docSidebar', sidebarId: 'tools', label: 'Tools', position: 'left' },
        { type: 'docSidebar', sidebarId: 'plugins', label: 'Plugins', position: 'left' },
        { to: '/security', label: 'Security', position: 'left' },
        {
          href: 'https://github.com/openwonton/openwonton',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            { label: 'Intro', to: '/intro' },
            { label: 'Docs', to: '/docs' },
            { label: 'API', to: '/api-docs' },
            { label: 'Legal', to: '/legal' },
          ],
        },
        {
          title: 'Community',
          items: [{ label: 'GitHub', href: 'https://github.com/openwonton/openwonton' }],
        },
      ],
      copyright: `Content licensed under MPL-2.0. Portions © HashiCorp, Inc.`,
    },
    prism: {
      additionalLanguages: ['hcl', 'ini'],
    },
  },
};
