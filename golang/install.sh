#!/bin/bash

# install golang

apt install golang-1.9
mv /usr/bin/go /usr/bin/go_16
mv /usr/bin/gofmt /usr/bin/gofmt_16
ln -s /usr/lib/go-1.9/bin/go /usr/bin/go
ln -s /usr/lib/go-1.9/bin/gofmt /usr/bin/gofmt

