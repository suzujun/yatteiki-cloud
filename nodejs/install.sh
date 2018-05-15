#!/bin/bash

# install nodejs

apt-get install nodejs
apt-get install npm
update-alternatives --install /usr/bin/node node /usr/bin/nodejs 10

