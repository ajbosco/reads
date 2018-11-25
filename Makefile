.EXPORT_ALL_VARIABLES:
NAME := reads
PKG := github.com/ajbosco/reads/cmd/reads
BUILD_DIR := $(shell pwd)/build
TARGET := ${BUILD_DIR}/${NAME}
VERSION := $(shell cat VERSION.txt)
LDFLAGS ?= -X github.com/ajbosco/reads/version.VERSION=${VERSION}

# List the GOOS and GOARCH to build
GOOSARCHES = $(shell cat .goosarch)

.PHONY: fmt
fmt: ## Verifies all files have been `gofmt`ed.
	@gofmt -s -l .

.PHONY: lint
lint: ## Verifies `golint` passes.
	@golint ./...

.PHONY: test
test: ## Runs the go tests.
	@go test -cover -race ./...

.PHONY: vet
vet: ## Verifies `go vet` passes.
	@go vet ./...

.PHONY: build-native
build-native: ## run go build for current OS
	@go build -ldflags "$(LDFLAGS)" -o "${TARGET}" ${PKG}

define buildrelease
GOOS=$(1) GOARCH=$(2) go build \
	 -ldflags "$(LDFLAGS)" \
	 -o $(BUILD_DIR)/$(NAME)-$(1)-$(2) ${PKG};
md5sum $(BUILD_DIR)/$(NAME)-$(1)-$(2) > $(BUILD_DIR)/$(NAME)-$(1)-$(2).md5;
sha256sum $(BUILD_DIR)/$(NAME)-$(1)-$(2) > $(BUILD_DIR)/$(NAME)-$(1)-$(2).sha256;
endef

.PHONY: release
release: VERSION.txt ## Builds the cross-compiled binaries, naming them in such a way for release (eg. binary-GOOS-GOARCH).
	@$(foreach GOOSARCH,$(GOOSARCHES), $(call buildrelease,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH))))

.PHONY: bump-version
BUMP := patch
bump-version: ## Bump the version in the version file. Set BUMP to [ patch | major | minor ].
	@go get -u github.com/jessfraz/junk/sembump # update sembump tool
	$(eval NEW_VERSION = $(shell sembump --kind $(BUMP) $(VERSION)))
	@echo "Bumping VERSION.txt from $(VERSION) to $(NEW_VERSION)"
	echo $(NEW_VERSION) > VERSION.txt
	@echo "Updating links to download binaries in README.md"
	sed -i s/$(VERSION)/$(NEW_VERSION)/g README.md
	git add VERSION.txt README.md
	git commit -vsam "Bump version to $(NEW_VERSION)"
	@echo "Run make tag to create and push the tag for new version $(NEW_VERSION)"

.PHONY: tag
tag: ## Create a new git tag to prepare to build a release.
	git tag -sa $(VERSION) -m "$(VERSION)"
	@echo "Run git push origin $(VERSION) to push your new tag to GitHub and trigger a travis build."

.PHONY: clean
clean: ## Cleanup any build binaries or packages.
	$(RM) -r $(BUILD_DIR)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
