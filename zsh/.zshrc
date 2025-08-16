export LANG="en_US.UTF-8"
export LC_ALL=${LANG}
export EDITOR=`which vim`
export BAT_THEME="Monokai Extended Bright"
export MANPAGER="sh -c 'col -bx | bat -l man -p'"
export PNPM_HOME="$HOME/Library/pnpm"

# PATH and PATH-like configurations
export GOPATH="$HOME/go"
export PATH="$HOME/bin:$PATH"
export PATH="$PATH:$GOPATH/bin"
export PATH="$PATH:$PNPM_HOME"

# Android Studio
export ANDROID_HOME=$HOME/Library/Android/sdk
export PATH=$PATH:$ANDROID_HOME/emulator
export PATH=$PATH:$ANDROID_HOME/platform-tools
export JAVA_HOME=/Library/Java/JavaVirtualMachines/zulu-17.jdk/Contents/Home

# Configure GPG
export GPG_TTY=$(tty)

# Start SSH agent
eval "$(ssh-agent -s)" &>/dev/null

# Configure Oh My Zsh https://ohmyz.sh/
export ZSH="$HOME/.oh-my-zsh"
ZSH_THEME="" # Do not load any zsh theme as I'm using starship!
zstyle ':omz:plugins:nvm' lazy yes

plugins=(
  zsh-autosuggestions
  zsh-syntax-highlighting
  starship
  git
  nvm
  fzf
  hacdias
)

setopt HIST_IGNORE_SPACE

[ -f  ${ZSH}/oh-my-zsh.sh ] && source ${ZSH}/oh-my-zsh.sh

# Load up my tokens
set -a
[ -f ~/.tokens ] && source ~/.tokens
set +a
