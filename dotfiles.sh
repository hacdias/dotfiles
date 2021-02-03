#/usr/bin/env bash

set -euo pipefail

# Wrap around stow to target the user home.
function stowu () {
  stow  -t ~ $@
}

# On new installations, .zshrc should be removed
# otherwise stow will complain.
rm -rf ~/.zshrc

# Stow the dotfiles.
stowu git
stowu gpg
stowu starship
stowu vim
stowu vscode
stowu zsh

echo "Please run:\n\tsource ~/.zshrc"
