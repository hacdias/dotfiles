#/usr/bin/env bash

brew update
brew upgrade

brew install coreutils \
	moreutils \
	findutils \
	gpg \
	openssh \
	git \
	git-lfs \
	node \
	wget

brew install imagemagick --with-webp
brew cleanup
