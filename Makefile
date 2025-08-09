# Variables
BINARY_NAME=countdown
SOURCE_FILES=*.go
BUILD_DIR=.
PREFIX=/usr/local
BINDIR=$(PREFIX)/bin
USER_BINDIR=$(HOME)/.local/bin

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) $(SOURCE_FILES)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Install binary to system
install: build
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 $(BINARY_NAME) $(DESTDIR)$(BINDIR)/

# Uninstall binary from system
uninstall:
	rm -f $(DESTDIR)$(BINDIR)/$(BINARY_NAME)

# Install binary to user directory
install-user: build
	install -d $(USER_BINDIR)
	install -m 755 $(BINARY_NAME) $(USER_BINDIR)/

# Uninstall binary from user directory
uninstall-user:
	rm -f $(USER_BINDIR)/$(BINARY_NAME)

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build the binary (default)"
	@echo "  release      - Build optimized binary for release"
	@echo "  clean        - Remove build artifacts"
	@echo "  install      - Install binary to $(BINDIR)"
	@echo "  uninstall    - Remove binary from $(BINDIR)"
	@echo "  install-user - Install binary to $(USER_BINDIR)"
	@echo "  uninstall-user - Remove binary from $(USER_BINDIR)"
	@echo "  help         - Show this help message"

.PHONY: all build release test clean install uninstall install-user uninstall-user fmt lint vet check dev help
