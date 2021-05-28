alias ll='ls -lh'
alias la='ls -lAh'
alias l='ls -CF'

alias open-ports="lsof -Pan -iTCP -sTCP:LISTEN"
alias ip="dig +short myip.opendns.com @resolver1.opendns.com"
alias fresh-install="rm -r node_modules && rm package-lock.json && npm install"
alias ftypes="find . -type f | awk -F'.' '{print \$NF}' | sort| uniq -c | sort -g"

# Testground
alias tg='testground'
alias tgd='testground daemon'
alias tgr='testground run'
alias tgb='testground build'

# Uni Stuff
alias cdmsce='cd "~/Google Drive/MSCE"'

alias swap_yubikey="killall gpg-agent && rm -rf ~/.gnupg/private-keys-v1.d/ && gpg --card-status"