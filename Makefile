#
#  Copyright (c) 2016 General Electric Company. All rights reserved.
#
#  The copyright to the computer software herein is the property of
#  General Electric Company. The software may be used and/or copied only
#  with the written permission of General Electric Company or in accordance
#  with the terms and conditions stipulated in the agreement/contract
#  under which the software has been supplied.
#
#  author: chia.chang@ge.com
#

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')


REV=v1
DIST=./dist
BINARY=${DIST}/pcs-${REV}
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.REV=${REV}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): pre-install darwin-build linux-build
	@echo Copying Settings
	@cp ./settings.json ./dist/

pre-install:
	@echo install libraries 
	@go get github.com/pborman/uuid
	@go get github.com/gorilla/mux
	@go get github.com/cloudfoundry-community/go-cfenv
	@go get github.com/rs/cors

darwin-build:
	@echo Creating Mac OS X artifact
	@GOOS=darwin go build ${LDFLAGS} -o ${BINARY}_darwin app.go

linux-build:
	@echo Creating amd64_x86 artifact
	@GOOS=linux go build ${LDFLAGS} -o ${BINARY}_linux app.go


.PHONY: install
install:
	go install ${LDFLAGS}

.PHONY: clean
clean:
	rm ${DIST}/*
#if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: deploy
deploy:
	@echo Checking CF Envs..
	@cf a &> /dev/null
	@if [ $$? -eq 0 ] ; then \
		echo Good you have logged in the CF; \
		cf push pcs-${USER} -c "./${DIST}/${BINARY}_linux" -b https://github.com/cloudfoundry/binary-buildpack.git; \
	 else \
		echo Please log in the Cloud Foundry org/space.; \
		exit 1; \
	 fi;