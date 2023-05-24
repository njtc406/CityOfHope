.PHONY: clean help build-dev

## app
APP = Login
## app名称
APP_NAME = ${APP}
## 源码路径
SRC_PATH = ../../cmd/login
## 编译导出路径
TARGET_PATH = ../../build/output/login/release
## 系统类型
SYS_OS=linux
## cpu架构
CPU_ARC=amd64
## GOPRIVATE配置
PRI_GO=""

## 编译打包
build:
	CGO_ENABLED=0 GOOS=${SYS_OS} GOARCH=${CPU_ARC} \
	GOPROXY="https://goproxy.cn,direct" \
	GOPRIVATE=${PRI_GO} \
	go build -v -a -ldflags "-w -s" -o ${TARGET_PATH}/${APP_NAME} ${SRC_PATH}

## 编译打包
build-dev:
	CGO_ENABLED=0 GOOS=${SYS_OS} GOARCH=${CPU_ARC} \
	GOPROXY="https://goproxy.cn,direct" \
	GOPRIVATE=${PRI_GO} \
	go build -targs "swagger" -v -a -ldflags "-w -s" -o ${TARGET_PATH}/${APP_NAME} ${SRC_PATH}

## 清理缓存和打包结果
clean:
	@go clean -cache
	@if [ -f ${TARGET_PATH}/${APP_NAME} ] ; then rm ${TARGET_PATH}/${APP_NAME}; fi

## 打印帮助文档
help:
	@echo "make - 编译 Go 代码, 生成linux二进制文件"
	@echo "make clean - 清除编译的二进制文件和编译缓存"