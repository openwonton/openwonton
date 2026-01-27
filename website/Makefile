.DEFAULT_GOAL := website

.PHONY: website
website:
	@echo "==> Installing dependencies..."
	@npm install
	@echo "==> Starting Docusaurus..."
	@npm start

.PHONY: website/build
website/build:
	@echo "==> Installing dependencies..."
	@npm install
	@echo "==> Building Docusaurus..."
	@npm run build
