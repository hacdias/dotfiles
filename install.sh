#/usr/bin/env bash

set -euo pipefail

function yn () {
  while true; do
    read -p "🙋 $1 [y/n] " yn
    case $yn in
      [Yy]* ) return 0;;
      [Nn]* ) return 1;;
      * ) echo "Please answer yes or no.";;
    esac
  done
}

if yn "Install homebrew?"; then
  # Install brew if not installed already
  # Homepage: https://brew.sh/
  if ! command -v brew; then
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  fi
fi

if yn "Install CLIs and App Store apps?"; then
  # Install homebrew packages and GUI apps
  ./brew.sh
  ./mas.sh
fi

if yn "Install oh-my-zsh and make zsh the default shell?"; then
  ./oh-my-zsh.sh
  # Make zsh the default shell
  sudo sh -c 'echo "/opt/homebrew/bin/zsh" >> /etc/shells'
  chsh -s /opt/homebrew/bin/zsh
fi

if yn "Sync dotfiles?"; then
  ./dotfiles.sh
fi

if yn "Install other tools and scripts?"; then
  ./scripts.sh
fi

if yn "Install node and nvm?"; then
  ./node.sh
fi
