
install: install-first install-go install-nginx install-node

install-first:
	apt-get update
	apt-get upgrade
	apt-get install net-tools ssh git emacs-nox zsh apache2-utils curl

install-go:
	apt install golang-1.9
	#mv /usr/bin/go /usr/bin/go_16
	#mv /usr/bin/gofmt /usr/bin/gofmt_16
	ln -s /usr/lib/go-1.9/bin/go /usr/bin/go
	ln -s /usr/lib/go-1.9/bin/gofmt /usr/bin/gofmt

install-nginx:
	apt-get install net-tools ssh git emacs-nox zsh
	curl http://nginx.org/keys/nginx_signing.key | apt-key add -
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

setup-nginx:
	htpasswd -c .htpasswd username
	cp .htpasswd /etc/nginx/conf.d
	mv /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf.bk
	cp ./nginx/default.conf /etc/nginx/conf.d

setup-goapp:
	add-apt-repository ppa:masterminds/glide && sudo apt-get update
	apt-get install glide
	cd goapp && glide i

setup-nodeapp:
	cd nodeapp && npm i

run: run-goapp run-nodeapp

run-goapp:
	DB_MASTER_HOST=suzujun-dbm.xxx-dev.local \
	DB_SLAVE_HOST=suzujun-dbs.xxx-dev.local \
	go run goapp/main.go &

run-nodeapp:
	cd nodeapp && npm start
