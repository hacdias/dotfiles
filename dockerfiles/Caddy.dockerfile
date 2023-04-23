FROM golang:1.20-buster

ENV CADDY_VERSION v2.6.4

RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest; \
  xcaddy build $CADDY_VERSION --with github.com/caddy-dns/cloudflare; \
  mv caddy /usr/bin/caddy

FROM alpine:3.17

RUN apk add --no-cache ca-certificates mailcap

COPY --from=0 /usr/bin/caddy /usr/bin/caddy

RUN chmod +x /usr/bin/caddy; \
	caddy version

# See https://caddyserver.com/docs/conventions#file-locations for details
ENV XDG_CONFIG_HOME /config
ENV XDG_DATA_HOME /data

VOLUME /config
VOLUME /data

EXPOSE 80
EXPOSE 443

WORKDIR /srv

CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]
