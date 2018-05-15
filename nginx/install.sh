#!/bin/bash

# @see https://qiita.com/mionatsume/items/1d802723c370984d544b
# install nginx

apt-get install net-tools ssh git emacs-nox zsh

curl http://nginx.org/keys/nginx_signing.key | sudo apt-key add -
VCNAME=`cat /etc/lsb-release | grep DISTRIB_CODENAME | cut -d= -f2` && sudo -E sh -c "echo \"deb http://nginx.org/packages/ubuntu/ $VCNAME nginx\" >> /etc/apt/sources.list"
VCNAME=`cat /etc/lsb-release | grep DISTRIB_CODENAME | cut -d= -f2` && sudo -E sh -c "echo \"deb-src http://nginx.org/packages/ubuntu/ $VCNAME nginx\" >> /etc/apt/sources.list"
apt-get update
apt-get install nginx

