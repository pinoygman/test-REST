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

#For build-only, not for cf push
FROM golang

MAINTAINER Chia Chang "chia.chang@ge.com"

#ENTRYPOINT $GOPATH/src/{DHOME}


COPY ./ $GOPATH/src/{DHOME}

WORKDIR $GOPATH/src/{DHOME}

ENV ARTIFACT {ARTIFACT}
ENV DIST ./{DIST}
ENV REV {REV}
ENV https_proxy {PROXY}
ENV BUILD_TIME {BUILD_TIME}
ENV BUILD_VER {BUILD_VER}
ENV LDFLAGS "{LDFLAGS}"
ENV SQLDSN "{SQLDSN}"

RUN go env
RUN go get github.com/cloudfoundry-community/go-cfenv
RUN go get github.com/pborman/uuid
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get github.com/lib/pq
RUN go get github.com/jmoiron/sqlx
RUN go get gopkg.in/redis.v4
RUN make
RUN ls -al && ls /{DIST} -al && pwd
#RUN chmod +x $GOPATH/src/{DHOME}
# No need to listen to a port. busted!
#EXPOSE 8080
