#/usr/bin/env bash

set -euo pipefail

# Wrap around stow to target the user home.
function stowu () {
  stow -t ~ $@
}

# On new installations, .zshrc should be removed
# otherwise stow will complain.
rm -rf ~/.zshrc

# Make sure the tokens file exists.
touch ~/.tokens

# Stow the dotfiles.
stowu git
stowu gpg
stowu starship
stowu vim
stowu zsh

echo "Please run:\n\tsource ~/.zshrc"
