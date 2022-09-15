GC=go
MAIN_GO_Dir=.
APP_NAME=ms-client
RELEASE_DIR=tools
VERSION=0.0.1

build:
	${GC} build -ldflags '-X main.version=${VERSION} -s -w' -o ${RELEASE_DIR}/${APP_NAME} ${MAIN_GO_Dir}

pkg :
	@echo "cp readme ..."
	cp -f readme.md ${RELEASE_DIR}

clean :
	@if [ -d ${RELEASE_DIR} ] ; then rm -rf ${RELEASE_DIR}; fi
