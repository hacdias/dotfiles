#/usr/bin/env bash

set -euo pipefail

echo "⛈  Symlinking shell scripts to ~/bin"
mkdir -p ~/bin
stow -d ./scripts -t ~/bin -D shell

echo "🔌 Symlinking JS scripts"
cd scripts/js
npm unlink &>/dev/null
npm link &>/dev/null

echo "🔨 Installing miscellaneous tools"
go install github.com/StackExchange/dnscontrol
