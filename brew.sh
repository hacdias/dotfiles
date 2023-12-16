#/usr/bin/env bash

set -euo pipefail

brew update
brew upgrade

# Install multiple tools
brew bundle install --file=mac/Brewfile

# Cleanup Cellar
brew cleanup
