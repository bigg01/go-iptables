# go-iptables

# GO BUILD Linux
```
GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o iptablenforcer cmd/iptablenforcer/main.go
```
# GO BUILD DOCKER
```
docker build --tag iptablenforcer .
```

#### later we will use  releaser
https://goreleaser.com/
```
goreleaser init
```