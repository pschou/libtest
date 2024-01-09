PROG_NAME := "libtest"
VERSION = 0.1.$(shell date -u +%Y%m%d.%H%M)
FLAGS := "-s -w -X main.version=${VERSION}"


build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags=${FLAGS} -o ${PROG_NAME}64 .
	upx --lzma ${PROG_NAME}64
	GOOS=linux GOARCH=386 CGO_ENABLED=1 go build -ldflags=${FLAGS} -o ${PROG_NAME}32 .
	upx --lzma ${PROG_NAME}32
