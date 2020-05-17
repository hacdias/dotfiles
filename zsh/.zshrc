export LANG="en_US.UTF-8"
export LC_ALL=${LANG}
export EDITOR=`which vim`

# PATH and PATH-like configurations
export GOPATH=${HOME}/go
export PATH=${HOME}/bin:${PATH}
export PATH=${PATH}:${GOPATH}/bin

# Configure GPG and SSH over GPG
export GPG_TTY=$(tty)
export SSH_AUTH_SOCK=$(gpgconf --list-dirs agent-ssh-socket)
gpgconf --launch gpg-agent

# Configure Oh My Zsh https://ohmyz.sh/
export ZSH="${HOME}/.oh-my-zsh"
ZSH_THEME="" # Do not load any zsh theme as I'm using starship!

plugins=(
  zsh-autosuggestions
  zsh-syntax-highlighting
  hacdias
)

source ${ZSH}/oh-my-zsh.sh

# Load Starship https://starship.rs/
eval "$(starship init zsh)"
