#/usr/bin/env bash

set -euo pipefail

echo "⛈  Symlinking shell scripts to ~/bin"
mkdir -p ~/bin
stow -d ./scripts -t ~/bin shell

echo "🐹 Installing Go scripts"
for dir in scripts/uno/cmd/*/; do
  cd $dir
  go mod download
  go install
  cd ../../../../
done
