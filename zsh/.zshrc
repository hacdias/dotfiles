# Node version manager https://github.com/Schniz/fnm
eval "$(fnm env --use-on-cd --shell zsh)"

# Configure GPG
export GPG_TTY=$(tty)

# Configure Oh My Zsh https://ohmyz.sh/
export ZSH="$HOME/.oh-my-zsh"
ZSH_THEME="" # Do not load any zsh theme as I'm using starship!

plugins=(
  zsh-autosuggestions
  zsh-syntax-highlighting
  starship
  git
  fzf
  hacdias
)

setopt HIST_IGNORE_SPACE

[ -f  ${ZSH}/oh-my-zsh.sh ] && source ${ZSH}/oh-my-zsh.sh
