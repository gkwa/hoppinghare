# List available recipes (default)
default:
    @just --list

# Run the demo in the example/demo directory
demo:
    @cd example/demo && just demo

# Set up the demo environment
setup:
    @cd example/demo && just setup

# Clean up after the demo
teardown:
    @cd example/demo && just teardown

# Build the hoppinghare binary
build:
    go build -o bin/hoppinghare
