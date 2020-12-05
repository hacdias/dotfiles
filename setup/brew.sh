#/usr/bin/env bash

set -euo pipefail

brew update
brew upgrade

# Install multiple tools
brew install coreutils
brew install moreutils
brew install findutils
brew install gpg
brew install openssh
brew install git
brew install git-lfs
brew install ffmpeg
brew install wget
brew install mas
brew install go
brew install hugo
brew install ufraw
brew install imagemagick
brew install caddy
brew install watch
brew install bat
brew install bandwhich
brew install wakeonlan
brew install wireguard-tools

# Shell and prompt
brew install zsh
brew install starship # https://starship.rs

# Cleanup Cellar
brew cleanup

# Testground tooling
brew install helm
brew install kops
brew install terraform
echo "Install Amazon CLI: https://aws.amazon.com/cli/"
