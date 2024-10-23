#!/bin/bash

if [[ "$EUID" -ne 0 ]]; then
    echo "Darkleaf must be run as root!"
    exit 1
fi

clear


./scripts/build-runtime.sh
./runtime/bin/runtime.exe