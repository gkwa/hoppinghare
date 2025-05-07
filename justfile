# List available recipes (default)
default:
    @just --list

# Run the demo in the examples/demo directory
demo:
    @cd examples/demo && just demo

# Set up the demo environment
setup:
    @cd examples/demo && just setup

# Clean up after the demo
teardown:
    @cd examples/demo && just teardown

# Build the hoppinghare binary
build:
    go build -o bin/hoppinghare
