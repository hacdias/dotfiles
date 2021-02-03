#/usr/bin/env bash

set -euo pipefail

# Install ZSH and Oh My Zsh
# ZSH install instructions: https://github.com/ohmyzsh/ohmyzsh/wiki/Installing-ZSH
# Oh-My-Zsh install instructions: https://ohmyz.sh/#install
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# Install some plugins
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
