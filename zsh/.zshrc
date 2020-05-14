export ZSH=$HOME/.oh-my-zsh
export PATH=$HOME/.scripts:$PATH

export LC_CTYPE=en_US.UTF-8
export LC_ALL=en_US.UTF-8
export EDITOR=`which code`

# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

ZSH_THEME="funny"

plugins=(
  git
  hacdias
)

source $ZSH/oh-my-zsh.sh
export PATH=$PATH:$HOME/go/bin
