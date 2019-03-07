#!/bin/bash

cd "$(dirname "${BASH_SOURCE}")"

git pull origin master

function doIt() {
	rsync --exclude ".git/" \
		--exclude ".DS_Store" \
		--exclude ".osx" \
		--exclude "setup.sh" \
		--exclude "README.md" \
		--exclude "LICENSE.md" \
		-avh --no-perms ./dots/ ~/
}

if [ "$1" == "--force" -o "$1" == "-f" ]; then
	doIt
else
	read -p "This may overwrite existing files in your home directory. Are you sure? (y/n) " -n 1
	echo ""
	if [[ $REPLY =~ ^[Yy]$ ]]; then
		doIt
	fi
	echo "Don't forget to run 'source ~/.zshrc' to update your zsh settings"
fi

unset -f doIt
