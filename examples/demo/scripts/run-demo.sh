#!/usr/bin/env bash
set -e

# Colors for pretty output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

# Function for section headers
section() {
  echo -e "\n${BLUE}=== $1 ===${NC}"
}

# Function to run commands with pretty output
run_cmd() {
  echo -e "${YELLOW}$ $1${NC}"
  eval "$1"
}

# Intro
section "Hopping Hare Demo"
echo "This demo shows how to use Hopping Hare to generate projects from templates."

# Show version
section "Checking version"
run_cmd "hoppinghare version"

# List available templates
section "Listing available templates"
run_cmd "hoppinghare list-templates --template-dir ./templates"

# Basic usage example
section "Generating project with command line variables"
run_cmd "hoppinghare generate --template-url ./templates/go-api --output-folder ./output/demo1 --var ProjectName=MyGoAPI --var GoVersion=1.24.2 --non-interactive"
run_cmd "ls -la ./output/demo1"
run_cmd "cat ./output/demo1/README.md"

# Using variable file
section "Generating project with variable file"
run_cmd "hoppinghare generate --template-url ./templates/go-api --output-folder ./output/demo2 --var-file ./templates/vars.yml --non-interactive"
run_cmd "ls -la ./output/demo2"
run_cmd "cat ./output/demo2/README.md"

# Override variable from file
section "Generating project with variable file and command line override"
run_cmd "hoppinghare generate --template-url ./templates/go-api --output-folder ./output/demo3 --var-file ./templates/vars.yml --var ProjectName=OverriddenAPI --non-interactive"
run_cmd "ls -la ./output/demo3"
run_cmd "cat ./output/demo3/README.md"

# Disable tests
section "Generating project without tests"
run_cmd "hoppinghare generate --template-url ./templates/go-api --output-folder ./output/demo4 --var ProjectName=NoTestsAPI --var IncludeTests=false --non-interactive"
run_cmd "ls -la ./output/demo4"
echo "Check if test file exists (should not exist):"
run_cmd "[ -f ./output/demo4/main_test.go ] && echo 'Test file exists' || echo 'Test file does not exist'"

# Verbose mode
section "Generating project with verbose logging"
run_cmd "hoppinghare generate --template-url ./templates/go-api --output-folder ./output/demo5 --var-file ./templates/vars.yml --var ProjectName=VerboseAPI --non-interactive -v"

# Success message
section "Demo completed successfully!"
echo "Generated projects are available in the ./output directory."
run_cmd "find ./output -type f | sort"
