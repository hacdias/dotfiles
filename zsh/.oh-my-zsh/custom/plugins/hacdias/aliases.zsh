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
alias swap_yubikey="gpg-connect-agent "scd serialno" "learn --force" /bye"
alias tailscale="/Applications/Tailscale.app/Contents/MacOS/Tailscale"
alias delete_empty='find . -type d -empty -delete'
alias jaeger='docker run --rm -it --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one'
alias tidy_all='gfind -type f -name "go.mod" | while read p; do cd $(dirname $p); echo $p; go mod tidy; cd -; done'
alias clean_docker='docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q) && docker system prune -a && docker container ls && docker image ls'
alias pretty_csv='mlr --icsv --opprint cat'
alias reload='source ~/.zshrc'

# Named directories for cd ~{dir}
hash -d c=$HOME/Code
hash -d w=$HOME/Code/work
hash -d d=$HOME/Documents
