#!/bin/bash
echo "======> Building the binary"
docker run --rm -v "$PWD":/usr/src/ecr \
  -w /usr/src/ecr \
  -e GO111MODULE=on \
  -e GOPROXY=https://goproxy.io \
  -e GOOS=linux \
  -e GOARCH=amd64 \
  -e  CGO_ENABLED=0 \
  golang:1.12 \
  go build -o ecr