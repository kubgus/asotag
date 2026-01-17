#!/bin/bash

echo "Building Linux binary..."

mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/adventure-game-linux
