FROM golang:1.21-alpine3.19 as build

ARG TARGETARCH

ENV XKCD_ARCHIVER_VERSION v0.1.0
ENV KUBO_VERSION v0.24.0
ENV FLARECTL_VERSION v0.83.0

RUN apk update && \
  apk add --no-cache git ca-certificates openssh openssl curl gcompat

RUN go install go.hacdias.com/xkcd-archiver@$XKCD_ARCHIVER_VERSION && \
  go install github.com/ipld/go-car/cmd/car@latest && \
  go install github.com/cloudflare/cloudflare-go/cmd/flarectl@$FLARECTL_VERSION && \
  cp $GOPATH/bin/xkcd-archiver /usr/bin && \
  cp $GOPATH/bin/car /usr/bin && \
  cp $GOPATH/bin/flarectl /usr/bin

RUN curl https://dist.ipfs.tech/kubo/${KUBO_VERSION}/kubo_${KUBO_VERSION}_linux-${TARGETARCH}.tar.gz -o kubo.tar.gz && \
  tar -xvf kubo.tar.gz && \
  mv kubo/ipfs /usr/bin/ && \
  rm -rf kubo.tar.gz

FROM alpine:3.19

COPY --from=0 /usr/bin/xkcd-archiver /usr/bin/xkcd-archiver
COPY --from=0 /usr/bin/car /usr/bin/car
COPY --from=0 /usr/bin/flarectl /usr/bin/flarectl
COPY --from=0 /usr/bin/ipfs /usr/bin/ipfs

RUN apk update && \
  apk add --no-cache git ca-certificates openssh openssl curl gcompat

RUN mkdir -p /app/output
COPY ./exec.sh /app/exec.sh
WORKDIR /app

VOLUME /app/output
ENTRYPOINT [ "/app/exec.sh" ]
