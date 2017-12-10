# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef TRAVIS
	export GOROOT=$(realpath ../go)
	export GOPATH=$(realpath ../paligo)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

GO_VERSION=1.9.2

WEBSITE_DIR=004-innerHTML-textContent

devserver: fmt js
	@# http://stackoverflow.com/a/5947779
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run server.go -dir=$(WEBSITE_DIR)

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@gopherjs build $(WEBSITE_DIR)/app.go -o $(WEBSITE_DIR)/app.js

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt server.go

lib_gopherjs:
	@echo "\033[92mInstalling GopherJS ...\033[0m"
	go get -u github.com/gopherjs/gopherjs

lib_godom:
	@echo "\033[92mInstalling godom ...\033[0m"
	go get -u github.com/siongui/godom

clean:
	rm $(WEBSITE_DIR)/*.js
	rm $(WEBSITE_DIR)/*.js.map

update_ubuntu:
	@echo "\033[92mUpdating Ubuntu ...\033[0m"
	@sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@cd ../ ; wget https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
	@cd ../ ; tar xvzf go$(GO_VERSION).linux-amd64.tar.gz
