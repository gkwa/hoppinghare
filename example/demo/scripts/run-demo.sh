#!/usr/bin/env bash
set -e

# Hopping Hare version
hoppinghare version

# List available templates
hoppinghare list-templates --template-dir ./templates

# Generate project with command line variables
hoppinghare generate \
  --template-url ./templates/go-api \
  --output-folder ./output/demo1 \
  --var ProjectName=MyGoAPI \
  --var GoVersion=1.24.2 \
  --non-interactive

# Generate project with variable file
hoppinghare generate \
  --template-url ./templates/go-api \
  --output-folder ./output/demo2 \
  --var-file ./templates/vars.yml \
  --non-interactive

# Generate with variable file and command line override
hoppinghare generate \
  --template-url ./templates/go-api \
  --output-folder ./output/demo3 \
  --var-file ./templates/vars.yml \
  --var ProjectName=OverriddenAPI \
  --non-interactive

# Generate project without tests
hoppinghare generate \
  --template-url ./templates/go-api \
  --output-folder ./output/demo4 \
  --var ProjectName=NoTestsAPI \
  --var IncludeTests=false \
  --non-interactive

[ -f ./output/demo4/main_test.go ] && echo "Test file exists" || echo "Test file does not exist"

# Generate with verbose logging
hoppinghare generate \
  --template-url ./templates/go-api \
  --output-folder ./output/demo5 \
  --var-file ./templates/vars.yml \
  --var ProjectName=VerboseAPI \
  --non-interactive \
  -v

# Show generated files
find ./output -type f
