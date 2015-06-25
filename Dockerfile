FROM gliderlabs/alpine:3.2
ENTRYPOINT ["/bin/signal-test"]

COPY . /go/src/github.com/Clever/signal-test
RUN apk-install -t build-deps go git \
    && cd /go/src/github.com/Clever/signal-test \
    && export GOPATH=/go \
    && go get \
    && go build -o /bin/signal-test \
    && rm -rf /go \
    && apk del --purge build-deps
