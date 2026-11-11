eval "$(/opt/homebrew/bin/brew shellenv)"

export LANG="en_US.UTF-8"
export LC_ALL=${LANG}
export EDITOR=nvim
export BAT_THEME_DARK="Catppuccin Mocha"
export BAT_THEME_LIGHT="Catppuccin Latte"
export MANPAGER="sh -c 'col -bx | bat -l man -p'"
export PNPM_HOME="$HOME/Library/pnpm"

# PATH and PATH-like configurations
export GOPATH="$HOME/go"
export PATH="$HOME/bin:$PATH"
export PATH="$PATH:$GOPATH/bin"
export PATH="$PATH:$PNPM_HOME"

# Load up my tokens
set -a
[ -f ~/.tokens ] && source ~/.tokens
set +a
