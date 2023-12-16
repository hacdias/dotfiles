#/usr/bin/env zsh

set -euo pipefail

NVM_VERSION=0.39.7

# 1. Install NVM: https://github.com/nvm-sh/nvm
#    Homebrew NVM version is broken.
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v${NVM_VERSION}/install.sh | bash

# Just make sure everything's updated!
source ~/.zshrc

# Install the latest Node.js version and make it default
nvm install node
nvm alias default node

# NCU is always useful
npm install -g npm-check-updates
