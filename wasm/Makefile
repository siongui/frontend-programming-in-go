# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


DIR=005-xmlhttprequest-xhr
MODULE=xhr.wasm

server: build
	@echo "\033[92mRunning dev server...\033[0m"
	go run ../server.go -dir=${DIR}/demo

build: fmt
	@echo "\033[92mBuild wasm module...\033[0m"
	GOARCH=wasm GOOS=js go build -o ${DIR}/demo/${MODULE} ${DIR}/*.go
	cp ${DIR}/index.html ${DIR}/demo/
	cd ${DIR}/demo/ ; ln -sf ../../001-hello-world/demo/wasm_exec.js wasm_exec.js

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt ${DIR}/*.go

install:
	GOARCH=wasm GOOS=js go get -u github.com/siongui/godom/wasm
