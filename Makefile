PROG_NAME := "libtest"
VERSION = 0.1.$(shell date -u +%Y%m%d.%H%M)
FLAGS := "-s -w -X main.version=${VERSION}"


build:
	CGO_ENABLED=0 go build -ldflags=${FLAGS} -o ${PROG_NAME} .
	upx --lzma ${PROG_NAME}
