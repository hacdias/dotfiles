#/usr/bin/env bash

set -euo pipefail

# Wrap around stow to target the user home.
function stowu () {
  stow -t ~ $@
}

# On new installations, .zshrc and .zprofile should be removed
# otherwise stow will complain.
rm -rf ~/.zshrc ~/.zprofile

# Make sure the tokens file exists.
touch ~/.tokens

# Stow the dotfiles.
stowu git
stowu gpg
stowu starship
stowu nvim
stowu zsh
stowu vscode
stowu linearmouse
stowu claude
stowu rex

echo "Please run:\n\texec zsh -l"
