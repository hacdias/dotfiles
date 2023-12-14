#!/usr/bin/env sh

set -eou pipefail

xkcd-archiver output

car create --version 1 --output output.car --no-wrap output

cid=$(car root output.car)
echo $cid

record="dnslink=/ipfs/$cid"
echo $record

ipfs --api $IPFS_API dag import output.car --pin-roots

flarectl dns create-or-update --zone $CF_ZONE \
  --type "TXT" \
  --name $CF_RECORD_NAME \
  --content $record
