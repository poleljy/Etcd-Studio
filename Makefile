# run from repository root

objects = EtcdStudio

SHELL = /bin/bash

REPO_PATH=$(shell pwd)
GOBIN=$(REPO_PATH)/Release

# 编译时间
COMPILE_TIME = $(shell date +"%Y-%M-%d %H:%M:%S")

# GIT版本号
GIT_REVISION = $(shell git show -s --pretty=format:%h)

.PHONY: all
all:
	@$(build-webui)
	@$(build-module)
	@$(compress-exe)
	@$(release)

.PHONY: clean
clean: go-clean
	@rm -rf $(GOBIN)

define build-webui
	@echo "building webui"
	#@npm set registry https://registry.npm.taobao.org
	#@npm set disturl https://npm.taobao.org/dist
	#@npm cache clean --force
	#@npm install -g cnpm --registry=https://registry.npm.taobao.org
	#@cnpm install vue -g
	#@cnpm install vue-cli -g
	#@cnpm install vue-cli-service -g
	@cd webui && cnpm install && npm run build
endef

define build-module
	@echo "building service"
	@go get -u github.com/go-bindata/go-bindata/...
	@./go-bindata -o src/bindata.go -fs -prefix "webui/dist" webui/dist/...
	@cd src;GOOS=linux GOARCH=amd64 go build -o ../Release/$(objects) . 
	@cd src;GOOS=windows GOARCH=amd64 go build -o ../Release/$(objects).exe . 
endef

define compress-exe
	@echo "compress executor"
	@cd Release;../upx $(objects)
	@cd Release;../upx $(objects).exe
endef

define release
	@echo "release "
	@cd Release;tar czvf $(objects)_$(GIT_REVISION).tar.gz $(objects) $(objects).exe
	@rm -rf ./Release/$(objects) ./Release/$(objects).exe
endef

.PHONY: go-clean
go-clean:
	@echo "cleaning build cache"
	@GOBIN=$(GOBIN) go clean
