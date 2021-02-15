#!/bin/bash
export GOROOT=/go
export GOPATH=/backend-go/
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

go version
cd /backend-go || exit; pwd; ls; cat main.go; go run main.go;