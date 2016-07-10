SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=roll

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/ariejan/roll/core.Version=${VERSION} -X github.com/ariejan/roll/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
    go build ${LDFLAGS} -o ${BINARY} main.go

.PHONY: install
install:
    go install ${LDFLAGS} ./...

.PHONY: clean
clean:
    if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
