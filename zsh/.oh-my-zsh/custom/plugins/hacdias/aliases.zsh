alias ll='ls -lh'
alias la='ls -lAh'
alias l='ls -CF'

alias open-ports="lsof -Pan -iTCP -sTCP:LISTEN"
alias ip="dig +short myip.opendns.com @resolver1.opendns.com"
alias fresh-install="rm -r node_modules && rm package-lock.json && npm install"

# Testground
alias tg='testground'
alias tgd='testground daemon'
alias tgr='testground run'
alias tgb='testground build'
