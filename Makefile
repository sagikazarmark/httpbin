GO?=go
GOFMT?=gofmt
GLIDE:=$(shell if which glide > /dev/null 2>&1; then echo "glide"; fi)
BINDATA?=go-bindata
GO_RUN_FILES=$(shell find . -type f -name "*.go" -not -name "*_test.go" -not -path "./vendor/*")
GO_SOURCE_FILES=$(shell find . -type f -name "*.go" -not -name "bindata.go" -not -path "./vendor/*")
GO_PACKAGES=$(shell go list ./... | grep -v /vendor/)

# Build the project, optionally in form of a binary
build:
ifdef BINARY
	$(GO) build $(BUILDOPTS) -o $(BINARY)
else
	$(GO) install $(BUILDOPTS) .
endif

# Install dependencies, optionally using go get
install:
ifdef GLIDE
	@$(GLIDE) install
else ifeq ($(FORCE), true)
	$(GO) get
else
	@echo "Glide is necessary for installing project dependencies: http://glide.sh/ Run this command with FORCE=true to fall back to go get" 1>&2 && exit 1
endif

# Clean Go environment
# TODO: add BINARY support?
clean:
	@$(GO) clean

# Run all sources if this is a console app
run:
	@$(GO) run $(GO_RUN_FILES)

# Run tests
test:
ifeq ($(VERBOSE), true)
	@$(GO) test -v $(GO_PACKAGES)
else
	@$(GO) test $(GO_PACKAGES)
endif

# Generate necessary files
generate:
	@$(GO) generate

# Check that all source files follow the Coding Style
check:
	@$(GOFMT) -l $(GO_SOURCE_FILES) | read && echo "Code differs from gofmt's style" 1>&2 && exit 1 || true
	@$(GO) vet $(GO_PACKAGES)

# Fix Coding Standard violations
fix:
ifeq ($(SIMPLYFY), true)
	@$(GOFMT) -l -w -s $(GO_SOURCE_FILES)
else
	@$(GOFMT) -l -w $(GO_SOURCE_FILES)
endif

.PHONY: build install clean run test generate check fix
