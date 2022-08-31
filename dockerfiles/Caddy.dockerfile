FROM golang:1.19-buster

ENV CADDY_VERSION v2.5.2

RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest; \
  xcaddy build $CADDY_VERSION --with github.com/caddy-dns/cloudflare --with github.com/aksdb/caddy-cgi/v2; \
  mv caddy /usr/bin/caddy

FROM alpine:3.16

RUN apk add --no-cache ca-certificates mailcap perl

COPY --from=0 /usr/bin/caddy /usr/bin/caddy

RUN chmod +x /usr/bin/caddy; \
	caddy version

# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian grep '^hosts:' /etc/nsswitch.conf
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

# See https://caddyserver.com/docs/conventions#file-locations for details
ENV XDG_CONFIG_HOME /config
ENV XDG_DATA_HOME /data

VOLUME /config
VOLUME /data

EXPOSE 80
EXPOSE 443

WORKDIR /srv

CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]
