export LANG="en_US.UTF-8"
export LC_ALL=${LANG}
export EDITOR=`which vim`

# PATH and PATH-like configurations
export GOPATH="$HOME/go"

export PATH="$HOME/bin:$PATH"
export PATH="$PATH:$GOPATH/bin"
export PATH="$HOME/.cargo/bin:$PATH"

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
  hacdias
)

source ${ZSH}/oh-my-zsh.sh

# Load Starship https://starship.rs/
eval "$(starship init zsh)"

# Load up my tokens
set -a
source ~/.tokens
set +a

# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/opt/homebrew/Caskroom/miniforge/base/bin/conda' 'shell.zsh' 'hook' 2> /dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/opt/homebrew/Caskroom/miniforge/base/etc/profile.d/conda.sh" ]; then
        . "/opt/homebrew/Caskroom/miniforge/base/etc/profile.d/conda.sh"
    else
        export PATH="/opt/homebrew/Caskroom/miniforge/base/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<

