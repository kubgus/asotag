#!/bin/bash

echo "Building Windows binary..."

mkdir -p bin
GOOS=windows GOARCH=amd64 go build -o bin/adventure-game-windows.exe
