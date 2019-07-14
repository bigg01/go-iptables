FROM golang:1.12.7-alpine AS builder
WORKDIR /go/src/
RUN apk --no-cache add ca-certificates git upx

COPY pkg  /go/src/pkg
COPY cmd  /go/src/cmd
COPY go.* /go/src/

#RUN cmd/iptablenforcer/main.go
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o iptablenforcer cmd/iptablenforcer/main.go && \
    upx --ultra-brute -qq iptablenforcer && \
    upx -t iptablenforcer

FROM scratch
WORKDIR /usr/local/bin
COPY --from=builder /go/src/iptablenforcer .
CMD ["iptablenforcer"]