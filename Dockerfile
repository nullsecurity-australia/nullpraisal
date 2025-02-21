FROM golang:1.21-alpine
LABEL maintainer="peer <https://github.com/nullsecurity-australia/nullpraisal>"
WORKDIR $GOPATH/src/github.com/nullsecurity-australia/nullpraisal
RUN apk --update add --no-cache --virtual build-dependencies git gcc musl-dev make bash && \
    git clone https://github.com/nullsecurity-australia/nullpraisal.git . && \
    export GO111MODULE=on ENV=prod && \
    make setup && \
    make build && \
    make install && \
    mkdir /evepraisal/ && \
    mv $GOPATH/bin/evepraisal /evepraisal/evepraisal && \
    rm -rf $GOPATH && \
    apk del build-dependencies && \
    mkdir /evepraisal/db
WORKDIR /evepraisal/
CMD ["./evepraisal"]
