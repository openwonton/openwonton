# OpenWonton Documentation Website

This directory contains the Docusaurus site that renders the OpenWonton docs.

## Running the site locally

- `npm install`
- `npm start`

The site will be available at `http://localhost:3000`. If you are using a
non-default base URL, set `DOCUSAURUS_BASE_URL` before starting the dev server.

## Editing content

Documentation content lives under `content/` in MDX format. The top-level
sections map to URL prefixes:

- `content/intro` → `/intro`
- `content/docs` → `/docs`
- `content/api-docs` → `/api-docs`
- `content/tools` → `/tools`
- `content/plugins` → `/plugins`

Partials referenced via `@include` are stored in `content/partials`.

## Sidebars and navigation

Sidebar structure is defined in `data/*-nav-data.json`. The Docusaurus
`sidebars.js` file converts these JSON files into Docusaurus sidebars at build
runtime.

## Redirects

Redirects are handled with `@docusaurus/plugin-client-redirects` in
`docusaurus.config.js`.

## Building for production

- `npm run build`

The static output is generated in `build/`.

## Deployment

The site is a static build. Configure your hosting provider to run
`npm run build` and serve the `build/` directory.

### Bunny.net (CI/CD)

This repo includes a GitHub Actions workflow to build and deploy the site to a
Bunny Storage Zone, and optionally purge the Pull Zone cache.

Required GitHub secrets:

- `BUNNY_STORAGE_ZONE` - Storage Zone name
- `BUNNY_STORAGE_PASSWORD` - Storage Zone password
- `BUNNY_STORAGE_ENDPOINT` - Storage endpoint host (e.g. `ny.storage.bunnycdn.com`)
- `BUNNY_API_KEY` - Bunny API key (for Pull Zone cache purge)
- `BUNNY_PULL_ZONE_ID` - Pull Zone ID to purge after deploy

The workflow publishes the contents of `website/build` to the root of the
Storage Zone.
