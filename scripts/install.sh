#!/bin/bash

binaryPath="/usr/local/bin/addigy"
githubLinkIntel="https://github.com/dsrosen6/addigy-cli/releases/latest/download/addigy-cli-intel"
githubLinkAppleSilicon="https://github.com/dsrosen6/addigy-cli/releases/latest/download/addigy-cli-apple-silicon"

# Check if Mac is running on Apple Silicon
if [[ $(uname -m) == "arm64" ]]; then
  link="$githubLinkAppleSilicon"
else
  link="$githubLinkIntel"
fi

# Download the binary
curl -L -o "$binaryPath" "$link"

# Make the binary executable
chmod +x "$binaryPath"