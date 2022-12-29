#/usr/bin/env bash

set -euo pipefail

echo "⛈  Symlinking shell scripts to ~/bin"
mkdir -p ~/bin
stow -d ./scripts -t ~/bin shell

# echo "🔌 Symlinking JS scripts"
# cd scripts/js
# npm unlink .
# npm link .
# cd ../../

echo "🐹 Installing Go scripts"
for dir in scripts/go/*/; do
  cd $dir
  go mod download
  go install
  cd ../../
done
