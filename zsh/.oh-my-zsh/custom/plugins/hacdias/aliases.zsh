alias ll='ls -lh'
alias la='ls -lAh'
alias l='ls -CF'

alias open_ports="lsof -Pan -iTCP -sTCP:LISTEN"
alias ip="dig +short myip.opendns.com @resolver1.opendns.com"
alias ftypes="find . -type f | awk -F'.' '{print \$NF}' | sort| uniq -c | sort -g"
alias swap_yubikey="killall gpg-agent && rm -rf ~/.gnupg/private-keys-v1.d/ && gpg --card-status"
alias tailscale="/Applications/Tailscale.app/Contents/MacOS/Tailscale"
alias delete_empty='find . -type d -empty -delete'
alias cat='bat --paging=never'
