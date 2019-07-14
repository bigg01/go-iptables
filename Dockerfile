FROM golang:1.12.7-alpine AS builder
WORKDIR /go/src/iptablenforcer
RUN apk --no-cache add ca-certificates git upx

COPY ./*  /go/src/iptablenforcer/
RUN ls -lht * && go build cmd/iptablenforcer/main.go

#CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o iptablenforcer . && \
#    upx --ultra-brute -qq iptablenforcer && \
#    upx -t iptablenforcer

FROM scratch
EXPOSE 8080
WORKDIR /usr/local/bin
COPY --from=builder /go/src/topz/cmd/iptablenforcer/iptablenforcer .
CMD ["iptablenforcer"]