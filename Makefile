
install: install-go install-nginx install-node

install-go:
	apt install golang-1.9
	mv /usr/bin/go /usr/bin/go_16
	mv /usr/bin/gofmt /usr/bin/gofmt_16
	ln -s /usr/lib/go-1.9/bin/go /usr/bin/go
	ln -s /usr/lib/go-1.9/bin/gofmt /usr/bin/gofmt

install-nginx:
	apt-get install net-tools ssh git emacs-nox zsh
	curl http://nginx.org/keys/nginx_signing.key | sudo apt-key add -
	VCNAME=`cat /etc/lsb-release | grep DISTRIB_CODENAME | cut -d= -f2` && sudo -E sh -c "echo \"deb http://nginx.org/packages/ubuntu/ $VCNAME nginx\" >> /etc/apt/sources.list"
	VCNAME=`cat /etc/lsb-release | grep DISTRIB_CODENAME | cut -d= -f2` && sudo -E sh -c "echo \"deb-src http://nginx.org/packages/ubuntu/ $VCNAME nginx\" >> /etc/apt/sources.list"
	apt-get update
	apt-get install nginx

install-node:
	apt-get install -y nodejs npm
	npm cache clean
	npm install n -g
	n stable
	ln -sf /usr/local/bin/node /usr/bin/node

run: run-goapp run-nodeapp

run-goapp:
	go run goapp/main.go &

run-nodeapp:
	cd nodeapp && npm start
