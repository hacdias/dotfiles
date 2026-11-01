alias ll='ls -lh'
alias la='ls -lAh'
alias l='ls -CF'

# Normalize `open` across Linux, macOS, and Windows.
if [ ! $(uname -s) = 'Darwin' ]; then
	if grep -q Microsoft /proc/version; then
		# Ubuntu on Windows using the Linux subsystem
		alias open='explorer.exe';
	else
		alias open='xdg-open';
	fi
fi

alias open_ports="lsof -Pan -iTCP -sTCP:LISTEN"
alias ip="dig +short myip.opendns.com @resolver1.opendns.com"
alias ftypes="find . -type f | awk -F'.' '{print \$NF}' | sort| uniq -c | sort -g"
alias delete_empty='find . -type d -empty -delete'
alias tidy_all='gfind -type f -name "go.mod" | while read p; do cd $(dirname $p); echo $p; go mod tidy; cd -; done'
alias clean_docker='docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q) && docker system prune -a && docker container ls && docker image ls'
alias clean_caches='go clean -modcache && nvm cache clear && brew cleanup && pnpm store prune && clean_docker'
alias pretty_csv='mlr --icsv --opprint cat'
alias reload='source ~/.zshrc'

# Named directories for cd ~{dir}
hash -d c=$HOME/Code
hash -d d=$HOME/Documents
