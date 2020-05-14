#/usr/bin/env bash

set -euo pipefail

# Install brew if not installed already
# Homepage: https://brew.sh/
if ! command -v brew; then
	/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
fi

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
# brew install node
brew install nvm
brew install wget
brew install mas
brew install go
brew install hugo
brew install ufraw
brew install imagemagick

# Install ZSH and Oh My Zsh
# ZSH install instructions: https://github.com/ohmyzsh/ohmyzsh/wiki/Installing-ZSH
# Oh-My-Zsh install instructions: https://ohmyz.sh/#install
brew install zsh
chsh -s /bin/zsh # Changes the default shell

if [ ! -d ~/.oh-my-zsh ]; then
	sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
fi

# Cleanup Cellar
brew cleanup

# Install GUI apps from the App Store
mas lucky "1Password"
mas lucky "Spark"
mas lucky "Lungo"
mas lucky "Slack"
mas lucky "Telegram"
mas lucky "Affinity Photo"
# mas lucky "Microsoft PowerPoint"
# mas lucky "Microsoft Excel"
# mas lucky "Microsoft Word"

# Syncs some dotfiles
stow git
stow vim
stow vscode
stow zsh --adopt # force adoption because oh-my-zsh already introduces some

# Install some Go tools
go install github.com/StackExchange/dnscontrol

# Install some Node.js tools
nvm install node
nvm alias default node

npm install -g asciify \
	standard \
	npm-check-updates
