FROM golang:1.12.7-alpine AS builder
WORKDIR /go/src/iptablenforcer
RUN apk --no-cache add ca-certificates git upx

COPY pkg  /go/src/iptablenforcer/pkg
COPY cmd  /go/src/iptablenforcer/cmd
RUN go get -v ./... && \
    cd /go/src/iptablenforcer/cmd/iptablenforcer && \
    CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o iptablenforcer . && \
    upx --ultra-brute -qq iptablenforcer && \
    upx -t iptablenforcer

FROM scratch
EXPOSE 8080
WORKDIR /usr/local/bin
COPY --from=builder /go/src/topz/cmd/iptablenforcer/iptablenforcer .
CMD ["iptablenforcer"]