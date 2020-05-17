export ZSH=$HOME/.oh-my-zsh
export PATH=$HOME/.scripts:$PATH

export LC_CTYPE=en_US.UTF-8
export LC_ALL=en_US.UTF-8
export EDITOR=`which vim`

export GPG_TTY=$(tty)

# Ensure we use gpg-agent instead of ssh-agent for SSH!
export SSH_AUTH_SOCK=$(gpgconf --list-dirs agent-ssh-socket)
gpgconf --launch gpg-agent

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

# THIS IS HUGELY SLOW. DONT LIKE IT.
export NVM_DIR="$HOME/.nvm"
[ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"  # This loads nvm
[ -s "/usr/local/opt/nvm/etc/bash_completion.d/nvm" ] && . "/usr/local/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion
