export ZSH=$HOME/.oh-my-zsh

# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

ZSH_THEME="funny"

plugins=(
  git
)

source $ZSH/oh-my-zsh.sh
source .functions

alias ll='ls -alF'
alias la='ls -A'
alias l='ls -CF'
