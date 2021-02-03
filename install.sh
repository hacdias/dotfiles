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
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
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
  sudo sh -c 'echo "/usr/local/bin/zsh" >> /etc/shells'
  chsh -s /usr/local/bin/zsh
fi

if yn "Sync dotfiles?"; then
  ./dotfiles.sh
fi

if yn "Install other tools and scripts?"; then
  # TODO

  # Install some Go tools
  go install github.com/StackExchange/dnscontrol
fi

if yn "Install node and nvm?"; then
  ./node.sh
fi

if yn "Install testground tools?"; then
  ./brew-testgrond.sh
fi
