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
	ffmpeg \
	gcc \
	node \
	wget \
	mas \
	go \
	hugo \
	ufraw

brew install imagemagick --with-webp
brew cleanup

mas lucky todoist
mas lucky 1password
mas lucky spark
mas lucky lungo
mas lucky slack
mas lucky telegram
