#!/bin/bash

cd "$(dirname "$0")/../runtime"

mkdir -p bin

go build -o bin/runtime.exe cmd/app/main.go

if [ $? -eq 0 ]; then
  echo "Build successful. Binary is located in runtime/bin/runtime.exe"
else
  echo "Build failed."
  exit 1
fi
