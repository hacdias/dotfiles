alias ll='ls -lh'
alias la='ls -lAh'
alias l='ls -CF'

alias open_ports="lsof -Pan -iTCP -sTCP:LISTEN"
alias ip="dig +short myip.opendns.com @resolver1.opendns.com"
alias ftypes="find . -type f | awk -F'.' '{print \$NF}' | sort| uniq -c | sort -g"
alias swap_yubikey="killall gpg-agent && rm -rf ~/.gnupg/private-keys-v1.d/ && gpg --card-status"
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
