#/usr/bin/env bash

set -euo pipefail

echo "⛈  Symlinking shell scripts to ~/bin"
mkdir -p ~/bin
stow -d ./scripts -t ~/bin shell

echo "🔌 Symlinking JS scripts"
cd scripts/js
npm unlink .
npm link .

echo "🔨 Installing miscellaneous tools"
go install github.com/StackExchange/dnscontrol
