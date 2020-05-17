#/usr/bin/env zsh

set -euo pipefail

# Install brew if not installed already
# Homepage: https://brew.sh/
if ! command -v brew; then
	/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
fi

# Install homebrew packages and GUI apps
./brew.sh
./apps.sh

# Changes the default shell to ZSH
sudo sh -c 'echo "/usr/local/bin/zsh" >> /etc/shells'
chsh -s /usr/local/bin/zsh

# Instsll zsh and plugins
./zsh.sh

# Syncs some dotfiles
stow git
stow vim
stow vscode
stow gpg
stow starship
rm ~/.zshrc
stow zsh

source ~/.zshrc

# Install some Go tools
go install github.com/StackExchange/dnscontrol

# Install Node
./node.sh
