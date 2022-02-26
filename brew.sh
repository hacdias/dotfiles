#/usr/bin/env bash

set -euo pipefail

brew update
brew upgrade

# Install multiple tools
brew bundle install --file=~/Brewfile

# Cleanup Cellar
brew cleanup
