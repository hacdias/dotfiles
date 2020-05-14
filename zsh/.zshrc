export ZSH=$HOME/.oh-my-zsh
export PATH=$HOME/.scripts:$PATH

export LC_CTYPE=en_US.UTF-8
export LC_ALL=en_US.UTF-8
export EDITOR=`which vim`

# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

ZSH_THEME="funny"

plugins=(
  git
  hacdias
)

export PATH=$PATH:$HOME/go/bin
export PATH="$PATH:$HOME/.npm-global/bin"

# export NODE_PATH="$NODE_PATH:$HOME/.npm-global/lib/node_modules"
source $ZSH/oh-my-zsh.sh
source $HOME/.cargo/env

export NVM_DIR="$HOME/.nvm"
[ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"  # This loads nvm
[ -s "/usr/local/opt/nvm/etc/bash_completion.d/nvm" ] && . "/usr/local/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion
