export LANG="en_US.UTF-8"
export LC_ALL=${LANG}
export EDITOR=`which vim`
export BAT_THEME="Monokai Extended Bright"
export MANPAGER="sh -c 'col -bx | bat -l man -p'"

# PATH and PATH-like configurations
export GOPATH="$HOME/go"

export PATH="$HOME/bin:$PATH"
export PATH="$PATH:$GOPATH/bin"
export PATH="$HOME/.cargo/bin:$PATH"
export PATH="$HOME/Library/Python/3.10/bin:$PATH"

# Configure GPG
export GPG_TTY=$(tty)

# Start GPG over SSH
# export SSH_AUTH_SOCK=$(gpgconf --list-dirs agent-ssh-socket)
# gpgconf --launch gpg-agent

# Start SSH agent
eval "$(ssh-agent -s)" &>/dev/null

# Configure Oh My Zsh https://ohmyz.sh/
export ZSH="$HOME/.oh-my-zsh"
ZSH_THEME="" # Do not load any zsh theme as I'm using starship!

plugins=(
  zsh-autosuggestions
  zsh-syntax-highlighting
  starship
  hacdias
  git
)

setopt HIST_IGNORE_SPACE

[ -f  ${ZSH}/oh-my-zsh.sh ] && source ${ZSH}/oh-my-zsh.sh
[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

# Load up my tokens
set -a
source ~/.tokens
set +a
