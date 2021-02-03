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
brew install pyenv

# Shell and prompt (https://starship.rs)
brew install zsh
brew install starship

# Cleanup Cellar
brew cleanup
