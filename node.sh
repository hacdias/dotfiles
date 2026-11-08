#/usr/bin/env zsh

set -euo pipefail

# fnm (https://github.com/Schniz/fnm) is installed via mac/Brewfile.

# Install the latest LTS Node.js version (the first install becomes the default)
fnm install --lts
eval "$(fnm env)"

# Always useful packages: ncu and corepack
npm install -g npm-check-updates
npm install -g corepack
